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
	s.SetStyleSheet("style2.css")

	t := s.AddText(100, 100, "Hello world")
	t.ID = "text"
	t.FontSize = 18
	t.FontFamily = "sans"
	t.FontStyle = svg.Italic
	t.TextDecoration = svg.Underline

	star := s.AddStar(400, 400, 5, 40, 100, 0)
	star.ID = "star"

	s.WriteToFile("out.svg")
	svg.Convert("out.svg", "out.png")
	// svg.ConvertToSize("out.svg", "out.png", 10000, 10000)
	util.ViewImage("out.png")
}
