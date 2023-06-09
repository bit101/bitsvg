// Package svg holds SVG object creation and manipulation methods.
package svg

import "encoding/xml"

// Line creates a new Line object.
type Line struct {
	GraphicElement
	XMLName xml.Name `xml:"line"`
	X1      float64  `xml:"x1,attr"`
	Y1      float64  `xml:"y1,attr"`
	X2      float64  `xml:"x2,attr"`
	Y2      float64  `xml:"y2,attr"`
}

// NewLine creates a new Line object.
func NewLine(x1, y1, x2, y2 float64) *Line {
	line := &Line{
		GraphicElement: *NewGraphicElement(),
		X1:             Round3(x1),
		Y1:             Round3(y1),
		X2:             Round3(x2),
		Y2:             Round3(y2),
	}
	line.Stroke = "black"
	return line
}
