package manager

import (
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
)

// UseCaseManager -> all use case
type UseCaseManager interface {
	BrandUseCase() usecase.BrandUseCase
	VehicleUseCase() usecase.VehicleUseCase
	CustomerUseCase() usecase.CustomerUseCase
	EmployeeUseCase() usecase.EmployeeUseCase
	UserUseCase() usecase.UserUseCase
	TransactionUseCase() usecase.TransactionUseCase
	FileUSeCase() usecase.FileUseCase
}

type useCaseManager struct {
	repoManger RepositoryManager
}

func (u *useCaseManager) FileUSeCase() usecase.FileUseCase {
	return usecase.NewFileUseCase(u.repoManger.FileRepo())
}

func (u *useCaseManager) BrandUseCase() usecase.BrandUseCase {
	return usecase.NewBrandUseCase(u.repoManger.BrandRepo())
}

func (u *useCaseManager) VehicleUseCase() usecase.VehicleUseCase {
	return usecase.NewVehicleUseCase(
		u.repoManger.VehicleRepo(),
		u.BrandUseCase(),
		u.FileUSeCase(),
	)
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repoManger.UserRepo())
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManger.CustomerRepo(), u.UserUseCase())
}

func (u *useCaseManager) EmployeeUseCase() usecase.EmployeeUseCase {
	return usecase.NewEmployeeUseCase(u.repoManger.EmployeeRepo(), u.UserUseCase())
}

func (u *useCaseManager) TransactionUseCase() usecase.TransactionUseCase {
	return usecase.NewTransactionUseCase(
		u.repoManger.TransactionRepo(),
		u.VehicleUseCase(),
		u.EmployeeUseCase(),
		u.CustomerUseCase(),
	)
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{repoManger: repoManager}
}
