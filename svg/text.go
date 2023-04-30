// Package svg holds SVG object creation and manipulation methods.
package svg

import "encoding/xml"

// Text creates a new Text object.
type Text struct {
	GraphicElement
	XMLName        xml.Name `xml:"text"`
	X              float64  `xml:"x,attr"`
	Y              float64  `xml:"y,attr"`
	Text           string   `xml:",innerxml"`
	FontSize       float64  `xml:"font-size,attr,omitempty"`
	FontFamily     string   `xml:"font-family,attr,omitempty"`
	FontWeight     string   `xml:"font-weight,attr,omitempty"`
	FontStyle      string   `xml:"font-style,attr,omitempty"`
	TextDecoration string   `xml:"text-decoration,attr,omitempty"`
}

// NewText creates a new Text object.
func NewText(x, y float64, text string) *Text {
	return &Text{
		GraphicElement: *NewGraphicElement(),
		X:              x,
		Y:              y,
		Text:           text,
	}
}

// SetFontSize sets the size of the font used in this text.
func (t *Text) SetFontSize(size float64) {
	t.FontSize = size
}

// SetFontFamily sets the font family in this text.
func (t *Text) SetFontFamily(fontFamily string) {
	t.FontFamily = fontFamily
}

// SetFontStyle sets the font style in this text.
func (t *Text) SetFontStyle(fontStyle string) {
	t.FontStyle = fontStyle
}

// SetTextDecoration sets the text decoration type of this text.
func (t *Text) SetTextDecoration(textDecoration string) {
	t.TextDecoration = textDecoration
}
