// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// Polyline creates a new Polyline object.
type Polyline struct {
	GraphicElement
	XMLName    xml.Name `xml:"polyline"`
	Points     string   `xml:"points,attr"`
	LineJoin   string   `xml:"stroke-linejoin,attr"`
	LineCap    string   `xml:"stroke-linecap,attr"`
	MiterLimit float64  `xml:"stroke-miterlimit,attr"`
}

// NewPolyline creates a new Polyline object.
func NewPolyline(points ...float64) *Polyline {
	pointStrings := []string{}
	for _, coord := range points {
		pointStrings = append(pointStrings, strconv.FormatFloat(coord, 'f', 3, 64))
	}
	ps := strings.Join(pointStrings, ",")
	polyline := &Polyline{
		GraphicElement: *NewGraphicElement(),
		Points:         ps,
		LineJoin:       Miter,
		LineCap:        Butt,
		MiterLimit:     4,
	}
	polyline.NoFill()
	polyline.SetStrokeWidth(1)
	return polyline
}

// SetLineCap sets the line cap of this polyline.
func (p *Polyline) SetLineCap(cap string) {
	p.LineCap = cap
}

// SetLineJoin sets the line join of this polyline.
func (p *Polyline) SetLineJoin(cap string) {
	p.LineJoin = cap
}

// SetMiterLimit sets the miter limit of this polyline.
func (p *Polyline) SetMiterLimit(limit float64) {
	p.MiterLimit = limit
}
