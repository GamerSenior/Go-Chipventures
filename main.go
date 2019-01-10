package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
)

// PPM (Pixels por Metro)
const PPM float64 = 100.0

// Game struct contendo todos os dados do jogo
type Game struct {
	Player       Player
	ScreenWidth  int32
	ScreenHeight int32
	Title        string
	FrameCounter int32
	GameOver     bool
	Pause        bool
	Debbug       bool

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
	floorY2 := 32 / PPM
	b1.Set(&box2d.Vec2{X: 100.0, Y: floorY2}, math.MaxFloat64)
	b1.Position = box2d.Vec2{X: 0.0, Y: -(float64(game.ScreenHeight) / PPM) + 1}
	b1.Friction = 0.5
	game.World.AddBody(&b1)

	game.World.AddBody(&game.Player.rigidBody)
	//---------------

	// Testando Draw de Tileset
	// ft := loadTileset("resources/tileset.png")
	// ---------------

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
		if game.Debbug {
			DrawHitbox(game.Player)
			DrawHitboxRect(GetHitbox(b1))
		}
		// rl.DrawTextureRec(ft, rl.Rectangle{0, 0, 32, 32}, rl.Vector2{0, 768}, rl.White)
		// rl.DrawTextureRec(ft, rl.Rectangle{32, 0, 32, 32}, rl.Vector2{32, 768}, rl.White)
		// rl.DrawTextureRec(ft, rl.Rectangle{32, 0, 32, 32}, rl.Vector2{64, 768}, rl.White)
		// rl.DrawTextureRec(ft, rl.Rectangle{32, 0, 32, 32}, rl.Vector2{96, 768}, rl.White)

		fmt.Println("Ground Position:", b1.Position)
		fmt.Printf("Player position: %.2f\n", game.Player.rigidBody.Position)

		contacts := box2d.Collide(&b1, &game.Player.rigidBody)
		if len(contacts) > 0 {
			fmt.Println("Contact point 1: ", contacts[0].Position)
			fmt.Println("Contact point 2: ", contacts[1].Position)
		}
		// fmt.Printf("Collision point: [%.2f, %.2f]", x1, y1)

		//-------------------------

		rl.ClearBackground(rl.Blue)
		rl.EndDrawing()

		game.TimeStep = float64(rl.GetFrameTime())

		game.World.Step(game.TimeStep)
	}

	rl.CloseWindow()
}

// Init inicializa estrutura do jogo
func (g *Game) Init() {
	g.ScreenHeight = 800
	g.ScreenWidth = 1280
	g.FrameCounter = 60
	g.GameOver = false
	g.Pause = false
	g.Title = "Chipventures - In Go Lang"
	g.KeysDown = make(map[int32]bool)
	g.Debbug = true
}

// GameObject é uma interface que qualquer objeto do jogo deve implementar
type GameObject interface {
	GetHitbox() rl.Rectangle
}

// DrawHitbox desenha na janela a hitbox de um GameObject
func DrawHitbox(o GameObject) {
	rl.DrawRectangleLinesEx(o.GetHitbox(), 1, rl.Green)
}

func DrawHitboxRect(r rl.Rectangle) {
	rl.DrawRectangleLinesEx(r, 1, rl.Green)
}

func GetHitbox(b box2d.Body) (r rl.Rectangle) {
	r.X = float32(b.Position.X * -PPM)
	r.Y = float32(b.Position.Y * -PPM)
	r.Width = float32(b.Width.X * PPM)
	r.Height = float32(b.Width.Y * PPM)
	return
}
