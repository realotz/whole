// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/realotz/whole/internal/apps/users/data/ent/migrate"

	"github.com/realotz/whole/internal/apps/users/data/ent/member"
	"github.com/realotz/whole/internal/apps/users/data/ent/wechat"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Member is the client for interacting with the Member builders.
	Member *MemberClient
	// Wechat is the client for interacting with the Wechat builders.
	Wechat *WechatClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Member = NewMemberClient(c.config)
	c.Wechat = NewWechatClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Member: NewMemberClient(cfg),
		Wechat: NewWechatClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config: cfg,
		Member: NewMemberClient(cfg),
		Wechat: NewWechatClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Member.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Member.Use(hooks...)
	c.Wechat.Use(hooks...)
}

// MemberClient is a client for the Member schema.
type MemberClient struct {
	config
}

// NewMemberClient returns a client for the Member from the given config.
func NewMemberClient(c config) *MemberClient {
	return &MemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `member.Hooks(f(g(h())))`.
func (c *MemberClient) Use(hooks ...Hook) {
	c.hooks.Member = append(c.hooks.Member, hooks...)
}

// Create returns a create builder for Member.
func (c *MemberClient) Create() *MemberCreate {
	mutation := newMemberMutation(c.config, OpCreate)
	return &MemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Member entities.
func (c *MemberClient) CreateBulk(builders ...*MemberCreate) *MemberCreateBulk {
	return &MemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Member.
func (c *MemberClient) Update() *MemberUpdate {
	mutation := newMemberMutation(c.config, OpUpdate)
	return &MemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MemberClient) UpdateOne(m *Member) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMember(m))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MemberClient) UpdateOneID(id int64) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMemberID(id))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Member.
func (c *MemberClient) Delete() *MemberDelete {
	mutation := newMemberMutation(c.config, OpDelete)
	return &MemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MemberClient) DeleteOne(m *Member) *MemberDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MemberClient) DeleteOneID(id int64) *MemberDeleteOne {
	builder := c.Delete().Where(member.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MemberDeleteOne{builder}
}

// Query returns a query builder for Member.
func (c *MemberClient) Query() *MemberQuery {
	return &MemberQuery{config: c.config}
}

// Get returns a Member entity by its id.
func (c *MemberClient) Get(ctx context.Context, id int64) (*Member, error) {
	return c.Query().Where(member.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MemberClient) GetX(ctx context.Context, id int64) *Member {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWechats queries the wechats edge of a Member.
func (c *MemberClient) QueryWechats(m *Member) *WechatQuery {
	query := &WechatQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(wechat.Table, wechat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.WechatsTable, member.WechatsColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MemberClient) Hooks() []Hook {
	return c.hooks.Member
}

// WechatClient is a client for the Wechat schema.
type WechatClient struct {
	config
}

// NewWechatClient returns a client for the Wechat from the given config.
func NewWechatClient(c config) *WechatClient {
	return &WechatClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `wechat.Hooks(f(g(h())))`.
func (c *WechatClient) Use(hooks ...Hook) {
	c.hooks.Wechat = append(c.hooks.Wechat, hooks...)
}

// Create returns a create builder for Wechat.
func (c *WechatClient) Create() *WechatCreate {
	mutation := newWechatMutation(c.config, OpCreate)
	return &WechatCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Wechat entities.
func (c *WechatClient) CreateBulk(builders ...*WechatCreate) *WechatCreateBulk {
	return &WechatCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Wechat.
func (c *WechatClient) Update() *WechatUpdate {
	mutation := newWechatMutation(c.config, OpUpdate)
	return &WechatUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WechatClient) UpdateOne(w *Wechat) *WechatUpdateOne {
	mutation := newWechatMutation(c.config, OpUpdateOne, withWechat(w))
	return &WechatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WechatClient) UpdateOneID(id int) *WechatUpdateOne {
	mutation := newWechatMutation(c.config, OpUpdateOne, withWechatID(id))
	return &WechatUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Wechat.
func (c *WechatClient) Delete() *WechatDelete {
	mutation := newWechatMutation(c.config, OpDelete)
	return &WechatDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WechatClient) DeleteOne(w *Wechat) *WechatDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WechatClient) DeleteOneID(id int) *WechatDeleteOne {
	builder := c.Delete().Where(wechat.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WechatDeleteOne{builder}
}

// Query returns a query builder for Wechat.
func (c *WechatClient) Query() *WechatQuery {
	return &WechatQuery{config: c.config}
}

// Get returns a Wechat entity by its id.
func (c *WechatClient) Get(ctx context.Context, id int) (*Wechat, error) {
	return c.Query().Where(wechat.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WechatClient) GetX(ctx context.Context, id int) *Wechat {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Wechat.
func (c *WechatClient) QueryOwner(w *Wechat) *MemberQuery {
	query := &MemberQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(wechat.Table, wechat.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, wechat.OwnerTable, wechat.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WechatClient) Hooks() []Hook {
	return c.hooks.Wechat
}
