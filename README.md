# Pixel art Generator

A Go package for generating 16×16 pixel art deterministically from a seed string. Includes multiple templates (person, dog, cat, tree, star), color palettes, accent overlays, and border styles for endless variety.

## Features

* Deterministic generation using SHA-256 of seed.
* 5 shape templates + customizable additions.
* 16 rich color palettes (BG, primary, accent, border).
* Optional accent overlay shapes.
* Four border styles: rectangle, circle, dots, checker.
* Horizontal flip based on seed bits.

## Installation

```bash
go get github.com/flynnfc/artHash
```

## Usage

```go
package main

import (
    "image/png"
    "os"

    "github.com/flynnfc/artHash"
)

func main() {
    img := artHash.Generate("example-seed")
    f, _ := os.Create("example.png")
    defer f.Close()
    png.Encode(f, img)
}
```

## Examples

Upload your generated images here to illustrate different seeds:

![Example 1](examples/seed1.png)
![Example 2](examples/seed2.png)

## Contributing

Feel free to open issues or PRs to add more templates, palettes, or features.

## License

MIT © Flynnfc
