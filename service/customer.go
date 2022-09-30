package service

import (
	"mugambi-ian/go-banking/domain"
	"mugambi-ian/go-banking/dto"
	"mugambi-ian/go-banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCusomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCusomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, customer := range c {
		response = append(response, customer.ToDto())
	}
	return response, nil
}

func (s DefaultCusomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ByID(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCusomerService {
	return DefaultCusomerService{repo: repository}
}
