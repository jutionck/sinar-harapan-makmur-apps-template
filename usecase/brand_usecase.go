package usecase

import (
	"fmt"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
)

type BrandUseCase interface {
	BaseUseCase[model.Brand]
	BaseUseCasePaging[model.Brand]
}

type brandUseCase struct {
	repo repository.BrandRepository
}

func (b *brandUseCase) DeleteData(id string) error {
	brand, err := b.FindById(id)
	if err != nil {
		return fmt.Errorf("brand with ID %s not found", id)
	}
	return b.repo.Delete(brand.ID)
}

func (b *brandUseCase) FindAll() ([]model.Brand, error) {
	return b.repo.List()
}

func (b *brandUseCase) FindById(id string) (*model.Brand, error) {
	brand, err := b.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("brand with ID %s not found", id)
	}
	return brand, nil
}

func (b *brandUseCase) SaveData(payload *model.Brand) error {
	err := payload.Validate()
	if err != nil {
		return err
	}
	// cek jika data sudah ada -> count > 0
	if err := b.repo.CountData(payload.Name, payload.ID); err != nil {
		return err
	}

	if payload.ID != "" {
		_, err := b.FindById(payload.ID)
		if err != nil {
			return fmt.Errorf("brand with ID %s not found", payload.ID)
		}
	}
	return b.repo.Save(payload)
}

func (b *brandUseCase) SearchBy(by map[string]interface{}) ([]model.Brand, error) {
	brands, err := b.repo.Search(by)
	if err != nil {
		return nil, fmt.Errorf("data not found")
	}
	return brands, nil
}

func (b *brandUseCase) Pagination(requestQueryParams dto.RequestQueryParams) ([]model.Brand, dto.Paging, error) {
	if !requestQueryParams.QueryParams.IsSortValid() {
		return nil, dto.Paging{}, fmt.Errorf("invalid sort by: %s", requestQueryParams.QueryParams.Sort)
	}
	return b.repo.Paging(requestQueryParams)
}

func NewBrandUseCase(repo repository.BrandRepository) BrandUseCase {
	return &brandUseCase{repo: repo}
}
