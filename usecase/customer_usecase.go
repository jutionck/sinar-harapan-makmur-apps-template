package usecase

import (
	"fmt"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/utils"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
)

type CustomerUseCase interface {
	BaseUseCase[model.Customer]
	BaseUseCaseEmailPhone[model.Customer]
	BaseUseCasePaging[model.Customer]
	AppendCustomerVehicle(payload *model.Customer, association interface{}) error
}

type customerUseCase struct {
	repo   repository.CustomerRepository
	userUC UserUseCase
}

func (c *customerUseCase) DeleteData(id string) error {
	customer, err := c.FindById(id)
	if err != nil {
		return fmt.Errorf("Customer with ID %s not found!", id)
	}
	return c.repo.Delete(customer.ID)
}

func (c *customerUseCase) FindAll() ([]model.Customer, error) {
	return c.repo.List()
}

func (c *customerUseCase) FindById(id string) (*model.Customer, error) {
	customer, err := c.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("Customer with ID %s not found!", id)
	}
	return customer, nil
}

func (c *customerUseCase) SaveData(payload *model.Customer) error {
	if payload.ID != "" {
		_, err := c.FindById(payload.ID)
		if err != nil {
			return fmt.Errorf("customer with ID %s not found", payload.ID)
		}
	}

	// create new user credential, set with default password ex: bod, 12345, password ...
	password, err := utils.HashPassword("password")
	if err != nil {
		return err
	}
	// define model user credential
	// Harus di cek dulu user sudah ada atau belum
	user, err := c.userUC.FindUserByUsername(payload.Email)
	if err != nil {
		return err
	}
	if user.ID != "" {
		payload.UserCredential = *user
		err := c.userUC.SaveData(user)
		if err != nil {
			return err
		}
	} else {
		userCredential := model.UserCredential{
			UserName: payload.Email, // unique
			Password: password,
			IsActive: false,
		}
		payload.UserCredential = userCredential
	}
	return c.repo.Save(payload)
}

func (c *customerUseCase) SearchBy(by map[string]interface{}) ([]model.Customer, error) {
	customers, err := c.repo.Search(by)
	if err != nil {
		return nil, fmt.Errorf("Data not found")
	}
	return customers, nil
}

func (c *customerUseCase) FindByEmail(email string) (*model.Customer, error) {
	customer, err := c.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("Customer with email %s not found!", email)
	}
	return customer, nil
}

func (c *customerUseCase) FindByPhone(phone string) (*model.Customer, error) {
	customer, err := c.repo.GetByPhone(phone)
	if err != nil {
		return nil, fmt.Errorf("Customer with phone number %s not found!", phone)
	}
	return customer, nil
}

func (c *customerUseCase) Pagination(requestQueryParams dto.RequestQueryParams) ([]model.Customer, dto.Paging, error) {
	if !requestQueryParams.QueryParams.IsSortValid() {
		return nil, dto.Paging{}, fmt.Errorf("invalid sort by: %s", requestQueryParams.QueryParams.Sort)
	}
	return c.repo.Paging(requestQueryParams)
}

func (c *customerUseCase) AppendCustomerVehicle(payload *model.Customer, association interface{}) error {
	return c.repo.CreateCustomerVehicle(payload, association)
}

func NewCustomerUseCase(
	repo repository.CustomerRepository,
	userUC UserUseCase,
) CustomerUseCase {
	return &customerUseCase{
		repo:   repo,
		userUC: userUC,
	}
}
