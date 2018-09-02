package main

import (
	"fmt"
	"image"
	"os"
)

func traverse(i *image.RGBA) {
	it := NewImagerator(i)
	lastY := 0
	for it.More() {
		x, y := it.Next()
		if y > lastY {
			lastY = y
			fmt.Println()
		}
		fmt.Printf("(%d,%d)", x, y)
	}
	fmt.Println()
}

func checkFileError(e error) {
	if e != nil {
		fmt.Printf("Could not save identicon: %s", e)
		os.Exit(2)
	}
}

func main() {
	fmt.Println("[[ Identicon ]]")
	message := "AWITTYUSERNAME"
	con := NewIdenticon(64, 64, message)
	f, ferr := os.Create("/tmp/output.PNG")
	defer f.Close()
	checkFileError(ferr)
	werr := con.AsPNG(f)
	checkFileError(werr)
}
