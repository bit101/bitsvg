// Package main is a development and demo package for SVG.
package main

import (
	"github.com/bit101/bitsvg/svg"
	"github.com/bit101/blgo/color"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

func main() {
	random.Seed(1)
	s := svg.NewSVG("my svg", 800, 800)
	s.SetBackgroundRGB(255, 255, 255)
	s.AddStyleSheet("style.css")
	s.AddStyleSheet("style2.css")
	filter := svg.NewDropShadow("drop-shadow", 2, 0.5, 8, 8)
	filter.SetBoundPercent(40)
	s.AddFilter(filter)
	filter = svg.NewDropShadow("drop-shadow2", 2, 0.5, -8, -8)
	s.AddFilter(filter)
	filter = svg.NewGlowFilter("glow", 4, 1, 0, 0, 0.5)
	s.AddFilter(filter)
	filter = svg.NewBlurFilter("blur", 4)
	s.AddFilter(filter)

	for i := 0; i < 30; i++ {
		size := random.FloatRange(50, 200)
		star := s.AddStar(random.FloatRange(0, 800), random.FloatRange(0, 800), 5, size*0.4, size, 0)
		star.SetFillColor(color.HSV(random.FloatRange(190, 260), 1, 1))
		// star.Fill = "re"
		star.SetFilters("drop-shadow", "blur")
	}

	s.WriteToFile("out.svg")
	svg.Convert("out.svg", "out.png")
	// svg.ConvertToSize("out.svg", "out.png", 10000, 10000)
	util.ViewImage("out.png")
}
