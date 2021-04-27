package data

import (
	"github.com/realotz/whole/internal/apps/admin/biz"
)

func (ar *employeeRepo) CreateToken(employee *biz.Employee, expireTime int64) (string, error) {
	return ar.data.tk.Encode(employee.Account, employee.ID, "employee", expireTime)
}
