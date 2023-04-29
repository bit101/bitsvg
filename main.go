// Package main is a development and demo package for SVG.
package main

import (
	"github.com/bit101/bitsvg/svg"
)

func main() {
	s := svg.NewSVG("my svg", 800, 800)
	s.Description = "Just a thing I made"
	s.SetBackgroundRGB(255, 128, 0)

	points := []float64{
		200, 100,
		300, 300,
		50, 150,
		350, 250,
		100, 300,
	}
	p := s.AddPolyline(points...)
	p.SetStrokeWidth(10)
	// p.SetFillRGB(255, 128, 0)

	s.WriteToFile("out.svg")
}
