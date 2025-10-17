package guidefs

import (
	"errors"
	"fmt"
	"image/color"
	"path/filepath"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	// GraphicsNames is an array with the list of names of all the graphical primitives
	GraphicsNames []string

	// Graphics provides the info about the type of canvas object primitives
	Graphics map[string]WidgetInfo
)

func initGraphics() {
	Graphics = map[string]WidgetInfo{
		"*canvas.Arc": {
			Name: "Arc",
			Create: func(Context) fyne.CanvasObject {
				rect := canvas.NewArc(0, 90, 0.5, color.Black)
				rect.StrokeColor = color.Black
				return rect
			},
			Edit: func(obj fyne.CanvasObject, _ Context, _ func([]*widget.FormItem), onchanged func()) []*widget.FormItem {
				a := obj.(*canvas.Arc)
				return []*widget.FormItem{
					widget.NewFormItem("Start Angle", newIntSliderButton(float64(a.StartAngle), -360, 360, func(f float64) {
						a.StartAngle = float32(f)
						a.Refresh()
						onchanged()
					})),
					widget.NewFormItem("End Angle", newIntSliderButton(float64(a.EndAngle), -360, 360, func(f float64) {
						a.EndAngle = float32(f)
						a.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Fill", newColorButton(a.FillColor, func(c color.Color) {
						a.FillColor = c
						a.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Corner", newIntSliderButton(float64(a.CornerRadius), 0, 32, func(f float64) {
						a.CornerRadius = float32(f)
						a.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Cutout Ratio", newRatioSliderButton(float64(a.CutoutRatio), 0, 1, func(f float64) {
						a.CutoutRatio = float32(f)
						a.Refresh()
						onchanged()
					})),

					widget.NewFormItem("Stroke", newIntSliderButton(float64(a.StrokeWidth), 0, 32, func(f float64) {
						a.StrokeWidth = float32(f)
						a.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Color", newColorButton(a.StrokeColor, func(c color.Color) {
						a.StrokeColor = c
						a.Refresh()
						onchanged()
					})),
				}
			},
			Packages: func(_ fyne.CanvasObject, _ Context) []string {
				return []string{"canvas", "image/color"}
			},
		},
		"*canvas.Circle": {
			Name: "Circle",
			Create: func(Context) fyne.CanvasObject {
				rect := canvas.NewCircle(color.Black)
				rect.StrokeColor = color.Black
				return rect
			},
			Edit: func(obj fyne.CanvasObject, _ Context, _ func([]*widget.FormItem), onchanged func()) []*widget.FormItem {
				r := obj.(*canvas.Circle)
				return []*widget.FormItem{
					widget.NewFormItem("Fill", newColorButton(r.FillColor, func(c color.Color) {
						r.FillColor = c
						r.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Stroke", newIntSliderButton(float64(r.StrokeWidth), 0, 32, func(f float64) {
						r.StrokeWidth = float32(f)
						r.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Color", newColorButton(r.StrokeColor, func(c color.Color) {
						r.StrokeColor = c
						r.Refresh()
						onchanged()
					})),
				}
			},
			Packages: func(_ fyne.CanvasObject, _ Context) []string {
				return []string{"canvas", "image/color"}
			},
		},
		"*canvas.Image": initImageGraphic(),
		"*canvas.LinearGradient": {
			Name: "LinearGradient",
			Create: func(Context) fyne.CanvasObject {
				return &canvas.LinearGradient{StartColor: color.White}
			},
			Edit: func(obj fyne.CanvasObject, _ Context, _ func([]*widget.FormItem), onchanged func()) []*widget.FormItem {
				r := obj.(*canvas.LinearGradient)
				angleSlide := widget.NewSlider(0, 360)
				angleSlide.Step = 90
				angleSlide.OnChanged = func(f float64) {
					r.Angle = f
					r.Refresh()
					onchanged()
				}
				return []*widget.FormItem{
					widget.NewFormItem("Start", newColorButton(r.StartColor, func(c color.Color) {
						r.StartColor = c
						r.Refresh()
						onchanged()
					})),
					widget.NewFormItem("End", newColorButton(r.EndColor, func(c color.Color) {
						r.EndColor = c
						r.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Angle", angleSlide),
				}
			},
			Packages: func(_ fyne.CanvasObject, _ Context) []string {
				return []string{"canvas", "image/color"}
			},
		},
		"*canvas.Polygon": {
			Name: "Polygon",
			Create: func(Context) fyne.CanvasObject {
				rect := canvas.NewPolygon(5, color.Black)
				rect.StrokeColor = color.Black
				return rect
			},
			Edit: func(obj fyne.CanvasObject, _ Context, _ func([]*widget.FormItem), onchanged func()) []*widget.FormItem {
				p := obj.(*canvas.Polygon)
				return []*widget.FormItem{
					widget.NewFormItem("Sides", newIntSliderButton(float64(p.Sides), 3, 18, func(f float64) {
						p.Sides = uint(f)
						p.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Fill", newColorButton(p.FillColor, func(c color.Color) {
						p.FillColor = c
						p.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Corner", newIntSliderButton(float64(p.CornerRadius), 0, 32, func(f float64) {
						p.CornerRadius = float32(f)
						p.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Angle", newIntSliderButton(float64(p.Angle), -360, 360, func(f float64) {
						p.Angle = float32(f)
						p.Refresh()
						onchanged()
					})),

					widget.NewFormItem("Stroke", newIntSliderButton(float64(p.StrokeWidth), 0, 32, func(f float64) {
						p.StrokeWidth = float32(f)
						p.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Color", newColorButton(p.StrokeColor, func(c color.Color) {
						p.StrokeColor = c
						p.Refresh()
						onchanged()
					})),
				}
			},
			Packages: func(_ fyne.CanvasObject, _ Context) []string {
				return []string{"canvas", "image/color"}
			},
		},
		"*canvas.RadialGradient": {
			Name: "RadialGradient",
			Create: func(Context) fyne.CanvasObject {
				return &canvas.RadialGradient{StartColor: color.White}
			},
			Edit: func(obj fyne.CanvasObject, _ Context, _ func([]*widget.FormItem), onchanged func()) []*widget.FormItem {
				r := obj.(*canvas.RadialGradient)
				return []*widget.FormItem{
					widget.NewFormItem("Start", newColorButton(r.StartColor, func(c color.Color) {
						r.StartColor = c
						r.Refresh()
						onchanged()
					})),
					widget.NewFormItem("End", newColorButton(r.EndColor, func(c color.Color) {
						r.EndColor = c
						r.Refresh()
						onchanged()
					})),
				}
			},
			Packages: func(_ fyne.CanvasObject, _ Context) []string {
				return []string{"canvas", "image/color"}
			},
		},
		"*canvas.Rectangle": {
			Name: "Rectangle",
			Create: func(Context) fyne.CanvasObject {
				rect := canvas.NewRectangle(color.Black)
				rect.StrokeColor = color.Black
				return rect
			},
			Edit: func(obj fyne.CanvasObject, _ Context, _ func([]*widget.FormItem), onchanged func()) []*widget.FormItem {
				r := obj.(*canvas.Rectangle)
				return []*widget.FormItem{
					widget.NewFormItem("Fill", newColorButton(r.FillColor, func(c color.Color) {
						r.FillColor = c
						r.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Corner", newIntSliderButton(float64(r.CornerRadius), 0, 32, func(f float64) {
						r.CornerRadius = float32(f)
						r.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Stroke", newIntSliderButton(float64(r.StrokeWidth), 0, 32, func(f float64) {
						r.StrokeWidth = float32(f)
						r.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Color", newColorButton(r.StrokeColor, func(c color.Color) {
						r.StrokeColor = c
						r.Refresh()
						onchanged()
					})),
				}
			},
			Packages: func(_ fyne.CanvasObject, _ Context) []string {
				return []string{"canvas", "image/color"}
			},
		},
		"*canvas.Text": {
			Name: "Text",
			Create: func(Context) fyne.CanvasObject {
				rect := canvas.NewText("Text", color.Black)
				return rect
			},
			Edit: func(obj fyne.CanvasObject, _ Context, _ func([]*widget.FormItem), onchanged func()) []*widget.FormItem {
				t := obj.(*canvas.Text)
				e := widget.NewEntry()
				e.SetText(t.Text)
				e.OnChanged = func(text string) {
					t.Text = text
					t.Refresh()
					onchanged()
				}

				bold := widget.NewCheck("", func(on bool) {
					t.TextStyle.Bold = on
					t.Refresh()
					onchanged()
				})
				bold.Checked = t.TextStyle.Bold
				italic := widget.NewCheck("", func(on bool) {
					t.TextStyle.Italic = on
					t.Refresh()
					onchanged()
				})
				italic.Checked = t.TextStyle.Italic
				mono := widget.NewCheck("", func(on bool) {
					t.TextStyle.Monospace = on
					t.Refresh()
					onchanged()
				})
				mono.Checked = t.TextStyle.Monospace

				return []*widget.FormItem{
					widget.NewFormItem("Text", e),
					widget.NewFormItem("Color", newColorButton(t.Color, func(c color.Color) {
						t.Color = c
						t.Refresh()
						onchanged()
					})),
					widget.NewFormItem("TextSize", newIntSliderButton(float64(t.TextSize), 4, 64, func(f float64) {
						t.TextSize = float32(f)
						t.Refresh()
						onchanged()
					})),
					widget.NewFormItem("Bold", bold),
					widget.NewFormItem("Italic", italic),
					widget.NewFormItem("Monospace", mono),
				}
			},
			Packages: func(_ fyne.CanvasObject, _ Context) []string {
				return []string{"canvas", "image/color"}
			},
		},
	}

	GraphicsNames = extractNames(Graphics)
}

func initImageGraphic() WidgetInfo {
	return WidgetInfo{
		Name: "Image",
		Create: func(Context) fyne.CanvasObject {
			return &canvas.Image{}
		},
		Edit: func(obj fyne.CanvasObject, c Context, _ func([]*widget.FormItem), onchanged func()) []*widget.FormItem {
			i := obj.(*canvas.Image)
			props := c.Metadata()[obj]

			minWidthInput := widget.NewEntry()
			minWidthInput.SetText(props["minWidth"])
			minHeightInput := widget.NewEntry()
			minHeightInput.SetText(props["minHeight"])

			minWidthInput.Validator = func(s string) error {
				if s == "" {
					return nil
				}

				f, err := strconv.ParseFloat(s, 32)
				if err != nil {
					return errors.New("invalid number format")
				}
				if f < 0 {
					return errors.New("negative minimum size")
				}
				return nil
			}
			minHeightInput.Validator = minWidthInput.Validator

			updateMin := func(_ string) {
				w, err := strconv.ParseFloat(minWidthInput.Text, 32)
				if err != nil {
					props["minWidth"] = ""
					return
				} else {
					props["minWidth"] = minWidthInput.Text
				}
				h, err := strconv.ParseFloat(minHeightInput.Text, 32)
				if err != nil {
					props["minHeight"] = ""
					return
				} else {
					props["minHeight"] = minHeightInput.Text
				}

				s := fyne.Size{}
				if w > 0 || h > 0 {
					s = fyne.NewSize(float32(w), float32(h))
				}
				i.SetMinSize(s)
				onchanged()
			}
			minWidthInput.OnChanged = updateMin
			minHeightInput.OnChanged = updateMin

			fill := widget.NewSelect([]string{"Stretch", "Contain", "Cover"}, func(s string) {
				mode := canvas.ImageFillStretch
				switch s {
				case "Contain":
					mode = canvas.ImageFillContain
				case "Cover":
					mode = canvas.ImageFillCover
				}
				i.FillMode = mode
				i.Refresh()
				onchanged()
			})
			fill.SetSelectedIndex(int(i.FillMode))

			cornerRadius := newIntSliderButton(float64(i.CornerRadius), 0, 32, func(f float64) {
				i.CornerRadius = float32(f)
				i.Refresh()
				onchanged()
			})

			// TODO move to go:embed for files!
			pathSelect := widget.NewButton("(No file)", nil)
			if i.File != "" {
				pathSelect.SetText(filepath.Base(i.File))
			}
			pathSelect.OnTapped = func() {
				dialog.ShowFileOpen(func(r fyne.URIReadCloser, err error) {
					path := ""
					if r != nil && err == nil {
						_ = r.Close()
						path = r.URI().Path()
					}

					i.File = path
					if path == "" {
						i.Image = nil
						pathSelect.SetText("(No file)")
					} else {
						pathSelect.SetText(r.URI().Name())
					}
					i.Refresh()
					onchanged()
				}, fyne.CurrentApp().Driver().AllWindows()[0])
			}

			resSelect := newIconSelectorButton(i.Resource, func(res fyne.Resource) {
				i.Resource = res
				i.Refresh()
				onchanged()
			}, true)
			resSelect.SetIcon(i.Resource)

			return []*widget.FormItem{
				widget.NewFormItem("Min Width", minWidthInput),
				widget.NewFormItem("Min Height", minHeightInput),
				widget.NewFormItem("Fill mode", fill),
				widget.NewFormItem("Corner", cornerRadius),
				widget.NewFormItem("Path", pathSelect),
				widget.NewFormItem("Resource", resSelect),
			}
		},
		Packages: func(obj fyne.CanvasObject, _ Context) []string {
			i := obj.(*canvas.Image)
			if i.Resource != nil {
				return []string{"canvas", "theme"}
			}
			return []string{"canvas"}
		},
		Gostring: func(obj fyne.CanvasObject, c Context, defs map[string]string) string {
			i := obj.(*canvas.Image)
			props := c.Metadata()[obj]
			minWidth, _ := props["minWidth"]
			minHeight, _ := props["minHeight"]
			hasMin := (minWidth != "" && minWidth != "0") || (minHeight != "" && minHeight != "0")

			code := ""
			if i.Resource != nil {
				res := "theme." + IconName(i.Resource) + "()"

				if !hasMin && i.FillMode == canvas.ImageFillStretch && i.CornerRadius == 0 {
					code = fmt.Sprintf("canvas.NewImageFromResource(%s)", res)
				} else {
					code = fmt.Sprintf("&canvas.Image{Resource: %s, FillMode: %s, CornerRadius: %f}", res, fillName(i.FillMode), i.CornerRadius)

					if hasMin {
						code = fmt.Sprintf("func() *canvas.Image {"+
							"img := %s; img.SetMinSize(%#v); return img}()", code, i.MinSize())
					}
				}
			} else {
				if !hasMin && i.FillMode == canvas.ImageFillStretch && i.CornerRadius == 0 {
					code = fmt.Sprintf("canvas.NewImageFromFile(\"%s\")", i.File)
				} else {
					code = fmt.Sprintf("&canvas.Image{File: \"%s\", FillMode: %s, CornerRadius: %f}", i.File, fillName(i.FillMode), i.CornerRadius)

					if hasMin {
						code = fmt.Sprintf("func() *canvas.Image {"+
							"img := %s; img.SetMinSize(%#v); return img}()", code, i.MinSize())
					}
				}
			}

			return widgetRef(props, defs, code)
		},
	}
}

func newColorButton(c color.Color, fn func(color.Color)) fyne.CanvasObject {
	// TODO get the window passed in somehow
	w := fyne.CurrentApp().Driver().AllWindows()[0]

	input := widget.NewEntry()
	input.SetText(formatColor(c))
	preview := newColorTapper(c, func(col color.Color) {
		raw := formatColor(col)
		input.SetText(raw)
		fn(col)
	}, w)

	input.OnChanged = func(raw string) {
		c := parseColor(raw)
		preview.setColor(c)
		fn(c)
	}
	return container.NewBorder(nil, nil, preview, nil, input)
}

type colorTapper struct {
	widget.BaseWidget

	r   *canvas.Rectangle
	fn  func(color.Color)
	win fyne.Window
}

func newColorTapper(c color.Color, fn func(color.Color), win fyne.Window) *colorTapper {
	preview := canvas.NewRectangle(c)
	preview.SetMinSize(fyne.NewSquareSize(32))

	t := &colorTapper{r: preview, fn: fn, win: win}
	t.ExtendBaseWidget(t)
	return t
}

func (c *colorTapper) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.r)
}

func (c *colorTapper) Tapped(_ *fyne.PointEvent) {
	d := dialog.NewColorPicker("Choose Color", "Pick a color", c.fn, c.win)
	d.Advanced = true
	d.Show()
}

func (c *colorTapper) setColor(col color.Color) {
	c.r.FillColor = col
	c.r.Refresh()
}

func newIntSliderButton(f float64, start, end float64, fn func(float64)) fyne.CanvasObject {
	return newSliderButtonWithConversion(f, start, end, 1, "%0.0f", fn)
}

func newRatioSliderButton(f float64, start, end float64, fn func(float64)) fyne.CanvasObject {
	return newSliderButtonWithConversion(f, start, end, 0.01, "%0.2f", fn)
}

func newSliderButtonWithConversion(f float64, start, end, step float64, format string, fn func(float64)) fyne.CanvasObject {
	input := widget.NewEntry()
	input.SetText(strconv.Itoa(int(f)))
	slide := widget.NewSlider(start, end)
	slide.SetValue(f)
	slide.Step = step

	input.OnChanged = func(s string) {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return
		}

		slide.SetValue(f)
		fn(f)
	}
	slide.OnChanged = func(f float64) {
		input.SetText(fmt.Sprintf(format, f)) // format like an int
		fn(f)
	}
	return container.NewBorder(nil, nil, input, nil, slide)
}

func parseColor(s string) color.Color {
	if s == "" {
		return color.Black
	}

	var rgb int
	_, err := fmt.Sscanf(s, "#%x", &rgb)
	if err != nil {
		return color.Transparent
	}

	hasAlpha := len(s) > 7
	a := 0xff
	offset := 0
	if hasAlpha {
		a = rgb & 0xff
		offset = 8
	}

	b := rgb >> offset & 0xff
	gg := rgb >> (offset + 8) & 0xff
	r := rgb >> (offset + 16) & 0xff
	return color.NRGBA{R: uint8(r), G: uint8(gg), B: uint8(b), A: uint8(a)}
}

func formatColor(c color.Color) string {
	if c == nil {
		return "#000000"
	}
	ch := color.RGBAModel.Convert(c).(color.RGBA)
	if ch.A == 0xff {
		return fmt.Sprintf("#%.2x%.2x%.2x", ch.R, ch.G, ch.B)
	}

	return fmt.Sprintf("#%.2x%.2x%.2x%.2x", ch.R, ch.G, ch.B, ch.A)
}
