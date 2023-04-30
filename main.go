// Package main is a development and demo package for SVG.
package main

import (
	"github.com/bit101/bitsvg/svg"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

func main() {
	random.Seed(1)
	s := svg.NewSVG("my svg", 800, 800)
	s.SetBackgroundRGB(255, 255, 255)
	s.AddStyleSheet("style.css")
	drop := svg.NewDropShadow("drop", 4, 0.7, 4, 4)
	drop.SetBoundPercent(40)
	s.AddFilter(drop)

	for i := 0; i < 1000; i++ {
		c := s.AddCircle(random.FloatRange(0, 800), random.FloatRange(0, 800), random.FloatRange(10, 100))
		c.SetFillRandom()
		c.SetFilters("drop")

	}

	path := s.AddPath()
	path.MoveTo(100, 100)

	s.WriteToFile("out.svg")
	svg.Convert("out.svg", "out.png")
	// svg.ConvertToSize("out.svg", "out.png", 10000, 10000)
	util.ViewImage("out.png")
}
