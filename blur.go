// Package screenpro capture the screen on Windows, Linux, and Mac with cool background.
package screenpro

import (
	"github.com/anthonynsimon/bild/blur"
	"github.com/fogleman/gg"
	"image"
	"os"
	"path/filepath"
	"strings"
)

func Blur(path string) image.Image {
	reader, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	ext := filepath.Ext(path)
	name := strings.TrimSuffix(filepath.Base(path), ext)

	inputImage, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}

	width := inputImage.Bounds().Dx()
	height := inputImage.Bounds().Dy()
	step := 75

	newWidth := width + step
	newHeight := height + step
	step = 10

	dc := gg.NewContext(newWidth, newHeight)
	dc.DrawRoundedRectangle(40, 40, float64(width-step), float64(height-step), 100)
	dc.SetRGBA(255, 255, 255, 1)
	dc.Fill()

	dropShadow := blur.Gaussian(dc.Image(), 400)
	dc = gg.NewContext(newWidth, newHeight)
	dc.DrawImage(dropShadow, -10, 25)

	dc.DrawImage(inputImage, 25, 25)
	dc.Fill()

	dc.SavePNG(name + "out" + ext)
	return dc.Image()
}
