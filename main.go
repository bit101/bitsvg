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

	path := svg.NewPath()
	path.MoveTo(100, 100)
	path.CubicCurveTo(700, 100, 100, 300, 700, 700)
	path.NoFill()
	path.Stroke = "black"

	g := s.AddGroupDef()
	g.ID = "bez"
	g.AddElement(path)

	s.AddUse("bez", -100, -100)

	s.WriteToFile("out.svg")
	svg.Convert("out.svg", "out.png")
	// svg.ConvertToSize("out.svg", "out.png", 10000, 10000)
	util.ViewImage("out.png")
}
