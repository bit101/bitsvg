// Package main is a development and demo package for SVG.
package main

import (
	"github.com/bit101/bitsvg/svg"
)

func main() {
	s := svg.NewSVG("my svg", 800, 800)
	s.SetBackgroundRGB(255, 128, 0)
	s.SetStyleSheet("style.css")

	x := 100.0
	y := 200.0
	r := 0.0
	for i := 0; i < 300; i++ {
		q := s.AddStar(x, y, 5, 40, 100, r)
		q.Class = "poly"
		x += 2
		y += 1.1
		r += 0.01
	}

	s.WriteToFile("out.svg")
}
