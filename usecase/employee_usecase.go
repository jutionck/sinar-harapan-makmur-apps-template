package usecase

import (
	"fmt"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model/dto"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/utils"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
)

type EmployeeUseCase interface {
	BaseUseCase[model.Employee]
	BaseUseCaseEmailPhone[model.Employee]
	BaseUseCasePaging[model.Employee]
	FindAllEmployeeByManager(managerId string) ([]model.Employee, error)
}

type employeeUseCase struct {
	repo   repository.EmployeeRepository
	userUC UserUseCase
}

func (e *employeeUseCase) DeleteData(id string) error {
	employee, err := e.FindById(id)
	if err != nil {
		return fmt.Errorf("employee with ID %s not found", id)
	}
	return e.repo.Delete(employee.ID)
}

func (e *employeeUseCase) FindAll() ([]model.Employee, error) {
	return e.repo.List()
}

func (e *employeeUseCase) FindById(id string) (*model.Employee, error) {
	employee, err := e.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("employee with ID %s not found", id)
	}
	return employee, nil
}

func (e *employeeUseCase) SaveData(payload *model.Employee) error {
	if payload.ID != "" {
		_, err := e.FindById(payload.ID)
		if err != nil {
			return fmt.Errorf("employee with ID %s not found", payload.ID)
		}
	}

	if payload.ManagerID != nil {
		manager, _ := e.FindById(*payload.ManagerID)
		payload.Manager = manager
	}

	// create new user credential, set with default password ex: bod, 12345, password ...
	password, err := utils.HashPassword("password")
	if err != nil {
		return err
	}
	// define model user credential
	// Harus di cek dulu user sudah ada atau belum
	user, err := e.userUC.FindUserByUsername(payload.Email)
	if err != nil {
		return err
	}
	if user.ID != "" {
		payload.UserCredential = *user
		err := e.userUC.SaveData(user)
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

	return e.repo.Save(payload)
}

func (e *employeeUseCase) SearchBy(by map[string]interface{}) ([]model.Employee, error) {
	employees, err := e.repo.Search(by)
	if err != nil {
		return nil, fmt.Errorf("Data not found")
	}
	return employees, nil
}

func (e *employeeUseCase) FindByEmail(email string) (*model.Employee, error) {
	employee, err := e.repo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("Employee with email %s not found!", email)
	}
	return employee, nil
}

func (e *employeeUseCase) FindByPhone(phone string) (*model.Employee, error) {
	employee, err := e.repo.GetByPhone(phone)
	if err != nil {
		return nil, fmt.Errorf("Employee with phone number %s not found!", phone)
	}
	return employee, nil
}

func (e *employeeUseCase) FindAllEmployeeByManager(managerId string) ([]model.Employee, error) {
	return e.repo.ListEmployeeByManager(managerId)
}

func (e *employeeUseCase) Pagination(requestQueryParams dto.RequestQueryParams) ([]model.Employee, dto.Paging, error) {
	if !requestQueryParams.QueryParams.IsSortValid() {
		return nil, dto.Paging{}, fmt.Errorf("invalid sort by: %s", requestQueryParams.QueryParams.Sort)
	}
	return e.repo.Paging(requestQueryParams)
}

func NewEmployeeUseCase(repo repository.EmployeeRepository, userUC UserUseCase) EmployeeUseCase {
	return &employeeUseCase{repo: repo, userUC: userUC}
}
