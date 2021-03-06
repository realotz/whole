// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/realotz/whole/internal/apps/users/data/ent/customer"
	"github.com/realotz/whole/internal/apps/users/data/ent/wechat"
)

// Wechat is the model entity for the Wechat schema.
type Wechat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Openid holds the value of the "openid" field.
	// openid
	Openid string `json:"openid,omitempty"`
	// UnionId holds the value of the "unionId" field.
	// unionId
	UnionId string `json:"unionId,omitempty"`
	// AppType holds the value of the "app_type" field.
	// 类型 小程序 公众号
	AppType wechat.AppType `json:"app_type,omitempty"`
	// MetaData holds the value of the "meta_data" field.
	// 微信原始数据
	MetaData []byte `json:"meta_data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WechatQuery when eager-loading is set.
	Edges            WechatEdges `json:"edges"`
	customer_wechats *int64
}

// WechatEdges holds the relations/edges for other nodes in the graph.
type WechatEdges struct {
	// Customers holds the value of the customers edge.
	Customers *Customer `json:"customers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CustomersOrErr returns the Customers value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WechatEdges) CustomersOrErr() (*Customer, error) {
	if e.loadedTypes[0] {
		if e.Customers == nil {
			// The edge customers was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customers, nil
	}
	return nil, &NotLoadedError{edge: "customers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Wechat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case wechat.FieldMetaData:
			values[i] = &[]byte{}
		case wechat.FieldID:
			values[i] = &sql.NullInt64{}
		case wechat.FieldOpenid, wechat.FieldUnionId, wechat.FieldAppType:
			values[i] = &sql.NullString{}
		case wechat.FieldCreateTime, wechat.FieldUpdateTime:
			values[i] = &sql.NullTime{}
		case wechat.ForeignKeys[0]: // customer_wechats
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Wechat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Wechat fields.
func (w *Wechat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wechat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case wechat.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				w.CreateTime = value.Time
			}
		case wechat.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				w.UpdateTime = value.Time
			}
		case wechat.FieldOpenid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field openid", values[i])
			} else if value.Valid {
				w.Openid = value.String
			}
		case wechat.FieldUnionId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unionId", values[i])
			} else if value.Valid {
				w.UnionId = value.String
			}
		case wechat.FieldAppType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field app_type", values[i])
			} else if value.Valid {
				w.AppType = wechat.AppType(value.String)
			}
		case wechat.FieldMetaData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta_data", values[i])
			} else if value != nil {
				w.MetaData = *value
			}
		case wechat.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field customer_wechats", value)
			} else if value.Valid {
				w.customer_wechats = new(int64)
				*w.customer_wechats = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryCustomers queries the "customers" edge of the Wechat entity.
func (w *Wechat) QueryCustomers() *CustomerQuery {
	return (&WechatClient{config: w.config}).QueryCustomers(w)
}

// Update returns a builder for updating this Wechat.
// Note that you need to call Wechat.Unwrap() before calling this method if this Wechat
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Wechat) Update() *WechatUpdateOne {
	return (&WechatClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Wechat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Wechat) Unwrap() *Wechat {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Wechat is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Wechat) String() string {
	var builder strings.Builder
	builder.WriteString("Wechat(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(w.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(w.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", openid=")
	builder.WriteString(w.Openid)
	builder.WriteString(", unionId=")
	builder.WriteString(w.UnionId)
	builder.WriteString(", app_type=")
	builder.WriteString(fmt.Sprintf("%v", w.AppType))
	builder.WriteString(", meta_data=")
	builder.WriteString(fmt.Sprintf("%v", w.MetaData))
	builder.WriteByte(')')
	return builder.String()
}

// Wechats is a parsable slice of Wechat.
type Wechats []*Wechat

func (w Wechats) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
