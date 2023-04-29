// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
)

// SVG is a struct representing an SVG object.
type SVG struct {
	XMLName        xml.Name `xml:"svg"`
	Width          float64  `xml:"width,attr"`
	Height         float64  `xml:"height,attr"`
	Namespace      string   `xml:"xmlns,attr"`
	Title          string   `xml:"title"`
	Description    string   `xml:"desc"`
	Style          string   `xml:"style"`
	BackgroundRect *Rect    `xml:"rect"`
	Elements       []Element
}

// NewSVG creates a new SVG instance.
func NewSVG(title string, width, height float64) *SVG {
	bgRect := NewRect(0, 0, width, height)
	bgRect.NoFill()
	return &SVG{
		Namespace:      "http://www.w3.org/2000/svg",
		Title:          title,
		Width:          width,
		Height:         height,
		Style:          "background-color: white",
		BackgroundRect: bgRect,
	}
}

// AddElement adds a new GraphicElement to the SVG.
func (s *SVG) AddElement(element Element) {
	s.Elements = append(s.Elements, element)
}

// AddCircle adds a new Circle object to the SVG.
func (s *SVG) AddCircle(x, y, r float64) *Circle {
	c := NewCircle(x, y, r)
	s.AddElement(c)
	return c
}

// AddEllipse adds a new Ellipse object to the SVG.
func (s *SVG) AddEllipse(x, y, rx, ry float64) *Ellipse {
	e := NewEllipse(x, y, rx, ry)
	s.AddElement(e)
	return e
}

// AddLine adds a new Line object to the SVG.
func (s *SVG) AddLine(x1, y1, x2, y2 float64) *Line {
	l := NewLine(x1, y1, x2, y2)
	s.AddElement(l)
	return l
}

// AddPolygon adds a new Polygon object to the SVG.
func (s *SVG) AddPolygon(points ...float64) *Polygon {
	p := NewPolygon(points...)
	s.AddElement(p)
	return p
}

// AddPolyline adds a new Polyline object to the SVG.
func (s *SVG) AddPolyline(points ...float64) *Polyline {
	p := NewPolyline(points...)
	s.AddElement(p)
	return p
}

// AddRect adds a new Rect object to the SVG.
func (s *SVG) AddRect(x, y, w, h float64) *Rect {
	r := NewRect(x, y, w, h)
	s.AddElement(r)
	return r
}

// AddRegularPolygon adds a new Rect object to the SVG.
func (s *SVG) AddRegularPolygon(x, y float64, points int, radius, rotation float64) *Polygon {
	r := NewRegularPolygon(x, y, points, radius, rotation)
	s.AddElement(r)
	return r
}

// SetBackgroundRGB creates a background rect of the given color.
func (s *SVG) SetBackgroundRGB(r, g, b int) {
	s.BackgroundRect.SetFillRGB(r, g, b)
}

// WriteToFile saves this SVG object as an SVG file.
func (s *SVG) WriteToFile(filename string) error {
	svg, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return errors.New("Unable to marshal svg: " + err.Error())
	}
	// add header
	svg = []byte(xml.Header + string(svg))

	err = ioutil.WriteFile(filename, svg, 0x777)
	if err != nil {
		return errors.New("Unable to save svg: " + err.Error())
	}
	return nil
}
