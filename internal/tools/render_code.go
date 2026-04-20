package tools

import (
	"strings"
	"text/template"
)

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
