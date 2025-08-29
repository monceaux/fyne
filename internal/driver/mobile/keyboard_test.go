package mobile

import (
	"testing"

	_ "github.com/monceaux/fyne/v2/test"
)

func TestDevice_HideVirtualKeyboard_BeforeRun(t *testing.T) {
	(&device{}).hideVirtualKeyboard() // should not crash!
}
