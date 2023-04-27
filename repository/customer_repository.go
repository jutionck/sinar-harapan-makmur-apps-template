package repository

import (
	"fmt"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/utils/common"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
	BaseRepositoryPaging[model.Customer]
	BaseRepositoryEmailPhone[model.Customer]
	ListCustomerUser() ([]model.Customer, error)
	GetByUser(userId string) (*model.Customer, error)
	CreateCustomerVehicle(payload *model.Customer, association interface{}) error
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) CreateCustomerVehicle(payload *model.Customer, association interface{}) error {
	vehicle := association.(*model.Vehicle) // casting interface to struct vehicle
	result := c.db.Model(payload).Association("Vehicles").Append(vehicle)
	if result != nil {
		return result
	}
	return nil
}

func (c *customerRepository) Search(by map[string]interface{}) ([]model.Customer, error) {
	var customers []model.Customer
	result := c.db.Where(by).Find(&customers).Error
	if result != nil {
		return nil, result
	}
	return customers, nil
}

func (c *customerRepository) List() ([]model.Customer, error) {
	var customers []model.Customer
	result := c.db.Find(&customers).Error
	if result != nil {
		return nil, result
	}
	return customers, nil
}

func (c *customerRepository) Get(id string) (*model.Customer, error) {
	var customer model.Customer
	result := c.db.First(&customer, "id=?", id).Error
	if result != nil {
		return nil, result
	}
	return &customer, nil
}

func (c *customerRepository) ListCustomerUser() ([]model.Customer, error) {
	var customers []model.Customer
	result := c.db.Preload("UserCredential").Order("created_at").Find(&customers).Error
	if result != nil {
		return nil, result
	}

	return customers, nil
}

func (c *customerRepository) GetByUser(userId string) (*model.Customer, error) {
	var customer model.Customer
	result := c.db.Preload("UserCredential").First(&customer, "user_credential_id=?", userId).Error
	if result != nil {
		return nil, result
	}

	return &customer, nil
}

func (c *customerRepository) Save(payload *model.Customer) error {
	return c.db.Save(payload).Error
}

func (c *customerRepository) Delete(id string) error {
	return c.db.Delete(&model.Customer{}, "id=?", id).Error
}

func (c *customerRepository) GetByEmail(email string) (*model.Customer, error) {
	var customer model.Customer
	result := c.db.Select("id, email").First(&customer, "email=?", email).Error
	if result != nil {
		return nil, result
	}
	return &customer, nil
}

func (c *customerRepository) GetByPhone(phone string) (*model.Customer, error) {
	var customer model.Customer
	result := c.db.Select("id, phone_number").First(&customer, "phone_number=?", phone).Error
	if result != nil {
		return nil, result
	}
	return &customer, nil
}

func (c *customerRepository) Paging(requestQueryParams dto.RequestQueryParams) ([]model.Customer, dto.Paging, error) {
	paginationQuery, orderQuery := c.pagingValidate(requestQueryParams)
	var customers []model.Customer
	result := c.db.Preload("UserCredential").Order(orderQuery).Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Find(&customers).Error
	if result != nil {
		return nil, dto.Paging{}, result
	}
	var totalRows int64
	result = c.db.Model(&model.Customer{}).Count(&totalRows).Error
	if result != nil {
		return nil, dto.Paging{}, result
	}
	return customers, common.Paginate(paginationQuery.Page, paginationQuery.Take, int(totalRows)), nil
}

func (c *customerRepository) pagingValidate(requestQueryParams dto.RequestQueryParams) (dto.PaginationQuery, string) {
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

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}
