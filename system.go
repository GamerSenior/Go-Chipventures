package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
)

// InitSystems inicializa os principais sistemas da aplicação
func InitSystems(game *Game) {
	gravity := box2d.Vec2{X: 0.0, Y: -10.0}
	game.World = box2d.NewWorld(gravity, 10)
	game.Player = NewPlayer()

	dispatcher := NewDispatcher()
	fmt.Println("Dispatcher do Keyboard criado")
	dispatcher.on("keyPressed", onKeyPressed)
	dispatcher.on("movePlayer", movePlayer)
	game.Dispatcher = dispatcher
}

func onKeyPressed(i ...interface{}) {
	game := i[0].(*Game)
	if rl.IsKeyDown(rl.KeyS) {
		// game.Dispatcher.dispatch("movePlayer", &game.Player, rl.Vector2{X: 0, Y: 1})
	}
	if rl.IsKeyDown(rl.KeyW) {
		// game.Dispatcher.dispatch("movePlayer", &game.Player, rl.Vector2{X: 0, Y: -1})
	}
	if rl.IsKeyDown(rl.KeyD) {
		game.Dispatcher.dispatch("movePlayer", &game.Player, box2d.Vec2{X: 1, Y: 0})
	}
	if rl.IsKeyDown(rl.KeyA) {
		game.Dispatcher.dispatch("movePlayer", &game.Player, box2d.Vec2{X: -1, Y: 0})
	}
}

func movePlayer(i ...interface{}) {
	player := i[0].(*Player)
	player.rigidBody.Velocity.X = i[1].(box2d.Vec2).X * -1
	// speed := i[1].(rl.Vector2)
	// player.position.X = player.position.X + speed.X*player.speed.X
	// player.position.Y = player.position.Y + speed.Y*player.speed.Y
}
