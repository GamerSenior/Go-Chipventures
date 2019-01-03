package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Player é uma estrutura contendo as informações básicas do jogador
type Player struct {
	texture  rl.Vector2
	position rl.Vector2
	speed    rl.Vector2
}

//NewPlayer retorna uma nova instancia de player
func NewPlayer() (p Player) {
	p.texture = rl.NewVector2(float32(50), float32(50))
	p.position = rl.NewVector2(float32(0), float32(0))
	p.speed = rl.NewVector2(float32(2), float32(2))
	return
}

// Update atualiza os dados do Player
func (p *Player) Update() {
	rl.DrawRectangleV(p.position, p.texture, rl.SkyBlue)
}
