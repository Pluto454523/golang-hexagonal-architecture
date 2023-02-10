package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDb struct {
	db *sqlx.DB
}

func NewAccountRepositoryDb(db *sqlx.DB) AccountRepository {

	return accountRepositoryDb{db: db}
}

func (accRepo accountRepositoryDb) Create(acc Account) (*Account, error) {

	query := "INSERT INTO accounts (`customer_id`, `opening_date`, `account_type`, `amount`, `status`) VALUES (?, ?, ?, ?, ?)"
	result, err := accRepo.db.Exec(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	acc.AccountID = int(id)

	return &acc, nil
}

func (accRepo accountRepositoryDb) GetAll(id int) ([]Account, error) {

	query := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id=?"
	accounts := []Account{}
	err := accRepo.db.Select(&accounts, query, id)

	if err != nil {
		return nil, err
	}

	return accounts, nil
}
