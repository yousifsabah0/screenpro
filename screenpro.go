// Package screenpro capture the screen on Windows, Linux, and Mac with cool background.
package screenpro

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"golang.design/x/hotkey"
	"image"
	"image/jpeg"
	"os"
)

// ScreenPro main func to start the program.
func ScreenPro() {
	// Shortcut is: Ctrl + Alt + P.
	key := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyP)

	// Register the shortcut
	if err := key.Register(); err != nil {
		panic(err)
	}

	// When user presses a key run the Capture() func.
	<-key.Keydown()
	Capture()
}

// Capture captures the screen and adds a blur effect.
func Capture() image.Image {
	var lastResult image.Image
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Active display not found")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)

		img, err := screenshot.CaptureDisplay(i)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.jpeg", i, bounds.Dx(), bounds.Dy())
		save(img, fileName)
		lastResult = Blur(fileName)
	}

	return lastResult
}

func save(img *image.RGBA, filePath string) {
	file, err := os.CreateTemp(filePath, "screenpro-temp")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	jpeg.Encode(file, img, nil)
}
