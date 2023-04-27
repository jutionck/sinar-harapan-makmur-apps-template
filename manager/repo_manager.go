package manager

import "github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"

// RepositoryManager -> all repo
type RepositoryManager interface {
	BrandRepo() repository.BrandRepository
	VehicleRepo() repository.VehicleRepository
	CustomerRepo() repository.CustomerRepository
	EmployeeRepo() repository.EmployeeRepository
	UserRepo() repository.UserRepository
	TransactionRepo() repository.TransactionRepository
	FileRepo() repository.FileRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) FileRepo() repository.FileRepository {
	return repository.NewFileRepository(r.infra.UploadPath())
}

func (r *repositoryManager) BrandRepo() repository.BrandRepository {
	return repository.NewBrandRepository(r.infra.Conn())
}

func (r *repositoryManager) VehicleRepo() repository.VehicleRepository {
	return repository.NewVehicleRepository(r.infra.Conn())
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Conn())
}

func (r *repositoryManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infra.Conn())
}

func (r *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func (r *repositoryManager) TransactionRepo() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.Conn())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{infra: infra}
}
