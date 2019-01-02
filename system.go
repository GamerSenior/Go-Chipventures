package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// InitSystems inicializa os principais sistemas da aplicação
func InitSystems(game *Game) {
	keyboardDispatcher := NewDispatcher()
	fmt.Println("Dispatcher do Keyboard criado")
	keyboardDispatcher.on("keyPressed", onKeyPressed)
	keyboardDispatcher.on("movePlayer", movePlayer)
	game.KeyboardDispatcher = keyboardDispatcher
}

func onKeyPressed(i ...interface{}) {
	game := i[0].(*Game)
	if rl.IsKeyDown(rl.KeyS) {
		game.KeyboardDispatcher.dispatch("movePlayer", &game.Player, rl.Vector2{X: 0, Y: 1})
	}
	if rl.IsKeyDown(rl.KeyW) {
		game.KeyboardDispatcher.dispatch("movePlayer", &game.Player, rl.Vector2{X: 0, Y: -1})
	}
	if rl.IsKeyDown(rl.KeyD) {
		game.KeyboardDispatcher.dispatch("movePlayer", &game.Player, rl.Vector2{X: 1, Y: 0})
	}
	if rl.IsKeyDown(rl.KeyA) {
		game.KeyboardDispatcher.dispatch("movePlayer", &game.Player, rl.Vector2{X: -1, Y: 0})
	}
}

func movePlayer(i ...interface{}) {
	player := i[0].(*Player)
	speed := i[1].(rl.Vector2)
	player.position.X = player.position.X + speed.X*player.speed.X
	player.position.Y = player.position.Y + speed.Y*player.speed.Y
}
