// Package svg holds SVG object creation and manipulation methods.
package svg

import "encoding/xml"

// Rect creates a new Rect object.
type Rect struct {
	GraphicElement
	XMLName xml.Name `xml:"rect"`
	X       float64  `xml:"x,attr"`
	Y       float64  `xml:"y,attr"`
	Width   float64  `xml:"width,attr"`
	Height  float64  `xml:"height,attr"`
	Rx      float64  `xml:"rx,attr,omitempty"`
	Ry      float64  `xml:"ry,attr,omitempty"`
}

// NewRect creates a new Rect object.
func NewRect(x, y, w, h float64) *Rect {
	return &Rect{
		GraphicElement: *NewGraphicElement(),
		X:              x,
		Y:              y,
		Width:          w,
		Height:         h,
	}
}
