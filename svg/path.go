// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"encoding/xml"
	"fmt"
)

// Path creates a new Path object.
type Path struct {
	GraphicElement
	XMLName xml.Name `xml:"path"`
	D       string   `xml:"d,attr"`
}

// NewPath creates a new Path object.
func NewPath() *Path {
	return &Path{
		GraphicElement: *NewGraphicElement(),
	}
}

// MoveTo moves the drawing cursor to the x y point.
func (p *Path) MoveTo(x, y float64) {
	p.D += fmt.Sprintf("M %g %g ", Round3(x), Round3(y))
	fmt.Println(p.D)
}
