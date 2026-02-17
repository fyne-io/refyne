package refyne

import (
	"reflect"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"github.com/fyne-io/refyne/internal/guidefs"
)

// CreateNew returns a new instance of the given widget type
func CreateNew(name string, d Context) fyne.CanvasObject {
	guidefs.InitOnce()

	if match := guidefs.Lookup(name); match != nil {
		return match.Create(d)
	}

	return nil
}

// EditorFor returns an array of FormItems for editing, taking the widget, properties, callback to refresh the form items,
// and an optional callback that fires after changes to the widget.
func EditorFor(o fyne.CanvasObject, d Context, refresh func([]*widget.FormItem), onchanged func()) []*widget.FormItem {
	guidefs.InitOnce()

	_, clazz := getTypeOf(o)

	if onchanged == nil {
		onchanged = func() {}
	}

	appendManualItems := func(items []*widget.FormItem) []*widget.FormItem {
		parent := ContainerOf(o, d)
		if c, ok := parent.(*fyne.Container); ok {
			without := c.Layout == nil
			if !without {
				if l, ok := c.Layout.(interface{ WithoutLayout() bool }); ok && l.WithoutLayout() {
					without = true
				}
			}
			if without {
				return append(items, manualLayoutItems(o, d.Metadata()[o])...)
			}
		}

		return items
	}

	var items []*widget.FormItem
	if match := guidefs.Lookup(clazz); match != nil {
		items = match.Edit(o, d, func(items []*widget.FormItem) {
			items = appendManualItems(items)

			refresh(items)
		}, onchanged)
	}

	items = appendManualItems(items)
	return items
}

// GoStringFor generates the Go code for the given widget
func GoStringFor(o fyne.CanvasObject, d Context, defs map[string]string) string {
	guidefs.InitOnce()

	name := reflect.TypeOf(o).String()

	if match := guidefs.Lookup(name); match != nil {
		return match.Gostring(o, d, defs)
	}

	return ""
}

func getTypeOf(o fyne.CanvasObject) (string, string) {
	class := reflect.TypeOf(o).String()
	name := NameOf(o)

	return name, class
}

// NameOf returns the name for a given object
func NameOf(o fyne.CanvasObject) string {
	typeName := reflect.TypeOf(o).Elem().Name()
	l := reflect.ValueOf(o).Elem()
	if typeName == "Entry" {
		if l.FieldByName("Password").Bool() {
			typeName = "PasswordEntry"
		} else if l.FieldByName("MultiLine").Bool() {
			typeName = "MultiLineEntry"
		}
	}

	return typeName
}

func manualLayoutItems(o fyne.CanvasObject, props map[string]string) []*widget.FormItem {
	header := widget.NewFormItem("Layout", widget.NewLabel(""))
	header.HintText = "(manual, static)"

	return []*widget.FormItem{
		header,
		widget.NewFormItem("X", float32Entry(o.Position().X, "x", props, func(f float32) {
			o.Move(fyne.NewPos(f, o.Position().Y))
		})),
		widget.NewFormItem("Y", float32Entry(o.Position().Y, "y", props, func(f float32) {
			o.Move(fyne.NewPos(o.Position().X, f))
		})),
		widget.NewFormItem("Width", float32Entry(o.Size().Width, "width", props, func(f float32) {
			o.Resize(fyne.NewSize(f, o.Size().Height))
		})),
		widget.NewFormItem("Height", float32Entry(o.Size().Height, "height", props, func(f float32) {
			o.Resize(fyne.NewSize(o.Size().Width, f))
		})),
	}
}

func floatEntry(in float64, id string, props map[string]string, out func(float64)) *widget.Entry {
	val := binding.BindFloat(&in)
	val.AddListener(binding.NewDataListener(func() {
		f, _ := val.Get()
		out(f)
		props[id] = strconv.FormatFloat(f, 'f', -2, 64)
	}))
	return widget.NewEntryWithData(binding.FloatToString(val))
}

func float32Entry(in float32, id string, props map[string]string, out func(float32)) *widget.Entry {
	return floatEntry(float64(in), id, props, func(f float64) {
		out(float32(f))
	})
}
