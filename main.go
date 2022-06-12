package main

import (
	"fmt"
	"github.com/lcnssantos/snake-game/app"
	"github.com/lcnssantos/snake-game/infra/ui"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

func main() {
	wnd, cv, err := sdlcanvas.CreateWindow(640, 640, "Snake Game")

	if err != nil {
		panic(err)
	}
	defer wnd.Destroy()

	width, height := wnd.Size()

	game := app.NewGame(int64(width), int64(height), func() {
		fmt.Println("Game over")
	})

	game.Init()

	userInterface := ui.NewUi(wnd, cv, &game)

	userInterface.Start()
}
