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
	p.D += fmt.Sprintf("M%g %g ", Round3(x), Round3(y))
}

// MoveToRel moves the drawing cursor relative to its current location.
func (p *Path) MoveToRel(dx, dy float64) {
	p.D += fmt.Sprintf("m%g %g ", Round3(dx), Round3(dy))
}

// LineTo moves the drawing cursor to the x y point.
func (p *Path) LineTo(x, y float64) {
	p.D += fmt.Sprintf("L%g %g ", Round3(x), Round3(y))
}

// LineToRel moves the drawing cursor to the x y point.
func (p *Path) LineToRel(dx, dy float64) {
	p.D += fmt.Sprintf("l%g %g ", Round3(dx), Round3(dy))
}

// QuadraticCurveTo draws a quadratic bezier curve.
func (p *Path) QuadraticCurveTo(x1, y1, x2, y2 float64) {
	p.D += fmt.Sprintf("Q%g %g %g %g", Round3(x1), Round3(y1), Round3(x2), Round3(y2))
}

// CubicCurveTo draws a quadratic bezier curve.
func (p *Path) CubicCurveTo(x1, y1, x2, y2, x3, y3 float64) {
	p.D += fmt.Sprintf("C%g %g %g %g %g %g", Round3(x1), Round3(y1), Round3(x2), Round3(y2), Round3(x3), Round3(y3))
}

// Close draws a line back to the starting point.
func (p *Path) Close() {
	p.D += fmt.Sprintf("Z ")
}
