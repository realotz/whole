package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/whole/internal/apps/admin/biz"
)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

// NewEmployeeRepo .
func NewRoleRepo(data *Data, logger log.Logger) biz.EmployeeRepo {
	return &employeeRepo{
		data: data,
		log:  log.NewHelper("employee_repo", logger),
	}
}
