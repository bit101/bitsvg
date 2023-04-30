// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Polyline creates a new Polyline object.
type Polyline struct {
	GraphicElement
	XMLName xml.Name `xml:"polyline"`
	Points  string   `xml:"points,attr"`
}

// NewPolyline creates a new Polyline object.
func NewPolyline(points ...float64) *Polyline {
	pointStrings := []string{}
	for _, coord := range points {
		pointStrings = append(pointStrings, fmt.Sprintf("%g", Round3(coord)))
	}
	ps := strings.Join(pointStrings, ",")
	polyline := &Polyline{
		GraphicElement: *NewGraphicElement(),
		Points:         ps,
	}
	polyline.NoFill()
	polyline.SetStrokeWidth(1)
	return polyline
}
