package main

import (
	"image"
	"image/color"

	"github.com/boombuler/barcode"
)

type RotatedBarcode struct {
	barcode.Barcode
	Rotation bool
}

// ColorModel returns the Image's color model.
func (rb RotatedBarcode) ColorModel() color.Model {
	return color.GrayModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (rb RotatedBarcode) Bounds() image.Rectangle {
	if !rb.Rotation {
		return rb.Barcode.Bounds()
	}

	bounds := rb.Barcode.Bounds()
	return image.Rect(0, 0, bounds.Size().Y, bounds.Size().X)
}

func (rb RotatedBarcode) At(x, y int) color.Color {
	if !rb.Rotation {
		return rb.Barcode.At(x, y)
	} else {
		return rb.Barcode.At(y, x)
	}
}
