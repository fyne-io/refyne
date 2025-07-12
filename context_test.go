package refyne

import (
	"image/color"
	"testing"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"

	"github.com/stretchr/testify/assert"
)

func TestDefaultContext(t *testing.T) {
	d := DefaultContext()
	assert.Equal(t, theme.DefaultTheme(), d.Theme())

	o := canvas.NewRectangle(color.Black)
	assert.Nil(t, d.Metadata()[o])
}
