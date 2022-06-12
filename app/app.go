package app

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lcnssantos/snake-game/domain"
)

func (g *Game) BuildSnake(buildNew *bool) {
	snake := domain.NewSnake(domain.Position{
		X: g.width / 2,
		Y: g.height,
	}, DEFAULT_SQUARE_SIZE, 3)

	g.snake = &snake
}

type Game struct {
	width      int64
	height     int64
	IsPaused   bool
	OnGameOver func()
	snake      *domain.Snake
	speedUp    bool
	Target     domain.Square
	Frame      uint64
}

func NewGame(width, height int64, onGameOver func()) Game {
	return Game{
		width:      width,
		height:     height,
		IsPaused:   false,
		OnGameOver: onGameOver,
		speedUp:    false,
		Frame:      0,
	}
}

func (g *Game) SetSize(width, height int64) {
	g.width = width
	g.height = height
}

func (g *Game) Init() {
	build := true
	g.BuildSnake(&build)
	g.generateNewTarget()
}

func (g *Game) Run() {
	g.Frame++

	if g.IsPaused {
		return
	}

	if g.snake == nil {
		return
	}

	if !g.speedUp && g.Frame%10 != 0 {
		return
	}

	if g.speedUp && g.Frame%3 != 0 {
		return
	}

	err := g.snake.Move()

	if err != nil {
		g.OnGameOver()
		buildNew := true
		g.BuildSnake(&buildNew)
		fmt.Printf("Error: %s", err.Error())
		return
	}

	if g.snake.IsTouching(g.Target, nil) {
		g.snake.Add(g.Target)
		g.generateNewTarget()
	}

	fmt.Printf("IsPaused: %t\n", g.IsPaused)
	fmt.Printf("SpeedUp: %t\n", g.speedUp)
	fmt.Printf("Snake Frame: %d\n", g.snake.Frame)
	fmt.Printf("Game Frame: %d\n", g.Frame)

}

func (g *Game) OnKeyPress(key Key) {
	switch key {
	case KeyLeft:
		g.snake.SetDirection(domain.DirectionLeft)
		break
	case KeyRight:
		g.snake.SetDirection(domain.DirectionRight)
		break
	case KeyUp:
		g.snake.SetDirection(domain.DirectionUp)
		break
	case KeyDown:
		g.snake.SetDirection(domain.DirectionDown)
		break
	case KeySpace:
		g.speedUp = true
		break
	case KeyP:
		g.IsPaused = !g.IsPaused
		break
	}

	fmt.Print("Key pressed: ")
	fmt.Println(key)
}

func (g *Game) OnKeyDown(key Key) {
	if key == KeySpace {
		g.speedUp = false
	}
}

func (g *Game) generateNewTarget() {
	x := g.width / DEFAULT_SQUARE_SIZE
	y := g.height / DEFAULT_SQUARE_SIZE

	rand.Seed(time.Now().UnixNano())

	position := domain.Position{
		X: rand.Int63n(x) * DEFAULT_SQUARE_SIZE,
		Y: rand.Int63n(y) * DEFAULT_SQUARE_SIZE,
	}

	g.Target = domain.Square{
		Size:     DEFAULT_SQUARE_SIZE,
		Position: position,
	}
}

func (g *Game) GetSquares() []domain.Square {
	return g.snake.GetSquares()
}
