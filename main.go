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

	rect := s.AddRect(200, 200, 400, 400)
	rect.Class = "rect"

	star := s.AddStar(400, 400, 8, 100, 400, 0)
	star.Class = "poly"

	s.WriteToFile("out.svg")
	svg.Convert("out.svg", "out.png")
	util.ViewImage("out.png")
}
