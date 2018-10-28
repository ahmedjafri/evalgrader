package main

import (
	"encoding/json"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/datamatrix"
	"github.com/boombuler/barcode/pdf417"
	"github.com/jung-kurt/gofpdf"
)

const questionsFile = "questions.json"

func main() {
	b, err := ioutil.ReadFile(questionsFile)
	if err != nil {
		panic(err)
	}

	var questions []Question
	err = json.Unmarshal(b, &questions)
	if err != nil {
		panic(err)
	}

	//createCombinedBarcode(qrCodeFileName, []string{"11231234234412", "1123123"})
	createPDF417Barcode(qrCodeFileName, "testasfasdfasdasfasfsfsdf1")

	createPDFWithBarcode(questions, qrCodeFileName)
}

func createPDF417Barcode(filename string, data string) {
	// Create the barcode
	// qrCode, err := qr.Encode("12", qr.M, qr.Numeric)
	qrCode, err := pdf417.Encode(data, 2)
	if err != nil {
		panic(err)
	}

	rb := RotatedBarcode{Rotation: true, Barcode: qrCode}

	// create the output file
	file, _ := os.Create(filename)
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, rb)
}

func createPDFWithBarcode(questions []Question, barcodeFileName string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)

	offset := float64(10)

	for _, question := range questions {
		pdf.ImageOptions(barcodeFileName, 15, offset+7, 10, 25, false, gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: "",
		}, 0, "")

		pdf.Text(30, offset+5, question.Question)

		answersOffset := float64(offset + 10)
		answerOffset := 6
		for i, answer := range question.Answers {
			aOffset := answersOffset + float64(answerOffset*i)
			pdf.Circle(32, aOffset, 2, "D")
			pdf.Text(36, aOffset+1, answer.Answer)
		}

		offset += 40
	}

	fileStr := "qrcode.pdf"
	err := pdf.OutputFileAndClose(fileStr)
	if err != nil {
		panic(err)
	}
}

func createCombinedBarcode(filename string, data []string) {
	var b []barcode.Barcode
	for _, datum := range data {
		b = append(b, singleBarcode(datum))
	}

	cb, err := InitCombinedBarcode(vertical, b,
		&CombinedBarcodeOptions{
			Padding: 10,
		},
	)

	if err != nil {
		panic(err)
	}

	// create the output file
	file, _ := os.Create(filename)
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

const qrCodeFileName = "qrcode.png"
