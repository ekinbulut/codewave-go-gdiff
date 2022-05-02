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

var tmp string = "temp.html"

// write string to file
func (f *FileWriter) Write(s string) error {

	f.File, _ = os.Create(tmp)
	defer f.File.Close()
	writer := bufio.NewWriter(f.File)
	_, err := writer.WriteString(s)
	if err != nil {
		return err
	}
	writer.Flush()

	// compare diff between tmp and outputFile
	// if diff, then copy tmp to outputFile
	// if not, then delete tmp

	return nil
}

// rename file
func (f *FileWriter) Rename(newName string) error {

	err := os.Rename(tmp, newName)
	if err != nil {
		return err
	}
	return nil
}

// read file
func (f *FileWriter) Read() (string, error) {

	file, err := os.Open(f.FileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	content, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return content, nil
}

// chech if file exists
func (f *FileWriter) Exists() bool {
	if _, err := os.Stat(f.FileName); os.IsNotExist(err) {
		return false
	}
	return true
}
