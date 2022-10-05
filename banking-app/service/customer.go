package service

import (
	"go-banking/banking-app/domain"
	"go-banking/banking-app/dto"
	"go-banking/banking-app/utils"
)

type CustomerService interface {
	NewCustomer(dto.NewCustomerRequest) (*dto.CustomerResponse, *utils.AppError)
	GetAllCustomer(string) ([]dto.CustomerResponse, *utils.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *utils.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *utils.AppError) {
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, customer := range c {
		c := customer.ToDto()
		response = append(response, c)
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *utils.AppError) {
	c, err := s.repo.ByID(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultCustomerService) NewCustomer(r dto.NewCustomerRequest) (*dto.CustomerResponse, *utils.AppError) {
	err := r.IsValid()
	if err != nil {
		return nil, err
	}
	customer := domain.Customer{
		Name:        r.Name,
		City:        r.City,
		Zipcode:     r.Zipcode,
		DateofBirth: r.DateofBirth,
		Status:      "1",
	}
	c, err := s.repo.Create(customer)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
