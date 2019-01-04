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

	p.rigidBody.Set(&box2d.Vec2{X: 1.0, Y: 1.0}, 50.0)
	p.rigidBody.Position = box2d.Vec2{X: 0.0, Y: 0.0}
	p.rigidBody.Friction = 0.2
	return
}

// Update atualiza os dados do Player
func (p *Player) Update() {
	pos := p.rigidBody.Position
	xPixel := int32(-pos.X * PPM)
	yPixel := int32(-pos.Y * PPM)
	fmt.Printf("X: %.2f Y: %2f\n", pos.X, pos.Y)
	rl.DrawTexture(p.texture, xPixel, yPixel, rl.RayWhite)
}

// GetHitbox retorna um retangulo contendo a hitbox do Player
func (p Player) GetHitbox() (h rl.Rectangle) {
	h.Width = float32(p.texture.Width)
	h.Height = float32(p.texture.Height)
	h.X = float32(-p.rigidBody.Position.X * PPM)
	h.Y = float32(-p.rigidBody.Position.Y * PPM)
	return
}
