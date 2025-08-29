package widget_test

import (
	"testing"

	"github.com/monceaux/fyne/v2"
	"github.com/monceaux/fyne/v2/canvas"
	"github.com/monceaux/fyne/v2/container"
	"github.com/monceaux/fyne/v2/test"
	"github.com/monceaux/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestNewPasswordEntry(t *testing.T) {
	p := widget.NewPasswordEntry()
	p.Text = "visible"
	r := test.TempWidgetRenderer(t, p)

	cont := r.Objects()[2].(*container.Scroll).Content.(fyne.Widget)
	r = test.TempWidgetRenderer(t, cont)
	rich := r.Objects()[1].(*widget.RichText)
	r = test.TempWidgetRenderer(t, rich)

	assert.Equal(t, "•••••••", r.Objects()[0].(*canvas.Text).Text)
}
