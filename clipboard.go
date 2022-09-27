// Package screenpro capture the screen on Windows, Linux, and Mac with cool background.
package screenpro

import (
	"golang.design/x/clipboard"
	"image"
)

func SaveClipboard(img image.Image) {
	if err := clipboard.Init(); err != nil {
		panic(err)
	}
}
