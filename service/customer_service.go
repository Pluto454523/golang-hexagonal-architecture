package service

import (
	"banking/errs"
	"banking/logs"
	"banking/repository"
	"database/sql"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

// ** ใช้งานรวมกับ customerService เพื่อทำงานเป็น Constructor Function
func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	customers, err := s.custRepo.GetAll()
	if err != nil {
		// log.Println(err)
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custRepoData := []CustomerResponse{}
	for _, v := range customers {
		custRepoElement := CustomerResponse{
			CustomerID: v.CustomerID,
			Name:       v.Name,
			Status:     v.Status,
		}
		custRepoData = append(custRepoData, custRepoElement)
	}

	return custRepoData, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	cutomer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custRepoData := CustomerResponse{
		CustomerID: cutomer.CustomerID,
		Name:       cutomer.Name,
		Status:     cutomer.Status,
	}

	return &custRepoData, nil
}
