// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/realotz/whole/internal/apps/users/data/ent/member"
	"github.com/realotz/whole/internal/apps/users/data/ent/predicate"
	"github.com/realotz/whole/internal/apps/users/data/ent/wechat"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where adds a new predicate for the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.predicates = append(mu.mutation.predicates, ps...)
	return mu
}

// SetSex sets the "sex" field.
func (mu *MemberUpdate) SetSex(m member.Sex) *MemberUpdate {
	mu.mutation.SetSex(m)
	return mu
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableSex(m *member.Sex) *MemberUpdate {
	if m != nil {
		mu.SetSex(*m)
	}
	return mu
}

// SetUUID sets the "uuid" field.
func (mu *MemberUpdate) SetUUID(u uuid.UUID) *MemberUpdate {
	mu.mutation.SetUUID(u)
	return mu
}

// SetAccount sets the "account" field.
func (mu *MemberUpdate) SetAccount(s string) *MemberUpdate {
	mu.mutation.SetAccount(s)
	return mu
}

// SetName sets the "name" field.
func (mu *MemberUpdate) SetName(s string) *MemberUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableName(s *string) *MemberUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// ClearName clears the value of the "name" field.
func (mu *MemberUpdate) ClearName() *MemberUpdate {
	mu.mutation.ClearName()
	return mu
}

// SetRole sets the "role" field.
func (mu *MemberUpdate) SetRole(s string) *MemberUpdate {
	mu.mutation.SetRole(s)
	return mu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableRole(s *string) *MemberUpdate {
	if s != nil {
		mu.SetRole(*s)
	}
	return mu
}

// ClearRole clears the value of the "role" field.
func (mu *MemberUpdate) ClearRole() *MemberUpdate {
	mu.mutation.ClearRole()
	return mu
}

// SetNickName sets the "nick_name" field.
func (mu *MemberUpdate) SetNickName(s string) *MemberUpdate {
	mu.mutation.SetNickName(s)
	return mu
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableNickName(s *string) *MemberUpdate {
	if s != nil {
		mu.SetNickName(*s)
	}
	return mu
}

// ClearNickName clears the value of the "nick_name" field.
func (mu *MemberUpdate) ClearNickName() *MemberUpdate {
	mu.mutation.ClearNickName()
	return mu
}

// SetEmail sets the "email" field.
func (mu *MemberUpdate) SetEmail(s string) *MemberUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetMobile sets the "mobile" field.
func (mu *MemberUpdate) SetMobile(s string) *MemberUpdate {
	mu.mutation.SetMobile(s)
	return mu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableMobile(s *string) *MemberUpdate {
	if s != nil {
		mu.SetMobile(*s)
	}
	return mu
}

// ClearMobile clears the value of the "mobile" field.
func (mu *MemberUpdate) ClearMobile() *MemberUpdate {
	mu.mutation.ClearMobile()
	return mu
}

// SetIDCard sets the "id_card" field.
func (mu *MemberUpdate) SetIDCard(s string) *MemberUpdate {
	mu.mutation.SetIDCard(s)
	return mu
}

// SetNillableIDCard sets the "id_card" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableIDCard(s *string) *MemberUpdate {
	if s != nil {
		mu.SetIDCard(*s)
	}
	return mu
}

// ClearIDCard clears the value of the "id_card" field.
func (mu *MemberUpdate) ClearIDCard() *MemberUpdate {
	mu.mutation.ClearIDCard()
	return mu
}

// SetBirthday sets the "birthday" field.
func (mu *MemberUpdate) SetBirthday(t time.Time) *MemberUpdate {
	mu.mutation.SetBirthday(t)
	return mu
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableBirthday(t *time.Time) *MemberUpdate {
	if t != nil {
		mu.SetBirthday(*t)
	}
	return mu
}

// ClearBirthday clears the value of the "birthday" field.
func (mu *MemberUpdate) ClearBirthday() *MemberUpdate {
	mu.mutation.ClearBirthday()
	return mu
}

// SetPassword sets the "password" field.
func (mu *MemberUpdate) SetPassword(s string) *MemberUpdate {
	mu.mutation.SetPassword(s)
	return mu
}

// SetSalt sets the "salt" field.
func (mu *MemberUpdate) SetSalt(s string) *MemberUpdate {
	mu.mutation.SetSalt(s)
	return mu
}

// SetLastIP sets the "last_ip" field.
func (mu *MemberUpdate) SetLastIP(s string) *MemberUpdate {
	mu.mutation.SetLastIP(s)
	return mu
}

// SetNillableLastIP sets the "last_ip" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableLastIP(s *string) *MemberUpdate {
	if s != nil {
		mu.SetLastIP(*s)
	}
	return mu
}

// ClearLastIP clears the value of the "last_ip" field.
func (mu *MemberUpdate) ClearLastIP() *MemberUpdate {
	mu.mutation.ClearLastIP()
	return mu
}

// SetLastTime sets the "last_time" field.
func (mu *MemberUpdate) SetLastTime(t time.Time) *MemberUpdate {
	mu.mutation.SetLastTime(t)
	return mu
}

// SetNillableLastTime sets the "last_time" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableLastTime(t *time.Time) *MemberUpdate {
	if t != nil {
		mu.SetLastTime(*t)
	}
	return mu
}

// ClearLastTime clears the value of the "last_time" field.
func (mu *MemberUpdate) ClearLastTime() *MemberUpdate {
	mu.mutation.ClearLastTime()
	return mu
}

// AddWechatIDs adds the "wechats" edge to the Wechat entity by IDs.
func (mu *MemberUpdate) AddWechatIDs(ids ...int) *MemberUpdate {
	mu.mutation.AddWechatIDs(ids...)
	return mu
}

// AddWechats adds the "wechats" edges to the Wechat entity.
func (mu *MemberUpdate) AddWechats(w ...*Wechat) *MemberUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return mu.AddWechatIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// ClearWechats clears all "wechats" edges to the Wechat entity.
func (mu *MemberUpdate) ClearWechats() *MemberUpdate {
	mu.mutation.ClearWechats()
	return mu
}

// RemoveWechatIDs removes the "wechats" edge to Wechat entities by IDs.
func (mu *MemberUpdate) RemoveWechatIDs(ids ...int) *MemberUpdate {
	mu.mutation.RemoveWechatIDs(ids...)
	return mu
}

// RemoveWechats removes "wechats" edges to Wechat entities.
func (mu *MemberUpdate) RemoveWechats(w ...*Wechat) *MemberUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return mu.RemoveWechatIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	mu.defaults()
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MemberUpdate) defaults() {
	if _, ok := mu.mutation.UpdateTime(); !ok {
		v := member.UpdateDefaultUpdateTime()
		mu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MemberUpdate) check() error {
	if v, ok := mu.mutation.Sex(); ok {
		if err := member.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf("ent: validator failed for field \"sex\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Account(); ok {
		if err := member.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf("ent: validator failed for field \"account\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Password(); ok {
		if err := member.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Salt(); ok {
		if err := member.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf("ent: validator failed for field \"salt\": %w", err)}
		}
	}
	return nil
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: member.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldUpdateTime,
		})
	}
	if value, ok := mu.mutation.Sex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: member.FieldSex,
		})
	}
	if value, ok := mu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: member.FieldUUID,
		})
	}
	if value, ok := mu.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldAccount,
		})
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldName,
		})
	}
	if mu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldName,
		})
	}
	if value, ok := mu.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldRole,
		})
	}
	if mu.mutation.RoleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldRole,
		})
	}
	if value, ok := mu.mutation.NickName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldNickName,
		})
	}
	if mu.mutation.NickNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldNickName,
		})
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldEmail,
		})
	}
	if value, ok := mu.mutation.Mobile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldMobile,
		})
	}
	if mu.mutation.MobileCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldMobile,
		})
	}
	if value, ok := mu.mutation.IDCard(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldIDCard,
		})
	}
	if mu.mutation.IDCardCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldIDCard,
		})
	}
	if value, ok := mu.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldBirthday,
		})
	}
	if mu.mutation.BirthdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: member.FieldBirthday,
		})
	}
	if value, ok := mu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldPassword,
		})
	}
	if value, ok := mu.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldSalt,
		})
	}
	if value, ok := mu.mutation.LastIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldLastIP,
		})
	}
	if mu.mutation.LastIPCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldLastIP,
		})
	}
	if value, ok := mu.mutation.LastTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldLastTime,
		})
	}
	if mu.mutation.LastTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: member.FieldLastTime,
		})
	}
	if mu.mutation.WechatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.WechatsTable,
			Columns: []string{member.WechatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wechat.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedWechatsIDs(); len(nodes) > 0 && !mu.mutation.WechatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.WechatsTable,
			Columns: []string{member.WechatsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.WechatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.WechatsTable,
			Columns: []string{member.WechatsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// SetSex sets the "sex" field.
func (muo *MemberUpdateOne) SetSex(m member.Sex) *MemberUpdateOne {
	muo.mutation.SetSex(m)
	return muo
}

// SetNillableSex sets the "sex" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableSex(m *member.Sex) *MemberUpdateOne {
	if m != nil {
		muo.SetSex(*m)
	}
	return muo
}

// SetUUID sets the "uuid" field.
func (muo *MemberUpdateOne) SetUUID(u uuid.UUID) *MemberUpdateOne {
	muo.mutation.SetUUID(u)
	return muo
}

// SetAccount sets the "account" field.
func (muo *MemberUpdateOne) SetAccount(s string) *MemberUpdateOne {
	muo.mutation.SetAccount(s)
	return muo
}

// SetName sets the "name" field.
func (muo *MemberUpdateOne) SetName(s string) *MemberUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableName(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// ClearName clears the value of the "name" field.
func (muo *MemberUpdateOne) ClearName() *MemberUpdateOne {
	muo.mutation.ClearName()
	return muo
}

// SetRole sets the "role" field.
func (muo *MemberUpdateOne) SetRole(s string) *MemberUpdateOne {
	muo.mutation.SetRole(s)
	return muo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableRole(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetRole(*s)
	}
	return muo
}

// ClearRole clears the value of the "role" field.
func (muo *MemberUpdateOne) ClearRole() *MemberUpdateOne {
	muo.mutation.ClearRole()
	return muo
}

// SetNickName sets the "nick_name" field.
func (muo *MemberUpdateOne) SetNickName(s string) *MemberUpdateOne {
	muo.mutation.SetNickName(s)
	return muo
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableNickName(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetNickName(*s)
	}
	return muo
}

// ClearNickName clears the value of the "nick_name" field.
func (muo *MemberUpdateOne) ClearNickName() *MemberUpdateOne {
	muo.mutation.ClearNickName()
	return muo
}

// SetEmail sets the "email" field.
func (muo *MemberUpdateOne) SetEmail(s string) *MemberUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetMobile sets the "mobile" field.
func (muo *MemberUpdateOne) SetMobile(s string) *MemberUpdateOne {
	muo.mutation.SetMobile(s)
	return muo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableMobile(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetMobile(*s)
	}
	return muo
}

// ClearMobile clears the value of the "mobile" field.
func (muo *MemberUpdateOne) ClearMobile() *MemberUpdateOne {
	muo.mutation.ClearMobile()
	return muo
}

// SetIDCard sets the "id_card" field.
func (muo *MemberUpdateOne) SetIDCard(s string) *MemberUpdateOne {
	muo.mutation.SetIDCard(s)
	return muo
}

// SetNillableIDCard sets the "id_card" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableIDCard(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetIDCard(*s)
	}
	return muo
}

// ClearIDCard clears the value of the "id_card" field.
func (muo *MemberUpdateOne) ClearIDCard() *MemberUpdateOne {
	muo.mutation.ClearIDCard()
	return muo
}

// SetBirthday sets the "birthday" field.
func (muo *MemberUpdateOne) SetBirthday(t time.Time) *MemberUpdateOne {
	muo.mutation.SetBirthday(t)
	return muo
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableBirthday(t *time.Time) *MemberUpdateOne {
	if t != nil {
		muo.SetBirthday(*t)
	}
	return muo
}

// ClearBirthday clears the value of the "birthday" field.
func (muo *MemberUpdateOne) ClearBirthday() *MemberUpdateOne {
	muo.mutation.ClearBirthday()
	return muo
}

// SetPassword sets the "password" field.
func (muo *MemberUpdateOne) SetPassword(s string) *MemberUpdateOne {
	muo.mutation.SetPassword(s)
	return muo
}

// SetSalt sets the "salt" field.
func (muo *MemberUpdateOne) SetSalt(s string) *MemberUpdateOne {
	muo.mutation.SetSalt(s)
	return muo
}

// SetLastIP sets the "last_ip" field.
func (muo *MemberUpdateOne) SetLastIP(s string) *MemberUpdateOne {
	muo.mutation.SetLastIP(s)
	return muo
}

// SetNillableLastIP sets the "last_ip" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableLastIP(s *string) *MemberUpdateOne {
	if s != nil {
		muo.SetLastIP(*s)
	}
	return muo
}

// ClearLastIP clears the value of the "last_ip" field.
func (muo *MemberUpdateOne) ClearLastIP() *MemberUpdateOne {
	muo.mutation.ClearLastIP()
	return muo
}

// SetLastTime sets the "last_time" field.
func (muo *MemberUpdateOne) SetLastTime(t time.Time) *MemberUpdateOne {
	muo.mutation.SetLastTime(t)
	return muo
}

// SetNillableLastTime sets the "last_time" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableLastTime(t *time.Time) *MemberUpdateOne {
	if t != nil {
		muo.SetLastTime(*t)
	}
	return muo
}

// ClearLastTime clears the value of the "last_time" field.
func (muo *MemberUpdateOne) ClearLastTime() *MemberUpdateOne {
	muo.mutation.ClearLastTime()
	return muo
}

// AddWechatIDs adds the "wechats" edge to the Wechat entity by IDs.
func (muo *MemberUpdateOne) AddWechatIDs(ids ...int) *MemberUpdateOne {
	muo.mutation.AddWechatIDs(ids...)
	return muo
}

// AddWechats adds the "wechats" edges to the Wechat entity.
func (muo *MemberUpdateOne) AddWechats(w ...*Wechat) *MemberUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return muo.AddWechatIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// ClearWechats clears all "wechats" edges to the Wechat entity.
func (muo *MemberUpdateOne) ClearWechats() *MemberUpdateOne {
	muo.mutation.ClearWechats()
	return muo
}

// RemoveWechatIDs removes the "wechats" edge to Wechat entities by IDs.
func (muo *MemberUpdateOne) RemoveWechatIDs(ids ...int) *MemberUpdateOne {
	muo.mutation.RemoveWechatIDs(ids...)
	return muo
}

// RemoveWechats removes "wechats" edges to Wechat entities.
func (muo *MemberUpdateOne) RemoveWechats(w ...*Wechat) *MemberUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return muo.RemoveWechatIDs(ids...)
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	var (
		err  error
		node *Member
	)
	muo.defaults()
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MemberUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdateTime(); !ok {
		v := member.UpdateDefaultUpdateTime()
		muo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MemberUpdateOne) check() error {
	if v, ok := muo.mutation.Sex(); ok {
		if err := member.SexValidator(v); err != nil {
			return &ValidationError{Name: "sex", err: fmt.Errorf("ent: validator failed for field \"sex\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Account(); ok {
		if err := member.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf("ent: validator failed for field \"account\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Password(); ok {
		if err := member.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Salt(); ok {
		if err := member.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf("ent: validator failed for field \"salt\": %w", err)}
		}
	}
	return nil
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: member.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Member.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldUpdateTime,
		})
	}
	if value, ok := muo.mutation.Sex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: member.FieldSex,
		})
	}
	if value, ok := muo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: member.FieldUUID,
		})
	}
	if value, ok := muo.mutation.Account(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldAccount,
		})
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldName,
		})
	}
	if muo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldName,
		})
	}
	if value, ok := muo.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldRole,
		})
	}
	if muo.mutation.RoleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldRole,
		})
	}
	if value, ok := muo.mutation.NickName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldNickName,
		})
	}
	if muo.mutation.NickNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldNickName,
		})
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldEmail,
		})
	}
	if value, ok := muo.mutation.Mobile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldMobile,
		})
	}
	if muo.mutation.MobileCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldMobile,
		})
	}
	if value, ok := muo.mutation.IDCard(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldIDCard,
		})
	}
	if muo.mutation.IDCardCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldIDCard,
		})
	}
	if value, ok := muo.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldBirthday,
		})
	}
	if muo.mutation.BirthdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: member.FieldBirthday,
		})
	}
	if value, ok := muo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldPassword,
		})
	}
	if value, ok := muo.mutation.Salt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldSalt,
		})
	}
	if value, ok := muo.mutation.LastIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldLastIP,
		})
	}
	if muo.mutation.LastIPCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: member.FieldLastIP,
		})
	}
	if value, ok := muo.mutation.LastTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: member.FieldLastTime,
		})
	}
	if muo.mutation.LastTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: member.FieldLastTime,
		})
	}
	if muo.mutation.WechatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.WechatsTable,
			Columns: []string{member.WechatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wechat.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedWechatsIDs(); len(nodes) > 0 && !muo.mutation.WechatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.WechatsTable,
			Columns: []string{member.WechatsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.WechatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.WechatsTable,
			Columns: []string{member.WechatsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
