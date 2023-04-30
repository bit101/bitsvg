// Package svg holds SVG object creation and manipulation methods.
package svg

import "encoding/xml"

// FilterEffect defines a filter effect.
type FilterEffect struct {
	XMLName xml.Name `xml:"filter"`
}
