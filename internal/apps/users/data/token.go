package data

import (
	"github.com/realotz/whole/internal/services/users/biz"
)

func (ar *employeeRepo) CreateToken(employee *biz.Employee, expireTime int64) (string, error) {
	return ar.data.tk.Encode(employee.Account, employee.ID, "employee", expireTime)
}

func (ar *customerRepo) CreateToken(customer *biz.Customer, expireTime int64) (string, error) {
	return ar.data.tk.Encode(customer.Account, customer.ID, "customer", expireTime)
}
