// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bit101/blgo/color"
)

// TODO:
// Make HSL, HSLA, HSV, HSVA methods using blgo.color

// Element is an interface defining an SVG Element.
type Element interface {
	IsElement() bool

	SetFillColor(color.Color)
	SetFillRGB(int, int, int)
	SetFillOpacity(float64)
	NoFill()

	SetStrokeColor(color.Color)
	SetStrokeRGB(int, int, int)
	SetStrokeWidth(float64)
	SetStrokeOpacity(float64)
	SetStrokeDash(...int)
	NoStroke()

	SetStyle(string)
}

// GraphicElement represents an SVG Node.
type GraphicElement struct {
	ID            string  `xml:"id,attr"`
	Stroke        string  `xml:"stroke,attr"`
	Fill          string  `xml:"fill,attr"`
	FillOpacity   float64 `xml:"fill-opacity,attr"`
	StrokeWidth   float64 `xml:"stroke-width,attr"`
	StrokeOpacity float64 `xml:"stroke-opacity,attr"`
	StrokeDash    string  `xml:"stroke-dasharray,attr"`
	Style         string  `xml:"style,attr"`
	Class         string  `xml:"class,attr"`
}

// NewGraphicElement returns a new GraphicElement
func NewGraphicElement() *GraphicElement {
	return &GraphicElement{
		FillOpacity:   1.0,
		StrokeOpacity: 1.0,
		Stroke:        "black",
	}
}

// SetID sets the id of the element.
func (ge *GraphicElement) SetID(id string) {
	ge.ID = id
}

// SetFillColor sets the fill attribute using a blgo Color
func (ge *GraphicElement) SetFillColor(c color.Color) {
	r := int(c.R * 255)
	g := int(c.G * 255)
	b := int(c.B * 255)
	ge.SetFillRGB(r, g, b)
}

// SetFillRGB sets the fill color of this element with rgb integer values and full opacity.
func (ge *GraphicElement) SetFillRGB(r, g, b int) {
	ge.Fill = fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

// SetFillOpacity sets the opacity of the stroke on this element.
func (ge *GraphicElement) SetFillOpacity(opacity float64) {
	ge.FillOpacity = opacity
}

// NoFill sets the fill of the element to "none".
func (ge *GraphicElement) NoFill() {
	ge.Fill = "none"
}

// SetStrokeColor sets the stroke attribute using a blgo Color
func (ge *GraphicElement) SetStrokeColor(c color.Color) {
	r := int(c.R * 255)
	g := int(c.G * 255)
	b := int(c.B * 255)
	ge.SetStrokeRGB(r, g, b)
}

// SetStrokeRGB sets the stroke color of this element with rgb integer values and full opacity.
func (ge *GraphicElement) SetStrokeRGB(r, g, b int) {
	ge.Stroke = fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

// SetStrokeWidth sets the width of the stroke on this element.
func (ge *GraphicElement) SetStrokeWidth(width float64) {
	ge.StrokeWidth = width
}

// SetStrokeOpacity sets the width of the stroke on this element.
func (ge *GraphicElement) SetStrokeOpacity(width float64) {
	ge.StrokeOpacity = width
}

// SetStrokeDash sets the dash values for a stroke.
func (ge *GraphicElement) SetStrokeDash(dash ...int) {
	strs := []string{}
	for _, d := range dash {
		strs = append(strs, strconv.Itoa(d))
	}
	ge.StrokeDash = strings.Join(strs, ",")
}

// NoStroke sets the element to render without a stroke.
func (ge *GraphicElement) NoStroke() {
	ge.StrokeWidth = 0
}

// SetStyle sets the style attribute.
func (ge *GraphicElement) SetStyle(style string) {
	ge.Style = style
}

// IsElement returns true for Elements.
func (ge *GraphicElement) IsElement() bool {
	return true
}
