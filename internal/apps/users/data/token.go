package data

import (
	"github.com/realotz/whole/internal/apps/users/biz"
)

func (ar *customerRepo) CreateToken(customer *biz.Customer, expireTime int64) (string, error) {
	return ar.data.tk.Encode(customer.Account, customer.ID, "customer", expireTime)
}
