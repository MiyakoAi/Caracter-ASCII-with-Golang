package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	stdDraw "image/draw"

	"golang.org/x/image/draw"
)

func main() {
	f, err := os.Open("Hoshimachi-Suisei.png") // file foto
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	srcImg, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// Resize
	const targetW = 100 // lebar
	bounds := srcImg.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	targetH := (targetW * h) / w / 2 // tinggi font

	dst := image.NewRGBA(image.Rect(0, 0, targetW, targetH))
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), srcImg, bounds, stdDraw.Over, nil)

	// Karakter ASCII
	asciiChars := "@%#*+=-:. "

	// Loop piksel
	for y := 0; y < targetH; y++ {
		for x := 0; x < targetW; x++ {
			r, g, b, _ := dst.At(x, y).RGBA()
			brightness := uint8((r*299 + g*587 + b*114 + 500) / 1000 >> 8)
			index := int(brightness) * (len(asciiChars) - 1) / 255
			fmt.Printf("%c", asciiChars[index])
		}
		fmt.Println()
	}
}
