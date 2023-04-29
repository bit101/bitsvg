// Package svg holds SVG object creation and manipulation methods.
package svg

import "encoding/xml"

// Use creates a new Use object.
type Use struct {
	GraphicElement
	XMLName xml.Name `xml:"use"`
	Href    string   `xml:"href,attr"`
	X       float64  `xml:"x,attr"`
	Y       float64  `xml:"y,attr"`
}

// NewUse creates a new Use object.
func NewUse(href string, x, y float64) *Use {
	return &Use{
		GraphicElement: *NewGraphicElement(),
		Href:           href,
		X:              x,
		Y:              y,
	}
}
