package refyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"

	"github.com/fyne-io/refyne/internal/guidefs"
)

// Context defines a graphical builder context that contains metadata and theme information.
type Context = guidefs.Context

type context struct {
	meta map[fyne.CanvasObject]map[string]string
}

// DefaultContext returns a simple context with an empty metadata map that will
// defer to the default theme for theming information.
func DefaultContext() Context {
	return &context{meta: make(map[fyne.CanvasObject]map[string]string)}
}

func (c *context) Metadata() map[fyne.CanvasObject]map[string]string {
	return c.meta
}

func (c *context) Theme() fyne.Theme {
	return theme.DefaultTheme()
}
