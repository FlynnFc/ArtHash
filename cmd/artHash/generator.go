package arthash

import (
	"crypto/sha256"
	"image"
	"image/color"
)

type Palette struct {
	BG      color.RGBA
	Primary color.RGBA
	Accent  color.RGBA
	Border  color.RGBA
}
type Size int

const (
	Small Size = iota
	Medium
	Large
)

var palettes = []Palette{
	{BG: color.RGBA{240, 240, 240, 255}, Primary: color.RGBA{30, 30, 30, 255}, Accent: color.RGBA{200, 0, 0, 255}, Border: color.RGBA{0, 0, 0, 255}},
	{BG: color.RGBA{10, 10, 50, 255}, Primary: color.RGBA{220, 220, 255, 255}, Accent: color.RGBA{255, 200, 0, 255}, Border: color.RGBA{50, 50, 100, 255}},
	{BG: color.RGBA{0, 0, 0, 255}, Primary: color.RGBA{100, 255, 100, 255}, Accent: color.RGBA{255, 50, 50, 255}, Border: color.RGBA{255, 255, 255, 255}},
	{BG: color.RGBA{255, 240, 240, 255}, Primary: color.RGBA{200, 100, 200, 255}, Accent: color.RGBA{50, 50, 150, 255}, Border: color.RGBA{150, 50, 150, 255}},
	{BG: color.RGBA{230, 230, 250, 255}, Primary: color.RGBA{75, 0, 130, 255}, Accent: color.RGBA{138, 43, 226, 255}, Border: color.RGBA{72, 61, 139, 255}},
	{BG: color.RGBA{255, 250, 205, 255}, Primary: color.RGBA{255, 215, 0, 255}, Accent: color.RGBA{218, 165, 32, 255}, Border: color.RGBA{184, 134, 11, 255}},
	{BG: color.RGBA{224, 255, 255, 255}, Primary: color.RGBA{0, 206, 209, 255}, Accent: color.RGBA{72, 209, 204, 255}, Border: color.RGBA{95, 158, 160, 255}},
	{BG: color.RGBA{255, 228, 225, 255}, Primary: color.RGBA{255, 105, 180, 255}, Accent: color.RGBA{255, 20, 147, 255}, Border: color.RGBA{199, 21, 133, 255}},
	{BG: color.RGBA{245, 245, 220, 255}, Primary: color.RGBA{160, 82, 45, 255}, Accent: color.RGBA{210, 105, 30, 255}, Border: color.RGBA{139, 69, 19, 255}},
	{BG: color.RGBA{224, 238, 224, 255}, Primary: color.RGBA{34, 139, 34, 255}, Accent: color.RGBA{0, 100, 0, 255}, Border: color.RGBA{85, 107, 47, 255}},
	{BG: color.RGBA{245, 222, 179, 255}, Primary: color.RGBA{210, 180, 140, 255}, Accent: color.RGBA{222, 184, 135, 255}, Border: color.RGBA{160, 82, 45, 255}},
	{BG: color.RGBA{176, 196, 222, 255}, Primary: color.RGBA{65, 105, 225, 255}, Accent: color.RGBA{25, 25, 112, 255}, Border: color.RGBA{0, 0, 139, 255}},
	{BG: color.RGBA{255, 228, 181, 255}, Primary: color.RGBA{255, 165, 0, 255}, Accent: color.RGBA{255, 140, 0, 255}, Border: color.RGBA{255, 69, 0, 255}},
	{BG: color.RGBA{240, 255, 255, 255}, Primary: color.RGBA{32, 178, 170, 255}, Accent: color.RGBA{0, 139, 139, 255}, Border: color.RGBA{47, 79, 79, 255}},
	{BG: color.RGBA{255, 239, 213, 255}, Primary: color.RGBA{244, 164, 96, 255}, Accent: color.RGBA{210, 105, 30, 255}, Border: color.RGBA{160, 82, 45, 255}},
	{BG: color.RGBA{253, 245, 230, 255}, Primary: color.RGBA{233, 150, 122, 255}, Accent: color.RGBA{250, 128, 114, 255}, Border: color.RGBA{205, 92, 92, 255}},
}

var templates = []func() [16][16]bool{
	templatePerson, templateDog, templateCat, templateTree, templateStar,
}

var borderShapes = []func() [16][16]bool{
	borderRect, borderCircle, borderDots, borderChecker,
}

func Generate(seed string, size Size) image.Image {
	base := generateBase(seed)

	// determine scaling factor
	scale := 1
	switch size {
	case Medium:
		scale = 2
	case Large:
		scale = 4
	}
	if scale == 1 {
		return base
	}

	// upscale with nearest-neighbor
	dst := image.NewRGBA(image.Rect(0, 0, 16*scale, 16*scale))
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			c := base.At(x, y)
			for dy := 0; dy < scale; dy++ {
				for dx := 0; dx < scale; dx++ {
					dst.Set(x*scale+dx, y*scale+dy, c)
				}
			}
		}
	}
	return dst
}

func generateBase(seed string) *image.RGBA {
	h := sha256.Sum256([]byte(seed))

	// pick template & palette
	tmpl := templates[int(h[0])%len(templates)]()
	pal := palettes[int(h[1])%len(palettes)]

	// optional horizontal flip
	if h[2]%2 == 0 {
		tmpl = flipMaskHoriz(tmpl)
	}

	// optional border
	addBorder := h[5]%2 == 0
	var borderMask [16][16]bool
	if addBorder {
		borderMask = borderShapes[int(h[6])%len(borderShapes)]()
	}

	img := image.NewRGBA(image.Rect(0, 0, 16, 16))
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			switch {
			case addBorder && borderMask[y][x]:
				img.Set(x, y, pal.Border)
			case tmpl[y][x]:
				img.Set(x, y, pal.Primary)
			default:
				img.Set(x, y, pal.BG)
			}
		}
	}
	return img
}

func flipMaskHoriz(m [16][16]bool) [16][16]bool {
	var out [16][16]bool
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			out[y][15-x] = m[y][x]
		}
	}
	return out
}

func borderRect() [16][16]bool {
	var m [16][16]bool
	for i := 0; i < 16; i++ {
		m[0][i], m[15][i], m[i][0], m[i][15] = true, true, true, true
	}
	return m
}

func borderCircle() [16][16]bool {
	var m [16][16]bool
	cx, cy, r := 7.5, 7.5, 7.0
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			d := (float64(x)-cx)*(float64(x)-cx) + (float64(y)-cy)*(float64(y)-cy)
			if d >= (r-1)*(r-1) && d <= r*r {
				m[y][x] = true
			}
		}
	}
	return m
}

func borderDots() [16][16]bool {
	var m [16][16]bool
	for i := 0; i < 16; i += 2 {
		m[0][i], m[15][i] = true, true
		m[i][0], m[i][15] = true, true
	}
	return m
}

// borderChecker draws a checker-pattern frame.
func borderChecker() [16][16]bool {
	var m [16][16]bool
	for i := 0; i < 16; i += 2 {
		m[0][i], m[15][i] = true, true
		m[i][0], m[i][15] = true, true
	}
	return m
}

func templatePerson() [16][16]bool {
	var m [16][16]bool
	// head (circle radius 2 at (7,2))
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
	// body
	for y := 5; y <= 10; y++ {
		m[y][7] = true
	}
	// arms
	for i := 0; i <= 4; i++ {
		dy, dx := i/2, i
		if inBounds(7+dx, 6+dy) {
			m[6+dy][7+dx] = true
			m[6+dy][7-dx] = true
		}
	}
	// legs
	for i := 0; i <= 5; i++ {
		if inBounds(7+i, 10+i) {
			m[10+i][7+i] = true
			m[10+i][7-i] = true
		}
	}
	return m
}

func templateDog() [16][16]bool {
	var m [16][16]bool
	// body
	for y := 7; y <= 11; y++ {
		for x := 3; x <= 12; x++ {
			m[y][x] = true
		}
	}
	// legs
	for y := 11; y < 16; y++ {
		for _, x := range []int{4, 5, 10, 11, 12, 13} {
			if inBounds(x, y) {
				m[y][x] = true
			}
		}
	}
	// head
	for dy := -2; dy <= 2; dy++ {
		for dx := -2; dx <= 2; dx++ {
			if dx*dx+dy*dy <= 4 {
				x, y := 2+dx, 9+dy
				if inBounds(x, y) {
					m[y][x] = true
				}
			}
		}
	}
	// tail
	for i := 0; i < 4; i++ {
		x, y := 12+i, 8-i
		if inBounds(x, y) {
			m[y][x] = true
		}
	}
	return m
}

func templateCat() [16][16]bool {
	var m [16][16]bool
	// head
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
	// ears
	m[2][6], m[2][10] = true, true
	// body
	for y := 6; y <= 10; y++ {
		for x := 6; x <= 10; x++ {
			m[y][x] = true
		}
	}
	// tail
	for i := 0; i < 4; i++ {
		m[10+i][10] = true
	}
	return m
}

func templateTree() [16][16]bool {
	var m [16][16]bool
	// foliage (triangle)
	for y := 0; y < 8; y++ {
		for x := 8 - y; x <= 7+y; x++ {
			if inBounds(x, y+2) {
				m[y+2][x] = true
			}
		}
	}
	// trunk
	for y := 10; y < 16; y++ {
		m[y][7], m[y][8] = true, true
	}
	return m
}

func templateStar() [16][16]bool {
	var m [16][16]bool
	cx, cy := 7, 7
	// cross
	for i := 0; i < 16; i++ {
		m[cy][i] = true
		m[i][cx] = true
	}
	// diagonals
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

func inBounds(x, y int) bool {
	return x >= 0 && x < 16 && y >= 0 && y < 16
}
