// Package main is a development and demo package for SVG.
package main

import (
	"github.com/bit101/bitsvg/svg"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

func main() {
	s := svg.NewSVG("my svg", 800, 800)
	s.SetBackgroundRGB(255, 128, 0)
	s.AddStyleSheet("style.css")
	s.AddStyleSheet("style2.css")
	filter := svg.NewDropShadow("drop-shadow", 6, 0.7, 6, 6)
	// filter.SetBoundPercent(35)
	s.AddFilter(filter)

	for i := 0; i < 300; i++ {
		size := random.FloatRange(50, 200)
		star := s.AddStar(random.FloatRange(0, 800), random.FloatRange(0, 800), 5, size*0.4, size, 0)
		star.SetFillRGB(random.IntRange(128, 255), random.IntRange(128, 255), random.IntRange(128, 255))
		star.Filter = "url(#drop-shadow)"
	}

	s.WriteToFile("out.svg")
	svg.Convert("out.svg", "out.png")
	// svg.ConvertToSize("out.svg", "out.png", 10000, 10000)
	util.ViewImage("out.png")
}
