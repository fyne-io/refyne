package tools

import (
	"testing"

	"fyne.io/fyne/v2/widget"

	"github.com/stretchr/testify/assert"
)

func TestVarNames(t *testing.T) {
	list := &widget.List{}
	assert.Equal(t, "list1", VarNames.Get(list))
	assert.Equal(t, "list2", VarNames.Get(&widget.List{}))
	assert.Equal(t, "button1", VarNames.Get(&widget.Button{}))
	assert.Equal(t, "list1", VarNames.Get(list))

	VarNames.Reset()
	assert.Equal(t, "list1", VarNames.Get(&widget.List{}))
}
