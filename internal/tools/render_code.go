package tools

import (
	"strings"
	"text/template"
)

// RenderCode takes a template string and any given data to render it, and
// returns the resulting string and possible error.
func RenderCode(text string, data any) (string, error) {
	tmpl, err := template.New("").Parse(text)
	if err != nil {
		return "", err
	}
	buf := &strings.Builder{}
	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
