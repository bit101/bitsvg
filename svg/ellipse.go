// Package svg holds SVG object creation and manipulation methods.
package svg

import "encoding/xml"

// Ellipse creates a new Ellipse object.
type Ellipse struct {
	GraphicElement
	XMLName xml.Name `xml:"ellipse"`
	Cx      float64  `xml:"cx,attr"`
	Cy      float64  `xml:"cy,attr"`
	Rx      float64  `xml:"rx,attr"`
	Ry      float64  `xml:"ry,attr"`
}

// NewEllipse creates a new Ellipse object.
func NewEllipse(x, y, rx, ry float64) *Ellipse {
	return &Ellipse{
		GraphicElement: *NewGraphicElement(),
		Cx:             x,
		Cy:             y,
		Rx:             rx,
		Ry:             ry,
	}
}
