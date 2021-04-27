// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/realotz/whole/internal/apps/cms/data/ent/category"
	"github.com/realotz/whole/internal/apps/cms/data/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCategory = "Category"
)

// CategoryMutation represents an operation that mutates the Category nodes in the graph.
type CategoryMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	create_time   *time.Time
	update_time   *time.Time
	name          *string
	pid           *int64
	addpid        *int64
	icon          *string
	desc          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Category, error)
	predicates    []predicate.Category
}

var _ ent.Mutation = (*CategoryMutation)(nil)

// categoryOption allows management of the mutation configuration using functional options.
type categoryOption func(*CategoryMutation)

// newCategoryMutation creates new mutation for the Category entity.
func newCategoryMutation(c config, op Op, opts ...categoryOption) *CategoryMutation {
	m := &CategoryMutation{
		config:        c,
		op:            op,
		typ:           TypeCategory,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCategoryID sets the ID field of the mutation.
func withCategoryID(id int64) categoryOption {
	return func(m *CategoryMutation) {
		var (
			err   error
			once  sync.Once
			value *Category
		)
		m.oldValue = func(ctx context.Context) (*Category, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Category.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCategory sets the old Category of the mutation.
func withCategory(node *Category) categoryOption {
	return func(m *CategoryMutation) {
		m.oldValue = func(context.Context) (*Category, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CategoryMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CategoryMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Category entities.
func (m *CategoryMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID
// is only available if it was provided to the builder.
func (m *CategoryMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCreateTime sets the "create_time" field.
func (m *CategoryMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *CategoryMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Category entity.
// If the Category object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CategoryMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *CategoryMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetUpdateTime sets the "update_time" field.
func (m *CategoryMutation) SetUpdateTime(t time.Time) {
	m.update_time = &t
}

// UpdateTime returns the value of the "update_time" field in the mutation.
func (m *CategoryMutation) UpdateTime() (r time.Time, exists bool) {
	v := m.update_time
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateTime returns the old "update_time" field's value of the Category entity.
// If the Category object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CategoryMutation) OldUpdateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateTime: %w", err)
	}
	return oldValue.UpdateTime, nil
}

// ResetUpdateTime resets all changes to the "update_time" field.
func (m *CategoryMutation) ResetUpdateTime() {
	m.update_time = nil
}

// SetName sets the "name" field.
func (m *CategoryMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *CategoryMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Category entity.
// If the Category object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CategoryMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ClearName clears the value of the "name" field.
func (m *CategoryMutation) ClearName() {
	m.name = nil
	m.clearedFields[category.FieldName] = struct{}{}
}

// NameCleared returns if the "name" field was cleared in this mutation.
func (m *CategoryMutation) NameCleared() bool {
	_, ok := m.clearedFields[category.FieldName]
	return ok
}

// ResetName resets all changes to the "name" field.
func (m *CategoryMutation) ResetName() {
	m.name = nil
	delete(m.clearedFields, category.FieldName)
}

// SetPid sets the "pid" field.
func (m *CategoryMutation) SetPid(i int64) {
	m.pid = &i
	m.addpid = nil
}

// Pid returns the value of the "pid" field in the mutation.
func (m *CategoryMutation) Pid() (r int64, exists bool) {
	v := m.pid
	if v == nil {
		return
	}
	return *v, true
}

// OldPid returns the old "pid" field's value of the Category entity.
// If the Category object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CategoryMutation) OldPid(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldPid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldPid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPid: %w", err)
	}
	return oldValue.Pid, nil
}

// AddPid adds i to the "pid" field.
func (m *CategoryMutation) AddPid(i int64) {
	if m.addpid != nil {
		*m.addpid += i
	} else {
		m.addpid = &i
	}
}

// AddedPid returns the value that was added to the "pid" field in this mutation.
func (m *CategoryMutation) AddedPid() (r int64, exists bool) {
	v := m.addpid
	if v == nil {
		return
	}
	return *v, true
}

// ClearPid clears the value of the "pid" field.
func (m *CategoryMutation) ClearPid() {
	m.pid = nil
	m.addpid = nil
	m.clearedFields[category.FieldPid] = struct{}{}
}

// PidCleared returns if the "pid" field was cleared in this mutation.
func (m *CategoryMutation) PidCleared() bool {
	_, ok := m.clearedFields[category.FieldPid]
	return ok
}

// ResetPid resets all changes to the "pid" field.
func (m *CategoryMutation) ResetPid() {
	m.pid = nil
	m.addpid = nil
	delete(m.clearedFields, category.FieldPid)
}

// SetIcon sets the "icon" field.
func (m *CategoryMutation) SetIcon(s string) {
	m.icon = &s
}

// Icon returns the value of the "icon" field in the mutation.
func (m *CategoryMutation) Icon() (r string, exists bool) {
	v := m.icon
	if v == nil {
		return
	}
	return *v, true
}

// OldIcon returns the old "icon" field's value of the Category entity.
// If the Category object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CategoryMutation) OldIcon(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldIcon is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldIcon requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIcon: %w", err)
	}
	return oldValue.Icon, nil
}

// ClearIcon clears the value of the "icon" field.
func (m *CategoryMutation) ClearIcon() {
	m.icon = nil
	m.clearedFields[category.FieldIcon] = struct{}{}
}

// IconCleared returns if the "icon" field was cleared in this mutation.
func (m *CategoryMutation) IconCleared() bool {
	_, ok := m.clearedFields[category.FieldIcon]
	return ok
}

// ResetIcon resets all changes to the "icon" field.
func (m *CategoryMutation) ResetIcon() {
	m.icon = nil
	delete(m.clearedFields, category.FieldIcon)
}

// SetDesc sets the "desc" field.
func (m *CategoryMutation) SetDesc(s string) {
	m.desc = &s
}

// Desc returns the value of the "desc" field in the mutation.
func (m *CategoryMutation) Desc() (r string, exists bool) {
	v := m.desc
	if v == nil {
		return
	}
	return *v, true
}

// OldDesc returns the old "desc" field's value of the Category entity.
// If the Category object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CategoryMutation) OldDesc(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDesc is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDesc requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDesc: %w", err)
	}
	return oldValue.Desc, nil
}

// ClearDesc clears the value of the "desc" field.
func (m *CategoryMutation) ClearDesc() {
	m.desc = nil
	m.clearedFields[category.FieldDesc] = struct{}{}
}

// DescCleared returns if the "desc" field was cleared in this mutation.
func (m *CategoryMutation) DescCleared() bool {
	_, ok := m.clearedFields[category.FieldDesc]
	return ok
}

// ResetDesc resets all changes to the "desc" field.
func (m *CategoryMutation) ResetDesc() {
	m.desc = nil
	delete(m.clearedFields, category.FieldDesc)
}

// Op returns the operation name.
func (m *CategoryMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Category).
func (m *CategoryMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CategoryMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.create_time != nil {
		fields = append(fields, category.FieldCreateTime)
	}
	if m.update_time != nil {
		fields = append(fields, category.FieldUpdateTime)
	}
	if m.name != nil {
		fields = append(fields, category.FieldName)
	}
	if m.pid != nil {
		fields = append(fields, category.FieldPid)
	}
	if m.icon != nil {
		fields = append(fields, category.FieldIcon)
	}
	if m.desc != nil {
		fields = append(fields, category.FieldDesc)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CategoryMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case category.FieldCreateTime:
		return m.CreateTime()
	case category.FieldUpdateTime:
		return m.UpdateTime()
	case category.FieldName:
		return m.Name()
	case category.FieldPid:
		return m.Pid()
	case category.FieldIcon:
		return m.Icon()
	case category.FieldDesc:
		return m.Desc()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CategoryMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case category.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case category.FieldUpdateTime:
		return m.OldUpdateTime(ctx)
	case category.FieldName:
		return m.OldName(ctx)
	case category.FieldPid:
		return m.OldPid(ctx)
	case category.FieldIcon:
		return m.OldIcon(ctx)
	case category.FieldDesc:
		return m.OldDesc(ctx)
	}
	return nil, fmt.Errorf("unknown Category field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CategoryMutation) SetField(name string, value ent.Value) error {
	switch name {
	case category.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case category.FieldUpdateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateTime(v)
		return nil
	case category.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case category.FieldPid:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPid(v)
		return nil
	case category.FieldIcon:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIcon(v)
		return nil
	case category.FieldDesc:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDesc(v)
		return nil
	}
	return fmt.Errorf("unknown Category field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CategoryMutation) AddedFields() []string {
	var fields []string
	if m.addpid != nil {
		fields = append(fields, category.FieldPid)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CategoryMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case category.FieldPid:
		return m.AddedPid()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CategoryMutation) AddField(name string, value ent.Value) error {
	switch name {
	case category.FieldPid:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPid(v)
		return nil
	}
	return fmt.Errorf("unknown Category numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CategoryMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(category.FieldName) {
		fields = append(fields, category.FieldName)
	}
	if m.FieldCleared(category.FieldPid) {
		fields = append(fields, category.FieldPid)
	}
	if m.FieldCleared(category.FieldIcon) {
		fields = append(fields, category.FieldIcon)
	}
	if m.FieldCleared(category.FieldDesc) {
		fields = append(fields, category.FieldDesc)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CategoryMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CategoryMutation) ClearField(name string) error {
	switch name {
	case category.FieldName:
		m.ClearName()
		return nil
	case category.FieldPid:
		m.ClearPid()
		return nil
	case category.FieldIcon:
		m.ClearIcon()
		return nil
	case category.FieldDesc:
		m.ClearDesc()
		return nil
	}
	return fmt.Errorf("unknown Category nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CategoryMutation) ResetField(name string) error {
	switch name {
	case category.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case category.FieldUpdateTime:
		m.ResetUpdateTime()
		return nil
	case category.FieldName:
		m.ResetName()
		return nil
	case category.FieldPid:
		m.ResetPid()
		return nil
	case category.FieldIcon:
		m.ResetIcon()
		return nil
	case category.FieldDesc:
		m.ResetDesc()
		return nil
	}
	return fmt.Errorf("unknown Category field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CategoryMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CategoryMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CategoryMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CategoryMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CategoryMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CategoryMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CategoryMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Category unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CategoryMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Category edge %s", name)
}
