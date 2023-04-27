package repository

import (
	"fmt"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/utils/common"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	BaseRepository[model.Vehicle]
	BaseRepositoryPaging[model.Vehicle]
	UpdateStock(count int, id string) error
}

type vehicleRepository struct {
	db *gorm.DB
}

func (v *vehicleRepository) Search(by map[string]interface{}) ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	result := v.db.Preload("Brand").Where(by).Find(&vehicles)
	if err := result.Error; err != nil {
		return vehicles, err
	}
	return vehicles, nil
}

func (v *vehicleRepository) List() ([]model.Vehicle, error) {
	var vehicles []model.Vehicle
	result := v.db.Preload("Brand").Find(&vehicles)
	if err := result.Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (v *vehicleRepository) Get(id string) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	result := v.db.Preload("Brand").First(&vehicle, "id = ?", id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (v *vehicleRepository) Save(payload *model.Vehicle) error {
	result := v.db.Save(payload)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (v *vehicleRepository) Delete(id string) error {
	return v.db.Delete(&model.Vehicle{}, "id=?", id).Error
}

func (v *vehicleRepository) UpdateStock(count int, id string) error {
	// UPDATE mst_vehicle SET stock = stock - {count} WHERE id = {id}
	result := v.db.Model(&model.Vehicle{}).Where("id=?", id).Update("stock", gorm.Expr("stock - ?", count))
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (v *vehicleRepository) Paging(requestQueryParams dto.RequestQueryParams) ([]model.Vehicle, dto.Paging, error) {
	paginationQuery, orderQuery := v.pagingValidate(requestQueryParams)
	var vehicles []model.Vehicle
	result := v.db.Preload("Brand").Order(orderQuery).Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Find(&vehicles).Error
	if result != nil {
		return nil, dto.Paging{}, result
	}
	var totalRows int64
	result = v.db.Model(&model.Brand{}).Count(&totalRows).Error
	if result != nil {
		return nil, dto.Paging{}, result
	}
	return vehicles, common.Paginate(paginationQuery.Page, paginationQuery.Take, int(totalRows)), nil
}

func (v *vehicleRepository) pagingValidate(requestQueryParams dto.RequestQueryParams) (dto.PaginationQuery, string) {
	var paginationQuery dto.PaginationQuery
	paginationQuery = common.GetPaginationParams(requestQueryParams.PaginationParam)
	orderQuery := "id"
	if requestQueryParams.QueryParams.Order != "" && requestQueryParams.QueryParams.Sort != "" {
		sorting := "ASC"
		if requestQueryParams.QueryParams.Sort == "desc" {
			sorting = "DESC"
		}
		orderQuery = fmt.Sprintf("%s %s", requestQueryParams.QueryParams.Order, sorting)
	}
	return paginationQuery, orderQuery
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
	return &vehicleRepository{db: db}
}
