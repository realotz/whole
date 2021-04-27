package data

import (
	"context"
	"fmt"
	casbinModel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/realotz/whole/internal/apps/users/data/ent/role"
)

func NewCasbinAdapter(d *Data, logger log.Logger) persist.Adapter {
	return &casbinAdapter{
		data: d,
		log:  log.NewHelper("casbin/rbac", logger),
	}
}

// CasbinAdapter casbin适配器
type casbinAdapter struct {
	data *Data
	log  *log.Helper
}

// LoadPolicy loads all policy rules from the storage.
func (a *casbinAdapter) LoadPolicy(model casbinModel.Model) error {
	ctx := context.Background()
	err := a.loadRolePolicy(ctx, model)
	if err != nil {
		a.log.Errorf("Load casbin role policy error: %s", err.Error())
		return err
	}
	err = a.loadUserPolicy(ctx, model)
	if err != nil {
		a.log.Errorf("Load casbin user policy error: %s", err.Error())
		return err
	}
	return nil
}

// 加载角色策略(p,role_id,path,method)
func (a *casbinAdapter) loadRolePolicy(ctx context.Context, m casbinModel.Model) error {
	roleResult, err := a.data.db.Role.Query().Where(role.StatusEQ(role.StatusOn)).All(ctx)
	if err != nil {
		return err
	} else if len(roleResult) == 0 {
		return nil
	}
	for _, item := range roleResult {
		permissions, err := item.QueryPermissions().All(ctx)
		if err == nil {
			for _, permission := range permissions {
				line := fmt.Sprintf("p,%d,%s,%s", item.ID, permission.Path, permission.Action)
				persist.LoadPolicyLine(line, m)
			}
		}
	}
	return nil
}

// 加载用户策略(g,user_id,role_id)
func (a *casbinAdapter) loadUserPolicy(ctx context.Context, m casbinModel.Model) error {
	userResult, err := a.data.db.Employee.Query().All(ctx)
	if err != nil {
		return err
	} else if len(userResult) > 0 {
		for _, user := range userResult {
			roles, err := user.QueryRoles().All(ctx)
			if err == nil {
				for _, r := range roles {
					line := fmt.Sprintf("g,%d,%d", user.ID, r.ID)
					persist.LoadPolicyLine(line, m)
				}
			}
		}
	}
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *casbinAdapter) SavePolicy(model casbinModel.Model) error {
	return nil
}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (a *casbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (a *casbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (a *casbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return nil
}
