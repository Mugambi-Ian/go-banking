package service

import (
	"go-banking/banking-app/domain"
	"go-banking/banking-app/dto"
	"go-banking/banking-app/utils"
	"time"
)

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}
type TransactionService interface {
	Transact(dto.TransactionRequest) (*dto.TransactionResponse, *utils.AppError)
}

func (s DefaultTransactionService) Transact(req dto.TransactionRequest) (*dto.TransactionResponse, *utils.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Transaction{
		AccountId:       req.AccountId,
		TransactionId:   "",
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05")}
	save, err := s.repo.Transact(a)
	if err != nil {
		return nil, err
	}
	response := save.ToTransactionResponseDTO()
	return &response, nil
}

func NewTransactionService(repository domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo: repository}
}
