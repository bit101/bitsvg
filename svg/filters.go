// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"encoding/xml"
	"fmt"
)

// FilterEffect defines a filter effect.
type FilterEffect struct {
	XMLName xml.Name `xml:"filter"`
	ID      string   `xml:"id,attr"`
	Filter  string   `xml:",innerxml"`
	X       string   `xml:"x,attr,omitempty"`
	Y       string   `xml:"y,attr,omitempty"`
	Width   string   `xml:"width,attr,omitempty"`
	Height  string   `xml:"height,attr,omitempty"`
}

// NewDropShadow creates a new drop shadow filter effect
func NewDropShadow(id string, blur, strength, dx, dy float64) *FilterEffect {
	fe := &FilterEffect{
		ID: id,
		Filter: fmt.Sprintf(`
    <feGaussianBlur in="SourceAlpha" stdDeviation="%f"/> 
    <feOffset dx="%f" dy="%f"/>
    <feComponentTransfer>
      <feFuncA type="linear" slope="%f"/>
    </feComponentTransfer>
    <feMerge> 
      <feMergeNode/>
      <feMergeNode in="SourceGraphic"/> 
    </feMerge>
		`, blur, dx, dy, strength),
	}
	return fe
}

// SetBoundPercent sets the bounds of the filter in terms of percent of the object's bounding box.
// Sets all bounds with a single value.
// The default values would be equivalent to passing 10 here, allowing for 10% more on each side.
func (fe *FilterEffect) SetBoundPercent(percent int) {
	fe.X = fmt.Sprintf("%d%%", -percent)
	fe.Y = fmt.Sprintf("%d%%", -percent)
	fe.Width = fmt.Sprintf("%d%%", 100+percent*2)
	fe.Height = fmt.Sprintf("%d%%", 100+percent*2)
}

// SetBoundPercents sets the bounds of the filter in terms of percent of the object's bounding box.
// Allows for setting the bounds on each edge separately.
// The default values would be equivalent to passing 10, 10, 120, 120 here, allowing for 10% more on each side.
func (fe *FilterEffect) SetBoundPercents(x, y, w, h int) {
	fe.X = fmt.Sprintf("%d%%", x)
	fe.Y = fmt.Sprintf("%d%%", y)
	fe.Width = fmt.Sprintf("%d%%", w)
	fe.Height = fmt.Sprintf("%d%%", h)
}
