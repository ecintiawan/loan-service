package file

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
)

func NewFileImpl() File {
	return &fileImpl{}
}

func (l *fileImpl) Write(content []byte, filePath, fileName string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, 0755)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s", filePath, fileName))
	if err != nil {
		return errorwrapper.E(err, errorwrapper.CodeInternal)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	_, err = w.Write(content)
	if err != nil {
		return errorwrapper.E(err, errorwrapper.CodeInternal)
	}
	w.Flush()

	return nil
}
