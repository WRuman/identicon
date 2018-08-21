package main

import (
	"image"
)

type Imagerator struct {
	rowsize int
	idx     int
	max     int
}

func NewImagerator(i image.Image) *Imagerator {
	rowsize := i.Bounds().Dx()
	return &Imagerator{
		rowsize,
		0,
		rowsize * i.Bounds().Dy(),
	}
}

func (iit *Imagerator) More() bool {
	return iit.idx < iit.max
}

func (iit *Imagerator) Next() (int, int) {
	x := iit.idx % iit.rowsize
	y := (iit.idx - x) / iit.rowsize
	iit.idx++
	return x, y
}
