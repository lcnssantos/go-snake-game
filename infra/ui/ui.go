package ui

import (
	"github.com/lcnssantos/snake-game/app"
	"github.com/lcnssantos/snake-game/domain"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
	"github.com/veandco/go-sdl2/sdl"
)

type Ui struct {
	window *sdlcanvas.Window
	canvas *canvas.Canvas
	game   *app.Game
}

func NewUi(window *sdlcanvas.Window, canvas *canvas.Canvas, game *app.Game) *Ui {
	return &Ui{window: window, canvas: canvas, game: game}
}

func (ui *Ui) drawSquare(square domain.Square, color string) {
	_width, _height := ui.window.Size()
	width := int64(_width)
	height := int64(_height)

	if square.Position.X < 0 || square.Position.X > width || square.Position.Y < 0 || square.Position.Y > height {
		ui.game.OnGameOver()
		ui.game.BuildSnake(nil)
		return
	}

	ctx := ui.canvas

	ctx.BeginPath()
	ctx.SetStrokeStyle(color)
	ctx.SetFillStyle(color)
	ctx.FillRect(float64(square.Position.X), float64(square.Position.Y), float64(square.Size), float64(square.Size))
	ctx.Stroke()
}

func (ui *Ui) drawSquares() {
	squares := ui.game.GetSquares()

	ui.canvas.ClearRect(0, 0, float64(ui.canvas.Width()), float64(ui.canvas.Height()))

	for _, square := range squares {
		ui.drawSquare(square, "#4ade80")
	}

	if ui.game.Frame%20 == 0 {
		ui.drawSquare(ui.game.Target, "#818cf8")
	}
}

func (ui *Ui) Start() {

	ui.window.Event = func(event sdl.Event) {
		switch event.GetType() {
		case sdl.KEYDOWN:
			ui.handleKeyDown(event)
			break
		case sdl.KEYUP:
			ui.handleKeyUp(event)
			break
		}
	}

	ui.window.MainLoop(func() {
		width, height := ui.window.Size()

		ui.game.SetSize(int64(width), int64(height))

		if !ui.game.IsPaused {
			ui.game.Run()
			ui.drawSquares()
		}
	})
}

func (ui *Ui) handleKeyDown(event sdl.Event) {
	keyboardEvent := event.(*sdl.KeyboardEvent)
	switch keyboardEvent.Keysym.Scancode {
	case sdl.SCANCODE_W, sdl.SCANCODE_UP:
		ui.game.OnKeyPress(app.KeyUp)
		break
	case sdl.SCANCODE_A, sdl.SCANCODE_LEFT:
		ui.game.OnKeyPress(app.KeyLeft)
		break
	case sdl.SCANCODE_S, sdl.SCANCODE_DOWN:
		ui.game.OnKeyPress(app.KeyDown)
		break
	case sdl.SCANCODE_D, sdl.SCANCODE_RIGHT:
		ui.game.OnKeyPress(app.KeyRight)
		break
	case sdl.SCANCODE_P, sdl.SCANCODE_PAUSE:
		ui.game.OnKeyPress(app.KeyP)
		break
	case sdl.SCANCODE_SPACE:
		ui.game.OnKeyPress(app.KeySpace)
		break
	}
}

func (ui *Ui) handleKeyUp(event sdl.Event) {
	keyboardEvent := event.(*sdl.KeyboardEvent)
	switch keyboardEvent.Keysym.Scancode {
	case sdl.SCANCODE_SPACE:
		ui.game.OnKeyDown(app.KeySpace)
		break
	}
}
