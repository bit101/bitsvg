// Package svg holds SVG object creation and manipulation methods.
package svg

import "encoding/xml"

// Circle creates a new Circle object.
type Circle struct {
	GraphicElement
	XMLName xml.Name `xml:"circle"`
	Cx      float64  `xml:"cx,attr"`
	Cy      float64  `xml:"cy,attr"`
	Radius  float64  `xml:"r,attr"`
}

// NewCircle creates a new Circle object.
func NewCircle(x, y, r float64) *Circle {
	return &Circle{
		GraphicElement: *NewGraphicElement(),
		Cx:             x,
		Cy:             y,
		Radius:         r,
	}
}
