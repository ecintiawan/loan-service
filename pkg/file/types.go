package file

type (
	File interface {
		Write(content []byte, filePath, fileName string) error
	}

	PDFGenerator interface {
		Generate(content string) ([]byte, error)
	}

	fileImpl         struct{}
	pdfGeneratorImpl struct{}
)
