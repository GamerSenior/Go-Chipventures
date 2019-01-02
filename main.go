package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game struct contendo todos os dados do jogo
type Game struct {
	Player       Player
	ScreenWidth  int32
	ScreenHeight int32
	Title        string
	FrameCounter int32
	GameOver     bool
	Pause        bool

	Dispatcher Dispatcher
}

func main() {
	game := Game{}
	game.Init()

	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, game.Title)
	InitSystems(&game)

	rl.SetTargetFPS(int32(60))
	for !rl.WindowShouldClose() {

		key := rl.GetKeyPressed()
		if key != -1 {
			game.Dispatcher.dispatch("keyPressed", &game)
		}

		rl.BeginDrawing()
		game.Player.Update()
		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

// Init inicializa estrutura do jogo
func (game *Game) Init() {
	game.ScreenHeight = 400
	game.ScreenWidth = 600
	game.FrameCounter = 60
	game.GameOver = false
	game.Pause = false
	game.Title = "Chipventures - In Go Lang"
	game.Player = NewPlayer()
}
