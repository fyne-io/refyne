package guidefs

import (
	"regexp"

	"fyne.io/fyne/v2"
)

var attrRe = regexp.MustCompile(`^([\d\w_]+)\b`)

func setAttr(m map[fyne.CanvasObject][]string, o fyne.CanvasObject, key, val string) {
	attrs := m[o]
	for n, attr := range attrs {
		x := attrRe.FindStringSubmatch(attr)
		if len(x) > 1 && x[1] == key {
			attrs[n] = key + val
			return
		}
	}
	if len(attrs) > 0 && attrs[len(attrs)-1] == key+val {
		return
	}
	m[o] = append(attrs, key+val)
}

func delAttr(m map[fyne.CanvasObject][]string, o fyne.CanvasObject, key string) {
	attrs := make([]string, 0, len(m[o]))
	for _, attr := range m[o] {
		x := attrRe.FindStringSubmatch(attr)
		if len(x) > 1 && x[1] == key {
			continue
		}
		attrs = append(attrs, attr)
	}
	if len(attrs) < len(m[o]) {
		m[o] = attrs
	}
}
