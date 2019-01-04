package main

import (
	"image/png"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func loadTileset(name string) (t rl.Texture2D) {
	r, err := os.Open(name)
	if err != nil {
		rl.TraceLog(rl.LogError, err.Error())
	}
	defer r.Close()

	img, err := png.Decode(r)
	if err != nil {
		rl.TraceLog(rl.LogError, err.Error())
	}

	im := rl.NewImageFromImage(img)
	t = rl.LoadTextureFromImage(im)

	rl.UnloadImage(im)
	return
}
