package main

import (
	"github.com/getlantern/systray"
	"github.com/robertarles/lidless/internal/ui"
)

func main() {
	systray.Run(ui.OnReady, ui.OnExit)
}
