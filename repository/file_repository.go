package repository

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileRepository interface {
	Save(file multipart.File, fileName string) (string, error)
}

type fileRepository struct {
	path string
}

func (f *fileRepository) Save(file multipart.File, fileName string) (string, error) {
	// menggabungkan file path + file name -> uploads/gambar-bagus.jpg
	fileLocation := filepath.Join(f.path, fileName)
	out, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer out.Close()
	// copy dari hasil file (multipart) kedalam file yang udah dibukan yang disimpan dalam variabel out
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return fileLocation, nil
}

func NewFileRepository(path string) FileRepository {
	return &fileRepository{path: path}
}
