package service

import "mugambi-ian/go-banking/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type DefaultCusomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCusomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCusomerService {
	return DefaultCusomerService{repo: repository}
}
