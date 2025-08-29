//go:build !debug && !release

package build

import "github.com/monceaux/fyne/v2"

// Mode is the application's build mode.
const Mode = fyne.BuildStandard
