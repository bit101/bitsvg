// Package main is a development and demo package for SVG.
package main

import (
	"github.com/bit101/bitsvg/svg"
	"github.com/bit101/blgo/util"
)

func main() {
	s := svg.NewSVG("my svg", 800, 800)
	s.SetBackgroundRGB(255, 128, 0)

	x := 100.0
	y := 200.0
	r := 0.0
	for i := 0; i < 300; i++ {
		q := s.AddStar(x, y, 5, 40, 100, r)
		q.NoFill()
		q.SetStrokeWidth(0.25)

		x += 2
		y += 1.1
		r += 0.01
	}

	s.WriteToFile("out.svg")
	svg.Convert("out.svg", "out.png")
	util.ViewImage("out.png")
}
