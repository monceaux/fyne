package lang

import (
	"github.com/monceaux/fyne/v2/internal/driver/mobile/app"

	"github.com/jeandeaual/go-locale"
)

func initRuntime() {
	locale.SetRunOnJVM(app.RunOnJVM)
}
