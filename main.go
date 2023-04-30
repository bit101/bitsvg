// Package main is a development and demo package for SVG.
package main

import (
	"math"

	"github.com/bit101/bitsvg/svg"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

func main() {
	random.Seed(1)
	s := svg.NewSVG("my svg", 800, 800)
	s.SetBackgroundRGB(255, 255, 255)
	s.AddStyleSheet("style.css")
	star := s.AddStar(0, 0, 5, 40, 100, 0)
	star.ID = "star"

	for i := 0.0; i < 360; i++ {
		u := s.AddUse("star", 0, 0)
		u.Translate(0+i*1.6, 400)
		u.RotateRad(i * math.Pi / 180)
		u.Uniscale(i / 180)
	}

	s.WriteToFile("out.svg")
	svg.Convert("out.svg", "out.png")
	// svg.ConvertToSize("out.svg", "out.png", 10000, 10000)
	util.ViewImage("out.png")
}
