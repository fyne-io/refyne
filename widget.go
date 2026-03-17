package refyne

import (
	"fyne.io/fyne/v2"

	"github.com/fyne-io/refyne/internal/guidefs"
)

type WidgetInfo = guidefs.WidgetInfo

// CollectionClassList returns the list of supported collection widget classes.
// These can be used for passing to `CreateNew` or `EditorFor`.
func CollectionClassList() []string {
	return guidefs.CollectionNames
}

// RegisterCollection allows a 3rd party collection widget to be added to those recognised.
// It is important that the Name field is populated with the typed name (i.e. *myPkg.MyWidget)
// and your package should be returned in the list from Packages() as well.
func RegisterCollection(info WidgetInfo) {
	guidefs.InitOnce()

	guidefs.CollectionNames = append(guidefs.CollectionNames, info.Name)
	guidefs.Collections[info.Name] = info
}

// ContainerClassList returns the list of supported container classes.
// These can be used for passing to `CreateNew` or `EditorFor`.
func ContainerClassList() []string {
	return guidefs.ContainerNames
}

// RegisterContainer allows a 3rd party container widget to be added to those recognised.
// It is important that the Name field is populated with the typed name (i.e. *myPkg.MyWidget)
// and your package should be returned in the list from Packages() as well.
func RegisterContainer(info WidgetInfo) {
	guidefs.InitOnce()

	guidefs.ContainerNames = append(guidefs.ContainerNames, info.Name)
	guidefs.Containers[info.Name] = info
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

// RegisterWidget allows a 3rd party Fyne widget to be added to those recognised.
// It is important that the Name field is populated with the typed name (i.e. *myPkg.MyWidget)
// and your package should be returned in the list from Packages() as well.
func RegisterWidget(info WidgetInfo) {
	guidefs.InitOnce()

	guidefs.WidgetNames = append(guidefs.WidgetNames, info.Name)
	guidefs.Widgets[info.Name] = info
}

// DropZonesForObject returns the children of a container that can be used as drag and drop target zones
func DropZonesForObject(o fyne.CanvasObject) []fyne.CanvasObject {
	class := guidefs.TypeName(o)
	info := guidefs.Lookup(class)

	if !info.IsContainer() {
		return nil
	}

	return info.Children(o)
}
