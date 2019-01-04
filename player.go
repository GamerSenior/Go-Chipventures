package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	box2d "github.com/neguse/go-box2d-lite/box2dlite"
)

// Player é uma estrutura contendo as informações básicas do jogador
type Player struct {
	texture   rl.Texture2D
	rigidBody box2d.Body
}

//NewPlayer retorna uma nova instancia de player
func NewPlayer() (p Player) {
	img := rl.LoadImage("resources/chip.png")
	p.texture = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)

	p.rigidBody.Set(&box2d.Vec2{X: 1.0, Y: 1.0}, 200.0)
	p.rigidBody.Position = box2d.Vec2{X: 0.0, Y: 0.0}
	p.rigidBody.Friction = 0.2
	return
}

// Update atualiza os dados do Player
func (p *Player) Update() {
	pos := p.rigidBody.Position
	xPixel := int32(-pos.X * 30)
	yPixel := int32(-pos.Y * 30)
	fmt.Println(pos)
	rl.DrawTexture(p.texture, xPixel, yPixel, rl.RayWhite)
}
