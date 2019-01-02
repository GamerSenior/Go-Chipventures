package main

import (
	"container/list"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game struct contendo todos os dados do jogo
type Game struct {
	ScreenWidth  int32
	ScreenHeight int32
	Title        string
	FrameCounter int32
	GameOver     bool
	Pause        bool
	Handlers     *list.List
}

func main() {
	game := Game{}
	game.Init()

	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, game.Title)
	InitSystems(game)
	//rl.InitWindow(int32(600), int32(400), "Teste")

	rl.SetTargetFPS(int32(60))
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
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
	game.Handlers = list.New()
	game.Title = "Chipventures - In Go Lang"
}
