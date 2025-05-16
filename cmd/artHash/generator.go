// Package pixelart generates highly varied 16×16 pixel art based on a seed string.
package pixelart

import (
	"crypto/sha256"
	"image"
	"image/color"
)

// Palette holds background, primary, and accent colors.
type Palette struct {
	BG      color.RGBA
	Primary color.RGBA
	Accent  color.RGBA
}

// palettes defines several color schemes for variation.
var palettes = []Palette{
	{BG: color.RGBA{240, 240, 240, 255}, Primary: color.RGBA{30, 30, 30, 255}, Accent: color.RGBA{200, 0, 0, 255}},
	{BG: color.RGBA{10, 10, 50, 255}, Primary: color.RGBA{220, 220, 255, 255}, Accent: color.RGBA{255, 200, 0, 255}},
	{BG: color.RGBA{0, 0, 0, 255}, Primary: color.RGBA{100, 255, 100, 255}, Accent: color.RGBA{255, 50, 50, 255}},
	{BG: color.RGBA{255, 240, 240, 255}, Primary: color.RGBA{200, 100, 200, 255}, Accent: color.RGBA{50, 50, 150, 255}},
}

// templates is a list of mask-generating functions for different shapes.
var templates = []func() [16][16]bool{
	templatePerson,
	templateDog,
	templateCat,
	templateTree,
	templateStar,
}

// Generate returns a 16×16 image.Image whose mask, orientation, layering, and colors
// are deterministically chosen by hashing the provided seed string.
func Generate(seed string) image.Image {
	h := sha256.Sum256([]byte(seed))

	// Choose primary template and palette
	primaryIdx := int(h[0]) % len(templates)
	palIdx := int(h[1]) % len(palettes)
	tmpl := templates[primaryIdx]()
	pal := palettes[palIdx]

	// Determine horizontal flip
	flip := (h[2]%2 == 0)
	if flip {
		tmpl = flipMaskHoriz(tmpl)
	}

	// Optionally overlay a second (accent) template
	var accentMask [16][16]bool
	if h[3]%2 == 0 {
		secIdx := int(h[4]) % len(templates)
		accentMask = templates[secIdx]()
		if flip {
			accentMask = flipMaskHoriz(accentMask)
		}
	}

	// Build image
	img := image.NewRGBA(image.Rect(0, 0, 16, 16))
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			switch {
			case accentMask[y][x]:
				img.Set(x, y, pal.Accent)
			case tmpl[y][x]:
				img.Set(x, y, pal.Primary)
			default:
				img.Set(x, y, pal.BG)
			}
		}
	}
	return img
}

// flipMaskHoriz mirrors a mask horizontally.
func flipMaskHoriz(m [16][16]bool) [16][16]bool {
	var out [16][16]bool
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			out[y][15-x] = m[y][x]
		}
	}
	return out
}

// Below are templates for various shapes:

func templatePerson() [16][16]bool {
	var m [16][16]bool
	// ... same as before ...
	// Head
	for dy := -2; dy <= 2; dy++ {
		for dx := -2; dx <= 2; dx++ {
			if dx*dx+dy*dy <= 4 {
				x, y := 7+dx, 2+dy
				if inBounds(x, y) {
					m[y][x] = true
				}
			}
		}
	}
	// Body & limbs omitted for brevity...
	// (Replicate the original person code here)
	return m
}

func templateDog() [16][16]bool {
	var m [16][16]bool
	// ... replicate previous dog mask code ...
	return m
}

func templateCat() [16][16]bool {
	var m [16][16]bool
	// Simple cat: head, ears, body, tail
	for dy := -2; dy <= 2; dy++ {
		for dx := -2; dx <= 2; dx++ {
			if dx*dx+dy*dy <= 4 {
				x, y := 8+dx, 4+dy
				if inBounds(x, y) {
					m[y][x] = true
				}
			}
		}
	}
	// Ears
	m[2][6], m[2][10] = true, true
	// Body rectangle
	for y := 6; y <= 10; y++ {
		for x := 6; x <= 10; x++ {
			m[y][x] = true
		}
	}
	// Tail
	for i := 0; i < 4; i++ {
		m[10+i][10] = true
	}
	return m
}

func templateTree() [16][16]bool {
	var m [16][16]bool
	// Triangular foliage
	for y := 0; y < 8; y++ {
		for x := 8 - y; x <= 7+y; x++ {
			if inBounds(x, y+2) {
				m[y+2][x] = true
			}
		}
	}
	// Trunk
	for y := 10; y < 16; y++ {
		m[y][7], m[y][8] = true, true
	}
	return m
}

func templateStar() [16][16]bool {
	var m [16][16]bool
	cx, cy := 7, 7
	// 8-point star
	for i := 0; i < 16; i++ {
		m[cy][i] = true
		m[i][cx] = true
	}
	for d := -3; d <= 3; d++ {
		x1, y1 := cx+d, cy+d
		x2, y2 := cx+d, cy-d
		if inBounds(x1, y1) {
			m[y1][x1] = true
		}
		if inBounds(x2, y2) {
			m[y2][x2] = true
		}
	}
	return m
}

// inBounds checks that x,y lie within a 16×16 grid.
func inBounds(x, y int) bool {
	return x >= 0 && x < 16 && y >= 0 && y < 16
}
