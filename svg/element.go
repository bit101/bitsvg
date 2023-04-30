// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"fmt"
	"math"
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
	ID            string  `xml:"id,attr,omitempty"`
	Stroke        string  `xml:"stroke,attr,omitempty"`
	Fill          string  `xml:"fill,attr,omitempty"`
	FillOpacity   float64 `xml:"fill-opacity,attr,omitempty"`
	StrokeWidth   float64 `xml:"stroke-width,attr,omitempty"`
	StrokeOpacity float64 `xml:"stroke-opacity,attr,omitempty"`
	StrokeDash    string  `xml:"stroke-dasharray,attr,omitempty"`
	Style         string  `xml:"style,attr,omitempty"`
	Class         string  `xml:"class,attr,omitempty"`
	LineCap       string  `xml:"stroke-linecap,attr,omitempty"`
	LineJoin      string  `xml:"stroke-linejoin,attr,omitempty"`
	MiterLimit    float64 `xml:"stroke-miterlimit,attr,omitempty"`
	Filters       string  `xml:"filter,attr,omitempty"`
	Transform     string  `xml:"transform,attr,omitempty"`
}

// NewGraphicElement returns a new GraphicElement
func NewGraphicElement() *GraphicElement {
	return &GraphicElement{}
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

// SetFillRandom sets the fill to a random color.
func (ge *GraphicElement) SetFillRandom() {
	ge.SetFillColor(color.RandomRGB())
}

// SetFillRGB sets the fill color of this element with rgb integer values and full opacity.
func (ge *GraphicElement) SetFillRGB(r, g, b int) {
	ge.Fill = fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

// SetFillOpacity sets the opacity of the stroke on this element.
func (ge *GraphicElement) SetFillOpacity(opacity float64) {
	ge.FillOpacity = Round3(opacity)
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
	ge.StrokeWidth = Round3(width)
}

// SetStrokeOpacity sets the width of the stroke on this element.
func (ge *GraphicElement) SetStrokeOpacity(width float64) {
	ge.StrokeOpacity = Round3(width)
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

// SetLineCap sets the line cap of this element.
func (ge *GraphicElement) SetLineCap(cap string) {
	ge.LineCap = cap
}

// SetLineJoin sets the line join of this polygon.
func (ge *GraphicElement) SetLineJoin(cap string) {
	ge.LineJoin = cap
}

// SetMiterLimit sets the miter limit of this polygon.
func (ge *GraphicElement) SetMiterLimit(limit float64) {
	ge.MiterLimit = Round3(limit)
}

// SetStyle sets an inline style.
func (ge *GraphicElement) SetStyle(style string) {
	ge.Style = style
}

// SetFilters sets a list of filters to use on this element.
func (ge *GraphicElement) SetFilters(filters ...string) {
	var sb strings.Builder
	for _, filter := range filters {
		fmt.Fprintf(&sb, "url(#%s) ", filter)
	}
	ge.Filters = sb.String()
}

// RotateDeg rotates the element the given number of degrees.
func (ge *GraphicElement) RotateDeg(deg float64) {
	ge.Transform += fmt.Sprintf("rotate(%g) ", Round3(deg))
}

// RotateRad rotates the element the given number of radians.
func (ge *GraphicElement) RotateRad(rad float64) {
	ge.Transform += fmt.Sprintf("rotate(%g) ", Round3(rad*180/math.Pi))
}

// Translate rotates the element the given number of degrees.
func (ge *GraphicElement) Translate(x, y float64) {
	ge.Transform += fmt.Sprintf("translate(%g, %g) ", Round3(x), Round3(y))
}

// Scale scales the element the given percent.
func (ge *GraphicElement) Scale(x, y float64) {
	ge.Transform += fmt.Sprintf("scale(%g, %g) ", Round3(x), Round3(y))
}

// Uniscale scales the element the given percent.
func (ge *GraphicElement) Uniscale(s float64) {
	ge.Transform += fmt.Sprintf("scale(%g) ", Round3(s))
}

// IsElement returns true for Elements.
func (ge *GraphicElement) IsElement() bool {
	return true
}
