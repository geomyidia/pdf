package httpimg_test

import (
	fpdf "github.com/geomyidia/pdf"
	"github.com/geomyidia/pdf/contrib/httpimg"
	"github.com/geomyidia/pdf/internal/example"
)

func ExampleRegister() {
	pdf := fpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(fpdf.RGB{R: 200, G: 200, B: 220})
	pdf.AddPage()

	url := "https://github.com/geomyidia/pdf/raw/main/image/logo_gofpdf.jpg"
	httpimg.Register(pdf, url, "")
	pdf.Image(url, 15, 15, 267, 0, false, "", 0, "")
	fileStr := example.Filename("contrib_httpimg_Register")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/contrib_httpimg_Register.pdf
}
