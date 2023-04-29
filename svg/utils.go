// Package svg holds SVG object creation and manipulation methods.
package svg

import (
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
