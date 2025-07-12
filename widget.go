package refyne

import (
	"reflect"

	"fyne.io/fyne/v2"

	"github.com/fyne-io/refyne/internal/guidefs"
)

// CollectionClassList returns the list of supported collection widget classes.
// These can be used for passing to `CreateNew` or `EditorFor`.
func CollectionClassList() []string {
	return guidefs.CollectionNames
}

// ContainerClassList returns the list of supported container classes.
// These can be used for passing to `CreateNew` or `EditorFor`.
func ContainerClassList() []string {
	return guidefs.ContainerNames
}

// GraphicsClassList returns the list of supported graphics primitives classes.
// These can be used for passing to `CreateNew` or `EditorFor`.
func GraphicsClassList() []string {
	return guidefs.GraphicsNames
}

// WidgetClassList returns the list of supported widget classes.
// These can be used for passing to `CreateNew` or `EditorFor`.
func WidgetClassList() []string {
	return guidefs.WidgetNames
}

// DropZonesForObject returns the children of a container that can be used as drag and drop target zones
func DropZonesForObject(o fyne.CanvasObject) []fyne.CanvasObject {
	class := reflect.TypeOf(o).String()
	info := guidefs.Lookup(class)

	if !info.IsContainer() {
		return nil
	}

	return info.Children(o)
}
