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
	s.SetStyleSheet("style.css")

	g := s.AddGroup()
	c := svg.NewCircle(400, 400, 50)
	c.NoFill()
	c.Stroke = "black"
	l1 := svg.NewLine(350, 400, 450, 400)
	l2 := svg.NewLine(400, 350, 400, 450)
	g.AddElement(c)
	g.AddElement(l1)
	g.AddElement(l2)
	g.ID = "group"

	for i := 0; i < 100; i++ {
		s.AddUse("#group", random.FloatRange(-400, 400), random.FloatRange(-400, 400))
	}

	s.WriteToFile("out.svg")
	// svg.Convert("out.svg", "out.png")
	// svg.ConvertToSize("out.svg", "out.png", 10000, 10000)
	util.ViewImage("out.svg")
}
