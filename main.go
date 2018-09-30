package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
)

func main() {
	cb, err := InitCombinedBarcode(
		singleBarcode("112312"),
	)
	if err != nil {
		panic(err)
	}

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, cb)
}

func singleBarcode(data string) barcode.Barcode {
	// Create the barcode
	// qrCode, err := qr.Encode("12", qr.M, qr.Numeric)
	qrCode, err := datamatrix.Encode(data)
	if err != nil {
		panic(err)
	}

	// Scale the barcode to 200x200 pixels
	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		panic(err)
	}

	return qrCode
}
