package files

import (
	"log"
	"os"
	"path/filepath"
)

type FileData struct {
	FileName	string
	FileData	*os.File
}

func LoadFile(path string) *FileData {
	file, err := os.Open(path)	
	if err != nil {
		log.Panic(err)
	}
	fd := new(FileData)
	fd.FileName = filepath.Base(path)
	fd.FileData = file
	return fd
}
