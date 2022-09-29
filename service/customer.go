package service

import (
	"mugambi-ian/go-banking/domain"
	"mugambi-ian/go-banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCusomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCusomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll(status)
}
func (s DefaultCusomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ByID(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCusomerService {
	return DefaultCusomerService{repo: repository}
}
