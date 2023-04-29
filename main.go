// Package main is a development and demo package for SVG.
package main

import (
	"github.com/bit101/bitsvg/svg"
)

func main() {
	s := svg.NewSVG("my svg", 800, 800)
	s.Description = "Just a thing I made"
	s.SetBackgroundRGB(255, 128, 0)

	x := 200.0
	y := 200.0
	r := 0.0
	for i := 0; i < 150; i++ {
		q := s.AddRegularPolygon(x, y, 5, 100, r)
		q.NoFill()
		q.StrokeWidth = 0.5
		x += 3
		y += 3
		r += 0.01
	}

	s.WriteToFile("out.svg")
}
