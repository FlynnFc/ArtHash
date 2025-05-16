package main

import (
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"time"

	arthash "github.com/flynnfc/monster-hash/cmd/artHash"
)

func randString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	// Seed the RNG
	rand.Seed(time.Now().UnixNano())

	// Ensure output directory exists
	outDir := "out"
	if err := os.MkdirAll(outDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create output dir: %v\n", err)
		os.Exit(1)
	}

	// Number of variations to generate
	n := 28
	for i := 0; i < n; i++ {
		seed := randString(8)
		img := arthash.Generate(seed, arthash.Large)

		// Sanitize filename to avoid special chars
		filename := fmt.Sprintf("%s/%s.png", outDir, seed)

		// Create file
		f, err := os.Create(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file %s: %v\n", filename, err)
			continue
		}
		// Encode to PNG and close
		if err := png.Encode(f, img); err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding PNG for %s: %v\n", seed, err)
		}
		f.Close()
		fmt.Printf("Saved %s.png\n", seed)
	}
}
