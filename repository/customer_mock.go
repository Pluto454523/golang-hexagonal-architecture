package repository

import "errors"

type customerRepositoryMock struct {
	customer []Customer
}

// ** Constructor Function ส่งค่า customerRepositoryMock คืนผ่าน CustomerRepository
func NewCustomerRepositoryMock() CustomerRepository {
	customersData := []Customer{
		{CustomerID: 1001, Name: "Ashish", City: "New Delhi", ZipCode: "110011", DateOfBirth: "2000-01-01", Status: 1},
		{CustomerID: 1002, Name: "Rob", City: "New Delhi", ZipCode: "110011", DateOfBirth: "2000-01-01", Status: 0},
	}
	return customerRepositoryMock{customer: customersData}
}

// ** Implement GetAll() ([]Customer, error) เพื่อดึงข้อมูลแบบ mock จากด้านบนมาใช้
func (r customerRepositoryMock) GetAll() ([]Customer, error) {

	return r.customer, nil
}

// ** Implement GetAll() ([]Customer, error) เพื่อดึงข้อมูลแบบ mock จากด้านบนมาใช้โดยเช็ค id
func (r customerRepositoryMock) GetById(id int) (*Customer, error) {

	for _, customer := range r.customer {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
