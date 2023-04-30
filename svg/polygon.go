// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"encoding/xml"
	"fmt"
	"math"
	"strings"
)

// Polygon creates a new Polygon object.
type Polygon struct {
	GraphicElement
	XMLName  xml.Name `xml:"polygon"`
	Points   string   `xml:"points,attr"`
	FillRule string   `xml:"fill-rule,attr,omitempty"`
}

// NewPolygon creates a new Polygon object.
func NewPolygon(points ...float64) *Polygon {
	pointStrings := []string{}
	for _, coord := range points {
		pointStrings = append(pointStrings, fmt.Sprintf("%g", Round3(coord)))
	}
	ps := strings.Join(pointStrings, ",")
	return &Polygon{
		GraphicElement: *NewGraphicElement(),
		Points:         ps,
	}
}

// NewRegularPolygon creates a regular polygon with the given number of points, radius and rotation.
func NewRegularPolygon(x, y float64, numPoints int, radius, rotation float64) *Polygon {
	coords := []float64{}
	for i := 0; i < numPoints; i++ {
		t := float64(i) / float64(numPoints) * math.Pi * 2
		coords = append(coords, x+math.Cos(t+rotation)*radius, y+math.Sin(t+rotation)*radius)
	}
	return NewPolygon(coords...)
}

// NewStar creates a star shape with the given number of points, inner and outer radii and rotation.
func NewStar(x, y float64, numPoints int, innerRadius, outerRadius, rotation float64) *Polygon {
	coords := []float64{}
	for i := 0; i < numPoints*2; i++ {
		r := innerRadius
		if i%2 == 0 {
			r = outerRadius
		}
		t := float64(i) / float64(numPoints) * math.Pi
		coords = append(coords, x+math.Cos(t+rotation)*r, y+math.Sin(t+rotation)*r)
	}
	return NewPolygon(coords...)
}
