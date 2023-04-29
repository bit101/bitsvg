// Package svg holds SVG object creation and manipulation methods.
package svg

import "encoding/xml"

// Group creates a new Circle object.
type Group struct {
	GraphicElement
	XMLName  xml.Name `xml:"g"`
	Elements []Element
}

// NewGroup creates a new Group object.
func NewGroup() *Group {
	return &Group{
		GraphicElement: *NewGraphicElement(),
	}
}

// AddElement adds and element to the group.
func (g *Group) AddElement(e Element) {
	g.Elements = append(g.Elements, e)
}
