package main

import (
	"fmt"
	"image"
)

func rgbaOfSize(x, y int) *image.RGBA {
	r := image.Rect(0, 0, x, y)
	i := image.NewRGBA(r)
	return i
}

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

func main() {
	fmt.Println("[[ Identicon ]]")
	i := rgbaOfSize(8, 8)
	traverse(i)
}
