package main

import (
	"image/png"
	"os"

	arthash "github.com/flynnfc/monster-hash/cmd/artHash"
)

func main() {
	img := arthash.Generate("HAZNOODLi")
	f, _ := os.Create("out2.png")
	defer f.Close()
	png.Encode(f, img)
}
