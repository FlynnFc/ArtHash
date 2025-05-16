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

![UdJYmTu1](https://github.com/user-attachments/assets/eb1a965b-8269-4e09-9918-bccb9e494224)
![UAXRcbGH](https://github.com/user-attachments/assets/f9087ea7-3f0e-469d-9a22-0905d8d52b2c)
![u7vecmid](https://github.com/user-attachments/assets/e111a552-2748-4d2b-83f2-6b4b0e6912e7)
![u1Xw5GMx](https://github.com/user-attachments/assets/cd9e08ae-ff96-4a5b-ae9e-23c926af2058)
![TXjRXopP](https://github.com/user-attachments/assets/1fde4501-4881-48e3-b5a8-8e57001695a9)
![shVknEQp](https://github.com/user-attachments/assets/5255144c-2f92-4154-8672-fbce3f2531bc)
![Ry646Zcr](https://github.com/user-attachments/assets/7b0fb8b5-505b-445f-b135-95f28e44fc01)
![rqR9WDE2](https://github.com/user-attachments/assets/f754fdd5-3cfc-44fe-913b-c0b3380753fa)
![rpGQU5Er](https://github.com/user-attachments/assets/8ffdfe5a-2c1c-4e9f-adb8-5c6938ada323)
![r77k7unL](https://github.com/user-attachments/assets/d5e5fe25-403e-4419-97f0-fa9642434268)
![qzRNnt6D](https://github.com/user-attachments/assets/d003c35c-32e5-406e-a6a9-372b2e067645)
![PhEVhxM6](https://github.com/user-attachments/assets/76b6ced2-0499-4078-a3e1-7029533f0003)
![nSyABAy6](https://github.com/user-attachments/assets/a14ea273-bdba-4d4f-9dfa-6eb46dcc90e7)
![nbwL3Evq](https://github.com/user-attachments/assets/77e88934-48fa-4f52-852f-2b8fdc8eb90e)
![mXEOvSfz](https://github.com/user-attachments/assets/9ecdfba6-3625-4240-9d38-a348d472e5d2)
![MGf55rEJ](https://github.com/user-attachments/assets/73fc0f57-b632-4b32-8a67-9ff8a38387fa)
![lUCqDSJE](https://github.com/user-attachments/assets/7ba5f26b-cd87-4fad-8f66-30722290494f)
![kQOBA8qV](https://github.com/user-attachments/assets/4e927634-7b97-4b72-b7dd-16bf64e21673)
![KNGHVmEF](https://github.com/user-attachments/assets/b18562ef-bb2b-4693-8ca6-bbddb26efa5c)
![k4eWdD5y](https://github.com/user-attachments/assets/c9a5d300-f541-4072-b811-8144e65ccc11)
![jB7twNkh](https://github.com/user-attachments/assets/abbdc3e7-adc5-4dae-be33-b95bad68cd23)
![j400G6x7](https://github.com/user-attachments/assets/8a6f7cf4-529c-4069-b126-8734c898cfd2)
![ipWTqiRI](https://github.com/user-attachments/assets/02c7c84e-8b5f-449f-a47d-5f8cbf6fb815)
![IfKqqtRd](https://github.com/user-attachments/assets/e5718a16-bc17-40f1-8c23-c8a45e4c9780)
![htVgemZT](https://github.com/user-attachments/assets/ecb86766-a608-43ef-8aa4-a5021531ba94)
![Ho3amaHp](https://github.com/user-attachments/assets/72eb5438-d227-49dc-92c1-8022532b6a38)
![gZLa8Qbw](https://github.com/user-attachments/assets/2c80c948-4c7f-4389-a94f-c7172e6d38f6)




## Contributing

Feel free to open issues or PRs to add more templates, palettes, or features.

## License

MIT © Flynnfc
