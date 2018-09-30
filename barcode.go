package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"

	"github.com/boombuler/barcode"
)

type Direction string

type barcodeRect struct {
	barcode.Barcode
	Rectangle image.Rectangle
}

type CombinedBarcode struct {
	direction Direction
	barcodes  []barcodeRect
	options   CombinedBarcodeOptions
}

type CombinedBarcodeOptions struct {
	Padding int
}

const (
	horizontal Direction = "horizontal"
	vertical   Direction = "vertical"
)

func InitCombinedBarcode(direction Direction, barcodes []barcode.Barcode, options *CombinedBarcodeOptions) (*CombinedBarcode, error) {
	if len(barcodes) == 0 {
		return nil, errors.New("No barcodes supplied")
	}

	// TODO(ajafri): support horizontal barcodes as well
	if direction != vertical {
		return nil, errors.New("Only the vertical direction supported currently")
	}
	var width, height int

	if direction == vertical {
		// init the width of the combined barcode to
		width = barcodes[0].Bounds().Max.X
	} else if direction == horizontal {
		// TODO(ajafri):
	}

	var ret CombinedBarcode
	ret.direction = direction
	if options != nil {
		ret.options = *options
	}

	for _, b := range barcodes {
		switch direction {
		case vertical:
			// Make sure that all the barcodes have the same width
			if width == b.Bounds().Max.X {
				width = b.Bounds().Max.X
			} else {
				return nil, errors.New("Barcodes are not the same width")
			}

			top := height
			// Add padding between barcodes
			if top != 0 {
				top += ret.options.Padding
			}

			height += b.Bounds().Max.Y
			ret.barcodes = append(ret.barcodes, barcodeRect{
				Barcode:   b,
				Rectangle: image.Rect(0, top, width, height),
			})
		case horizontal:
			// TODO(ajafri): finish this
		}
	}

	return &ret, nil
}

// ColorModel returns the Image's color model.
func (cb CombinedBarcode) ColorModel() color.Model {
	return color.GrayModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (cb CombinedBarcode) Bounds() image.Rectangle {
	var width, height int
	for _, b := range cb.barcodes {
		if cb.direction == vertical {
			width = b.Bounds().Max.X
			height += b.Bounds().Max.Y
		} else if cb.direction == horizontal {
			// TODO(ajafri):
		}
	}

	return image.Rect(0, 0, width, height)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (cb CombinedBarcode) At(x, y int) color.Color {
	for _, b := range cb.barcodes {
		p := image.Point{X: x, Y: y}
		if p.In(b.Rectangle) {
			if cb.direction == vertical {
				// For vertical barcodes we need to adjust the Y coordinate
				// to the local barcode when getting the color value at that coordinate
				localY := y - b.Rectangle.Min.Y
				return b.Barcode.At(x, localY)
			} else {
				panic(fmt.Errorf("Unsupported direction %q", cb.direction))
			}

		}
	}

	return color.White
}
