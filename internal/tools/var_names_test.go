package tools

import (
	"testing"

	"fyne.io/fyne/v2/widget"

	"github.com/stretchr/testify/assert"
)

func TestVarNames(t *testing.T) {
	list := &widget.List{}
	assert.Equal(t, "List1", VarNames.Get(list))
	assert.Equal(t, "List2", VarNames.Get(&widget.List{}))
	assert.Equal(t, "Button1", VarNames.Get(&widget.Button{}))
	assert.Equal(t, "List1", VarNames.Get(list))

	VarNames.Reset()
	assert.Equal(t, "List1", VarNames.Get(&widget.List{}))
}
