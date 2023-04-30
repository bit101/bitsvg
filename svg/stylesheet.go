// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"encoding/xml"
	"fmt"
)

// StyleSheet respresents an internal stylesheet.
type StyleSheet struct {
	XMLName   xml.Name `xml:"style"`
	Type      string   `xml:"type,attr"`
	StyleData string   `xml:",cdata"`
}

// NewStyleSheet creates a new StyleSheet with the given style data.
func NewStyleSheet(path string) *StyleSheet {
	styleData := LoadStyle(path)
	return &StyleSheet{
		Type:      "text/css",
		StyleData: fmt.Sprintf("/* %s */\n%s", path, styleData),
	}
}
