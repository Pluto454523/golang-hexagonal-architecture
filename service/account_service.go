package service

import (
	"banking/errs"
	"banking/logs"
	"banking/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (accSrv accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {

	//Validate
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5,000")
	}
	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := accSrv.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	reponse := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}
	return &reponse, nil
}

func (accSrv accountService) GetAccount(customerID int) ([]AccountResponse, error) {

	accountS, err := accSrv.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	reponse := []AccountResponse{}
	for _, account := range accountS {
		reponse = append(reponse, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return reponse, nil
}
