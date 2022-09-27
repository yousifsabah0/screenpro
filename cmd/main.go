package main

import (
	"github.com/yousifsabah0/screenpro"

	"golang.design/x/mainthread"
)

func main() {
	// Run on main thread.
	mainthread.Init(screenpro.ScreenPro)
}
