package main

import (
	"image"
	"image/png"
	"io"
)

type Identicon struct {
	buffer  image.Image
	message string
}

func NewIdenticon(width, height int, message string) *Identicon {
	r := image.Rect(0, 0, width, height)
	i := image.NewNRGBA(r)
	return &Identicon{
		buffer:  i,
		message: message,
	}
}

func (idcon *Identicon) AsPNG(w io.Writer) error {
	return png.Encode(w, idcon.buffer)
}
