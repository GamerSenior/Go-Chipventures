package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Player é uma estrutura contendo as informações básicas do jogador
type Player struct {
	texture rl.Texture2D
}

//NewPlayer retorna uma nova instancia de player
func NewPlayer() Player {
	p := Player{}
	return p
}
