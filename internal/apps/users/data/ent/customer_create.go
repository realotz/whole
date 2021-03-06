// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/realotz/whole/internal/apps/users/data/ent/customer"
	"github.com/realotz/whole/internal/apps/users/data/ent/wechat"
)

// CustomerCreate is the builder for creating a Customer entity.
type CustomerCreate struct {
	config
	mutation *CustomerMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (cc *CustomerCreate) SetCreateTime(t time.Time) *CustomerCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableCreateTime(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CustomerCreate) SetUpdateTime(t time.Time) *CustomerCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableUpdateTime(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetSex sets the "sex" field.
func (cc *CustomerCreate) SetSex(c customer.Sex) *CustomerCreate {
	cc.mutation.SetSex(c)
	return cc
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableSex(c *customer.Sex) *CustomerCreate {
	if c != nil {
		cc.SetSex(*c)
	}
	return cc
}

// SetUUID sets the "uuid" field.
func (cc *CustomerCreate) SetUUID(u uuid.UUID) *CustomerCreate {
	cc.mutation.SetUUID(u)
	return cc
}

// SetAccount sets the "account" field.
func (cc *CustomerCreate) SetAccount(s string) *CustomerCreate {
	cc.mutation.SetAccount(s)
	return cc
}

// SetAvatar sets the "avatar" field.
func (cc *CustomerCreate) SetAvatar(s string) *CustomerCreate {
	cc.mutation.SetAvatar(s)
	return cc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableAvatar(s *string) *CustomerCreate {
	if s != nil {
		cc.SetAvatar(*s)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CustomerCreate) SetName(s string) *CustomerCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableName(s *string) *CustomerCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetNickName sets the "nick_name" field.
func (cc *CustomerCreate) SetNickName(s string) *CustomerCreate {
	cc.mutation.SetNickName(s)
	return cc
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableNickName(s *string) *CustomerCreate {
	if s != nil {
		cc.SetNickName(*s)
	}
	return cc
}

// SetEmail sets the "email" field.
func (cc *CustomerCreate) SetEmail(s string) *CustomerCreate {
	cc.mutation.SetEmail(s)
	return cc
}

// SetMobile sets the "mobile" field.
func (cc *CustomerCreate) SetMobile(s string) *CustomerCreate {
	cc.mutation.SetMobile(s)
	return cc
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableMobile(s *string) *CustomerCreate {
	if s != nil {
		cc.SetMobile(*s)
	}
	return cc
}

// SetIDCard sets the "id_card" field.
func (cc *CustomerCreate) SetIDCard(s string) *CustomerCreate {
	cc.mutation.SetIDCard(s)
	return cc
}

// SetNillableIDCard sets the "id_card" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableIDCard(s *string) *CustomerCreate {
	if s != nil {
		cc.SetIDCard(*s)
	}
	return cc
}

// SetBirthday sets the "birthday" field.
func (cc *CustomerCreate) SetBirthday(t time.Time) *CustomerCreate {
	cc.mutation.SetBirthday(t)
	return cc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableBirthday(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetBirthday(*t)
	}
	return cc
}

// SetPassword sets the "password" field.
func (cc *CustomerCreate) SetPassword(s string) *CustomerCreate {
	cc.mutation.SetPassword(s)
	return cc
}

// SetSalt sets the "salt" field.
func (cc *CustomerCreate) SetSalt(s string) *CustomerCreate {
	cc.mutation.SetSalt(s)
	return cc
}

// SetLastIP sets the "last_ip" field.
func (cc *CustomerCreate) SetLastIP(s string) *CustomerCreate {
	cc.mutation.SetLastIP(s)
	return cc
}

// SetNillableLastIP sets the "last_ip" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableLastIP(s *string) *CustomerCreate {
	if s != nil {
		cc.SetLastIP(*s)
	}
	return cc
}

// SetLastTime sets the "last_time" field.
func (cc *CustomerCreate) SetLastTime(t time.Time) *CustomerCreate {
	cc.mutation.SetLastTime(t)
	return cc
}

// SetNillableLastTime sets the "last_time" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableLastTime(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetLastTime(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CustomerCreate) SetID(i int64) *CustomerCreate {
	cc.mutation.SetID(i)
	return cc
}

// AddWechatIDs adds the "wechats" edge to the Wechat entity by IDs.
func (cc *CustomerCreate) AddWechatIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddWechatIDs(ids...)
	return cc
}

// AddWechats adds the "wechats" edges to the Wechat entity.
func (cc *CustomerCreate) AddWechats(w ...*Wechat) *CustomerCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cc.AddWechatIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cc *CustomerCreate) Mutation() *CustomerMutation {
	return cc.mutation
}

// Save creates the Customer in the database.
func (cc *CustomerCreate) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CustomerCreate) SaveX(ctx context.Context) *Customer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cc *CustomerCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := customer.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := customer.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.Sex(); !ok {
		v := customer.DefaultSex
		cc.mutation.SetSex(v)
	}
	if _, ok := cc.mutation.UUID(); !ok {
		v := customer.DefaultUUID()
		cc.mutation.SetUUID(v)
	}
	if _, ok := cc.mutation.Name(); !ok {
		v := customer.DefaultName
		cc.mutation.SetName(v)
	}
	if _, ok := cc.mutation.NickName(); !ok {
		v := customer.DefaultNickName
		cc.mutation.SetNickName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CustomerCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := cc.mutation.Sex(); !ok {
		return &ValidationError{Name: "sex", err: errors.New("ent: missing required field \"sex\"")}
	}
	if v, ok := cc.mutation.Sex(); ok {
		if err := customer.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf("ent: validator failed for field \"sex\": %w", err)}
		}
	}
	if _, ok := cc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New("ent: missing required field \"uuid\"")}
	}
	if _, ok := cc.mutation.Account(); !ok {
		return &ValidationError{Name: "account", err: errors.New("ent: missing required field \"account\"")}
	}
	if v, ok := cc.mutation.Account(); ok {
		if err := customer.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf("ent: validator failed for field \"account\": %w", err)}
		}
	}
	if _, ok := cc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if _, ok := cc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if v, ok := cc.mutation.Password(); ok {
		if err := customer.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if _, ok := cc.mutation.Salt(); !ok {
		return &ValidationError{Name: "salt", err: errors.New("ent: missing required field \"salt\"")}
	}
	if v, ok := cc.mutation.Salt(); ok {
		if err := customer.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf("ent: validator failed for field \"salt\": %w", err)}
		}
	}
	return nil
}

func (cc *CustomerCreate) sqlSave(ctx context.Context) (*Customer, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (cc *CustomerCreate) createSpec() (*Customer, *sqlgraph.CreateSpec) {
	var (
		_node = &Customer{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: customer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: customer.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.Sex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: customer.FieldSex,
		})
		_node.Sex = value
	}
	if value, ok := cc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: customer.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := cc.mutation.Account(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAccount,
		})
		_node.Account = value
	}
	if value, ok := cc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldAvatar,
		})
		_node.Avatar = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.NickName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldNickName,
		})
		_node.NickName = value
	}
	if value, ok := cc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := cc.mutation.Mobile(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldMobile,
		})
		_node.Mobile = value
	}
	if value, ok := cc.mutation.IDCard(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldIDCard,
		})
		_node.IDCard = value
	}
	if value, ok := cc.mutation.Birthday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldBirthday,
		})
		_node.Birthday = value
	}
	if value, ok := cc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := cc.mutation.Salt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldSalt,
		})
		_node.Salt = value
	}
	if value, ok := cc.mutation.LastIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldLastIP,
		})
		_node.LastIP = value
	}
	if value, ok := cc.mutation.LastTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: customer.FieldLastTime,
		})
		_node.LastTime = value
	}
	if nodes := cc.mutation.WechatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.WechatsTable,
			Columns: []string{customer.WechatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wechat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomerCreateBulk is the builder for creating many Customer entities in bulk.
type CustomerCreateBulk struct {
	config
	builders []*CustomerCreate
}

// Save creates the Customer entities in the database.
func (ccb *CustomerCreateBulk) Save(ctx context.Context) ([]*Customer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Customer, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CustomerCreateBulk) SaveX(ctx context.Context) []*Customer {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
