package file

import (
	"bytes"
	"strings"

	"github.com/go-pdf/fpdf"
)

func NewPDFGeneratorImpl() PDFGenerator {
	return &pdfGeneratorImpl{}
}

func (l *pdfGeneratorImpl) Generate(content string) ([]byte, error) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	cellWidth := 190.0

	contentLines := strings.Split(content, "|")
	for _, val := range contentLines {
		pdf.MultiCell(cellWidth, 10, val, "", "L", false)
	}

	var buf bytes.Buffer

	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
