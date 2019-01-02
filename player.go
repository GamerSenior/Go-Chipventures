package main

import rl "github.com/gen2brain/raylib-go/raylib"

type player struct {
	texture rl.Texture2D
}

func newPlayer() player {
	p := player{}
	return p
}
