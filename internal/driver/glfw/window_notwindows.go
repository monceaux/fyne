//go:build !windows

package glfw

import (
	"github.com/monceaux/fyne/v2"
	"github.com/monceaux/fyne/v2/internal/scale"
)

func (w *window) setDarkMode() {
}

func (w *window) computeCanvasSize(width, height int) fyne.Size {
	return fyne.NewSize(scale.ToFyneCoordinate(w.canvas, width), scale.ToFyneCoordinate(w.canvas, height))
}
