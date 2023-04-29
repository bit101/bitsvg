// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// EvenOdd is the fill rule "evenodd"
const EvenOdd = "evenodd"

// NonZero is the fill rule "NonZero"
const NonZero = "nonzero"

// Polygon creates a new Polygon object.
type Polygon struct {
	GraphicElement
	XMLName  xml.Name `xml:"polygon"`
	Points   string   `xml:"points,attr"`
	FillRule string   `xml:"fill-rule,attr"`
}

// NewPolygon creates a new Polygon object.
func NewPolygon(points ...float64) *Polygon {
	pointStrings := []string{}
	for _, coord := range points {
		pointStrings = append(pointStrings, strconv.FormatFloat(coord, 'f', 3, 64))
	}
	ps := strings.Join(pointStrings, ",")
	return &Polygon{
		GraphicElement: *NewGraphicElement(),
		Points:         ps,
		FillRule:       EvenOdd,
	}
}
