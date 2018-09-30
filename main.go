package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)

	cb, err := InitCombinedBarcode(vertical,
		[]barcode.Barcode{
			singleBarcode("11231234234412"),
			singleBarcode("1123123"),
		},
		&CombinedBarcodeOptions{
			Padding: 10,
		},
	)

	if err != nil {
		panic(err)
	}

	qrCodeFileName := "qrcode.png"
	// create the output file
	file, _ := os.Create(qrCodeFileName)
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, cb)

	pdf.Image(qrCodeFileName, 10, 10, 30, 0, false, "", 0, "")
	pdf.Text(50, 20, qrCodeFileName)
	fileStr := "qrcode.pdf"
	err = pdf.OutputFileAndClose(fileStr)
	if err != nil {
		panic(err)
	}
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
