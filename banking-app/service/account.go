package service

import (
	"go-banking/banking-app/domain"
	"go-banking/banking-app/dto"
	"go-banking/banking-app/utils"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *utils.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRespository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *utils.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1"}
	save, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := save.ToNewAccountResponseDTO()
	return &response, nil
}

func NewAccountService(repository domain.AccountRespository) DefaultAccountService {
	return DefaultAccountService{repo: repository}
}
