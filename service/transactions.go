package service

import (
	"mugambi-ian/go-banking/domain"
	"mugambi-ian/go-banking/dto"
	"mugambi-ian/go-banking/errs"
	"time"
)

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}
type TransactionService interface {
	Transact(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

func (s DefaultTransactionService) Transact(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
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
	response := save.ToTransactionResponseDTO(float64(save.Balance))
	return &response, nil
}

func NewTransactionService(repository domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo: repository}
}
