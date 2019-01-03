package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Player é uma estrutura contendo as informações básicas do jogador
type Player struct {
	texture  rl.Texture2D
	position rl.Vector2
	speed    rl.Vector2
}

//NewPlayer retorna uma nova instancia de player
func NewPlayer() (p Player) {
	img := rl.LoadImage("resources/Sprite-0001.png")
	p.texture = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	p.position = rl.NewVector2(float32(0), float32(0))
	p.speed = rl.NewVector2(float32(2), float32(2))
	return
}

// Update atualiza os dados do Player
func (p *Player) Update() {
	rl.DrawTexture(p.texture, int32(p.position.X), int32(p.position.Y), rl.RayWhite)
}
