package main

import (
	"math"

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
	KeysDown   map[int32]bool
}

func main() {
	game := Game{}
	game.Init()

	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, game.Title)
	InitSystems(&game)

	rl.SetTargetFPS(int32(60))

	// Testando box2D
	game.World.Clear()
	var b1 box2d.Body
	b1.Set(&box2d.Vec2{X: 100.0, Y: 20.0}, math.MaxFloat64)
	b1.Position = box2d.Vec2{X: 0.0, Y: -17}
	b1.Friction = 0.5
	game.World.AddBody(&b1)

	game.World.AddBody(&game.Player.rigidBody)
	//---------------

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		key := rl.GetKeyPressed()
		if key != -1 {
			game.Dispatcher.dispatch("keyPressed", &game)
		}
		for k := range game.KeysDown {
			if rl.IsKeyReleased(k) {
				game.Dispatcher.dispatch("keyReleased", &game, k)
			}
		}

		game.Player.Update()

		//------ Debbuging --------
		// fmt.Println("Ground Position:", b1.Position)
		// fmt.Println("Player position: ", game.Player.rigidBody.Position)
		//-------------------------

		rl.ClearBackground(rl.RayWhite)
		rl.EndDrawing()

		game.TimeStep = float64(rl.GetFrameTime())

		game.World.Step(game.TimeStep)
	}

	rl.CloseWindow()
}

// Init inicializa estrutura do jogo
func (g *Game) Init() {
	g.ScreenHeight = 400
	g.ScreenWidth = 600
	g.FrameCounter = 60
	g.GameOver = false
	g.Pause = false
	g.Title = "Chipventures - In Go Lang"
	g.KeysDown = make(map[int32]bool)
}
