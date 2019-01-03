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
}

func main() {
	game := Game{}
	game.Init()

	rl.InitWindow(game.ScreenWidth, game.ScreenHeight, game.Title)
	InitSystems(&game)

	rl.SetTargetFPS(int32(60))

	// Testando box2D
	game.World.Clear()
	var b1, b2 box2d.Body

	b1.Set(&box2d.Vec2{X: 100.0, Y: 20.0}, math.MaxFloat64)
	b1.Position = box2d.Vec2{X: 0.0, Y: -0.5 * b1.Width.Y}
	game.World.AddBody(&b1)

	b2.Set(&box2d.Vec2{X: 1.0, Y: 1.0}, 200)
	b2.Position = box2d.Vec2{X: 0.0, Y: 4.5 * b1.Width.Y}
	game.World.AddBody(&b2)

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
	for _, b := range g.World.Bodies {
		g.DrawBody(b)
	}
}

// DrawBody desenha um corpo do tipo box2d.Body
func (g *Game) DrawBody(b *box2d.Body) {
	R := box2d.Mat22ByAngle(b.Rotation)
	x := b.Position
	h := box2d.MulSV(0.5, b.Width)

	o := box2d.Vec2{X: 400, Y: 400}
	S := box2d.Mat22{Col1: box2d.Vec2{X: 20.0, Y: 0.0}, Col2: box2d.Vec2{X: 0.0, Y: -20.0}}

	v1 := o.Add(S.MulV(x.Add(R.MulV(box2d.Vec2{-h.X, -h.Y}))))
	v2 := o.Add(S.MulV(x.Add(R.MulV(box2d.Vec2{h.X, -h.Y}))))
	v3 := o.Add(S.MulV(x.Add(R.MulV(box2d.Vec2{h.X, h.Y}))))
	v4 := o.Add(S.MulV(x.Add(R.MulV(box2d.Vec2{-h.X, h.Y}))))

	rl.DrawLine(int32(v1.X), int32(v1.Y), int32(v2.X), int32(v2.Y), rl.RayWhite)
	rl.DrawLine(int32(v2.X), int32(v2.Y), int32(v3.X), int32(v3.Y), rl.RayWhite)
	rl.DrawLine(int32(v3.X), int32(v3.Y), int32(v4.X), int32(v4.Y), rl.RayWhite)
	rl.DrawLine(int32(v4.X), int32(v4.Y), int32(v1.X), int32(v1.Y), rl.RayWhite)
}

// Init inicializa estrutura do jogo
func (g *Game) Init() {
	g.ScreenHeight = 400
	g.ScreenWidth = 600
	g.FrameCounter = 60
	g.GameOver = false
	g.Pause = false
	g.Title = "Chipventures - In Go Lang"
	g.Player = NewPlayer()
}
