package tiff_test

import (
	fpdf "github.com/geomyidia/pdf"
	"github.com/geomyidia/pdf/contrib/tiff"
	"github.com/geomyidia/pdf/internal/example"
)

// ExampleRegisterFile demonstrates the loading and display of a TIFF image.
func ExampleRegisterFile() {
	pdf := fpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(fpdf.RGB{R: 200, G: 200, B: 220})
	pdf.AddPageFormat("L", fpdf.SizeType{Wd: 200, Ht: 200})
	opt := fpdf.ImageOptions{ImageType: "tiff", ReadDpi: false}
	_ = tiff.RegisterFile(pdf, "sample", opt, "../../image/golang-gopher.tiff")
	pdf.Image("sample", 0, 0, 200, 200, false, "", 0, "")
	fileStr := example.Filename("Fpdf_Contrib_Tiff")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
	// Output:
	// Successfully generated ../../pdf/Fpdf_Contrib_Tiff.pdf
}
