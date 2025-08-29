//go:build !darwin || no_native_menus

package glfw

import "github.com/monceaux/fyne/v2"

func setupNativeMenu(_ *window, _ *fyne.MainMenu) {
	// no-op
}
