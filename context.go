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
	attr map[fyne.CanvasObject]map[string]string
	root fyne.CanvasObject
}

// DefaultContext returns a simple context with an empty metadata map that will
// defer to the default theme for theming information.
func DefaultContext() Context {
	return &context{
		meta: make(map[fyne.CanvasObject]map[string]string),
		attr: make(map[fyne.CanvasObject]map[string]string),
	}
}

func (c *context) Metadata() map[fyne.CanvasObject]map[string]string {
	return c.meta
}

func (c *context) Attrs() map[fyne.CanvasObject]map[string]string {
	return c.attr
}

func (c *context) Root() fyne.CanvasObject {
	return c.root
}

func (c *context) Theme() fyne.Theme {
	return theme.DefaultTheme()
}

// ContainerOf returns the parent of the given CanvasObject, in the specified Context.
// The returned object will be in the tree descended from `c.Root()`, or nil.
func ContainerOf(obj fyne.CanvasObject, c Context) fyne.CanvasObject {
	return containerOf(obj, c.Root())
}

func containerOf(obj fyne.CanvasObject, root fyne.CanvasObject) fyne.CanvasObject {
	switch c := root.(type) {
	case *fyne.Container:
		for _, w := range c.Objects {
			if w == obj {
				return root
			}

			parent := containerOf(obj, w)
			if parent != nil {
				return parent
			}
		}

	case fyne.Widget:
		drops := DropZonesForObject(root)

		for _, child := range drops {
			if child == obj {
				return root
			}

			parent := containerOf(obj, child)
			if parent != nil {
				return parent
			}
		}
	}

	return nil
}
