package playground

import (
	"github.com/monceaux/fyne/v2/driver/software"
	"github.com/monceaux/fyne/v2/test"
)

// NewSoftwareCanvas creates a new canvas in memory that can render without hardware support
func NewSoftwareCanvas() test.WindowlessCanvas {
	return software.NewCanvas()
}
