package srv

import (
	"bufio"
	"os"
)

type FileWriter struct {
	FileName string
	File     *os.File
}

func NewFileWriter(fileName string) *FileWriter {

	return &FileWriter{
		FileName: fileName,
		File:     &os.File{},
	}

}

// write string to file
func (f *FileWriter) Write(s string) error {

	f.File, _ = os.Create(f.FileName)
	defer f.File.Close()
	writer := bufio.NewWriter(f.File)
	_, err := writer.WriteString(s)
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}
