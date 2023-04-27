package usecase

import (
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
	"mime/multipart"
)

type FileUseCase interface {
	Save(file multipart.File, fileName string) (string, error)
}

type fileUseCase struct {
	repo repository.FileRepository
}

func (f *fileUseCase) Save(file multipart.File, fileName string) (string, error) {
	return f.repo.Save(file, fileName)
}

func NewFileUseCase(repo repository.FileRepository) FileUseCase {
	return &fileUseCase{repo: repo}
}
