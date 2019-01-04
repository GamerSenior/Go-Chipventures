package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
)

const meterToPixel = 100.0
const offsetX = 6.0
const offsetY = 4.0

// InitSystems inicializa os principais sistemas da aplicação
func InitSystems(game *Game) {
	gravity := box2d.Vec2{X: 0.0, Y: -10.0}
	game.World = box2d.NewWorld(gravity, 10)
	game.Player = NewPlayer()

	dispatcher := NewDispatcher()
	fmt.Println("Dispatcher do Keyboard criado")
	dispatcher.on("keyPressed", onKeyPressed)
	dispatcher.on("keyReleased", onKeyReleased)
	dispatcher.on("movePlayer", movePlayer)
	game.Dispatcher = dispatcher
}

func onKeyPressed(i ...interface{}) {
	game := i[0].(*Game)
	if rl.IsKeyDown(rl.KeyS) {
		game.KeysDown[rl.KeyS] = true
		// game.Dispatcher.dispatch("movePlayer", &game.Player, rl.Vector2{X: 0, Y: 1})
	}
	if rl.IsKeyPressed(rl.KeyW) {
		game.KeysDown[rl.KeyW] = true
		game.Dispatcher.dispatch("movePlayer", game, box2d.Vec2{X: 0, Y: -1})
	}
	if rl.IsKeyDown(rl.KeyD) {
		game.KeysDown[rl.KeyD] = true
		game.Dispatcher.dispatch("movePlayer", game, box2d.Vec2{X: 1, Y: 0})
	}
	if rl.IsKeyDown(rl.KeyA) {
		game.KeysDown[rl.KeyA] = true
		game.Dispatcher.dispatch("movePlayer", game, box2d.Vec2{X: -1, Y: 0})
	}
}

func onKeyReleased(i ...interface{}) {
	game := i[0].(*Game)
	key := i[1].(int32)
	fmt.Println("Key released: ", key)
	delete(game.KeysDown, key)
}

func movePlayer(i ...interface{}) {
	// game := i[0].(*Game)
	player := &i[0].(*Game).Player
	mVector := i[1].(box2d.Vec2)

	if mVector.Y != 0 && math.Round(player.rigidBody.Velocity.Y) == 0 {
		player.rigidBody.Velocity.Y = mVector.Y * -8
	}
	fmt.Println("mVector: ", mVector.X)
	fmt.Println("Player velocity: ", math.Round(player.rigidBody.Velocity.X*100)/100)
	player.rigidBody.Velocity.X = mVector.X * -5
}
