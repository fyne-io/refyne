package guidefs

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func TestSetAttr(t *testing.T) {
	attrs := make(map[fyne.CanvasObject][]string)
	obj := &widget.Label{}
	setAttr(attrs, obj, "Text", ` = "foo"`)
	setAttr(attrs, obj, "Text", ` = "bar"`)
	assert.Len(t, attrs, 1)
	setAttr(attrs, obj, "Something", ` = ""`)
	setAttr(attrs, obj, "Something", ` = ""`)
	assert.Len(t, attrs, 1)
}

func TestDelAttr(t *testing.T) {
	attrs := make(map[fyne.CanvasObject][]string)
	obj := &widget.Label{}
	attrs[obj] = append(attrs[obj], `SetText("")`)
	attrs[obj] = append(attrs[obj], `SetText("")`)
	delAttr(attrs, obj, "SetText")
	assert.Len(t, attrs[obj], 0)
}
