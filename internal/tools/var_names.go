package tools

import (
	"reflect"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
)

// VarNames is a shared generator instance for variable names from CanvasObjects.
var VarNames = NewVarNameGen()

// VarNameGen is a variable name generator that groups and counts objects by type
// and remembers the names returned for a given object.
type VarNameGen struct {
	m map[fyne.CanvasObject]string
	c map[string]int
}

// NewVarNameGen returns a new instance ready to use.
func NewVarNameGen() *VarNameGen {
	vn := &VarNameGen{}
	vn.Reset()
	return vn
}

// Reset allocates a new and empty internal state.
func (vn *VarNameGen) Reset() {
	vn.m = make(map[fyne.CanvasObject]string)
	vn.c = make(map[string]int)
}

// Get returns the variable name for a given object, and remembers previous
// values to maintain stable results.
func (vn *VarNameGen) Get(obj fyne.CanvasObject) string {
	if name, found := vn.m[obj]; found {
		return name
	}

	s := reflect.TypeOf(obj).String()
	i := strings.Index(s, ".")
	t := s[i+1:]

	vn.c[t]++
	vn.m[obj] = t + strconv.Itoa(vn.c[t])

	return vn.m[obj]
}
