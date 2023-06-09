// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os/exec"
)

// Convert uses ImageMagick convert to convert an image from one format to another.
func Convert(input, output string) {
	cmd := exec.Command("rsvg-convert", input, "-o", output)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// ConvertToSize uses rsvg-convert to convert an svg to a png with a given size
func ConvertToSize(input, output string, width, height int) {
	w := fmt.Sprintf("%d", width)
	h := fmt.Sprintf("%d", height)
	cmd := exec.Command("rsvg-convert", "-w", w, "-h", h, input, "-o", output)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// LoadStyle loads an external stylesheet and returns the content.
func LoadStyle(path string) string {
	styleData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(styleData)
}

// Round3 rounds a number down to 3 decimals.
func Round3(val float64) float64 {
	return math.Round(val*1000) / 1000
}
