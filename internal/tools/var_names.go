package tools

import (
	"reflect"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
)

var VarNames = NewVarNameGen()

type VarNameGen struct {
	m map[fyne.CanvasObject]string
	c map[string]int
}

func NewVarNameGen() *VarNameGen {
	vn := &VarNameGen{}
	vn.Reset()
	return vn
}

func (vn *VarNameGen) Reset() {
	vn.m = make(map[fyne.CanvasObject]string)
	vn.c = make(map[string]int)
}

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
