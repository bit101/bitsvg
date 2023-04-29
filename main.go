// Package main is a development and demo package for SVG.
package main

import (
	"github.com/bit101/bitsvg/svg"
	"github.com/bit101/blgo/util"
)

func main() {
	s := svg.NewSVG("my svg", 800, 800)
	s.SetBackgroundRGB(255, 128, 0)
	s.SetStyleSheet("style.css")

	star := s.AddStar(100, 100, 8, 40, 100, 0)
	star.ID = "star"

	for i := 0.0; i < 200; i += 3 {
		s.AddUse("#star", i, i)
	}

	s.WriteToFile("out.svg")
	// svg.Convert("out.svg", "out.png")
	// svg.ConvertToSize("out.svg", "out.png", 10000, 10000)
	util.ViewImage("out.svg")
}
