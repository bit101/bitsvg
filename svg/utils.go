// Package svg holds SVG object creation and manipulation methods.
package svg

import (
	"fmt"
	"log"
	"os/exec"
)

// Convert uses ImageMagick convert to convert an image from one format to another.
func Convert(input, output string) {
	cmd := exec.Command("convert", input, output)
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
