package main

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"log"
	"runtime"
)

func main() {
	img := generateDumbImage()
	err := imaging.Save(img, "dumb.png")
	if err != nil {
		log.Fatalf("Could not save the image: %v", err)
	}
}

func generateDumbImage() image.Image {
	img := imaging.New(65536, 65536, color.Black)
	//pix := uint32(math.MaxUint32)
	var pix color.NRGBA
	bounds := img.Bounds()
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			img.Set(x, y, pix)
			incrementColour(&pix)
			//c := uint32ToColour(pix)
			//img.Set(x, y, c)
			//pix--
		}
	}

	return img
}

func incrementColour(pix *color.NRGBA) {
	if pix.R < 255 {
		pix.R++
		return
	}
	pix.R = 0
	if pix.G < 255 {
		pix.G++
		return
	}
	pix.G = 0
	if pix.B < 255 {
		pix.B++
		return
	}
	runtime.GC() // We OOM without this
	pix.B = 0
	if pix.A < 255 {
		pix.A++
		return
	}
	pix.A = 0
}

//func uint32ToColour(n uint32) color.NRGBA {
//	return color.NRGBA{
//		A: uint8(n >> 24),
//		R: uint8(n >> 16),
//		G: uint8(n >> 8),
//		B: uint8(n),
//	}
//}
