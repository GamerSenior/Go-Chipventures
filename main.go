package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
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
	World      *box2d.World
	TimeStep   float64
}

func main() {
	game := Game{}
	game.Init()

	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, game.Title)
	InitSystems(&game)

	rl.SetTargetFPS(int32(60))

	// Testando box2D
	game.World.Clear()
	//---------------

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		key := rl.GetKeyPressed()
		if key != -1 {
			game.Dispatcher.dispatch("keyPressed", &game)
		}

		game.Player.Update()

		game.Draw()

		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()

		game.TimeStep = float64(rl.GetFrameTime())

		game.World.Step(game.TimeStep)
	}

	rl.CloseWindow()
}

// Draw desenha todos os objetos pertencentes ao mundo
func (g *Game) Draw() {
	// for _, b := range g.World.Bodies {
	// Implementar funcão de desenhar bodies
	// }
}

// Init inicializa estrutura do jogo
func (g *Game) Init() {
	g.ScreenHeight = 400
	g.ScreenWidth = 600
	g.FrameCounter = 60
	g.GameOver = false
	g.Pause = false
	g.Title = "Chipventures - In Go Lang"
}
