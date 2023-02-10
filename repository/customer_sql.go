package repository

import (
	"github.com/jmoiron/sqlx"
)

type customerRepositoryDb struct {
	db *sqlx.DB
}

// ** ใช้งานรวมกับ customerRepositoryDb เพื่อทำงานเป็น Constructor Function
func NewCustomerRepositoryDb(db *sqlx.DB) CustomerRepository {

	return customerRepositoryDb{db: db}
}

// ** Implement GetAll() ([]Customer, error)
func (r customerRepositoryDb) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// ** Implement GetById(id int) (*Customer, error)
func (r customerRepositoryDb) GetById(id int) (*Customer, error) {

	customers := Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id=?"
	err := r.db.Get(&customers, query, id)
	if err != nil {
		return nil, err
	}
	return &customers, nil

}
