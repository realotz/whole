// Code generated by entc, DO NOT EDIT.

package wechat

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/realotz/whole/internal/apps/users/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Openid applies equality check predicate on the "openid" field. It's identical to OpenidEQ.
func Openid(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOpenid), v))
	})
}

// UnionId applies equality check predicate on the "unionId" field. It's identical to UnionIdEQ.
func UnionId(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnionId), v))
	})
}

// MetaData applies equality check predicate on the "meta_data" field. It's identical to MetaDataEQ.
func MetaData(v []byte) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMetaData), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// OpenidEQ applies the EQ predicate on the "openid" field.
func OpenidEQ(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOpenid), v))
	})
}

// OpenidNEQ applies the NEQ predicate on the "openid" field.
func OpenidNEQ(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOpenid), v))
	})
}

// OpenidIn applies the In predicate on the "openid" field.
func OpenidIn(vs ...string) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOpenid), v...))
	})
}

// OpenidNotIn applies the NotIn predicate on the "openid" field.
func OpenidNotIn(vs ...string) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOpenid), v...))
	})
}

// OpenidGT applies the GT predicate on the "openid" field.
func OpenidGT(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOpenid), v))
	})
}

// OpenidGTE applies the GTE predicate on the "openid" field.
func OpenidGTE(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOpenid), v))
	})
}

// OpenidLT applies the LT predicate on the "openid" field.
func OpenidLT(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOpenid), v))
	})
}

// OpenidLTE applies the LTE predicate on the "openid" field.
func OpenidLTE(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOpenid), v))
	})
}

// OpenidContains applies the Contains predicate on the "openid" field.
func OpenidContains(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOpenid), v))
	})
}

// OpenidHasPrefix applies the HasPrefix predicate on the "openid" field.
func OpenidHasPrefix(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOpenid), v))
	})
}

// OpenidHasSuffix applies the HasSuffix predicate on the "openid" field.
func OpenidHasSuffix(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOpenid), v))
	})
}

// OpenidEqualFold applies the EqualFold predicate on the "openid" field.
func OpenidEqualFold(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOpenid), v))
	})
}

// OpenidContainsFold applies the ContainsFold predicate on the "openid" field.
func OpenidContainsFold(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOpenid), v))
	})
}

// UnionIdEQ applies the EQ predicate on the "unionId" field.
func UnionIdEQ(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnionId), v))
	})
}

// UnionIdNEQ applies the NEQ predicate on the "unionId" field.
func UnionIdNEQ(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnionId), v))
	})
}

// UnionIdIn applies the In predicate on the "unionId" field.
func UnionIdIn(vs ...string) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUnionId), v...))
	})
}

// UnionIdNotIn applies the NotIn predicate on the "unionId" field.
func UnionIdNotIn(vs ...string) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUnionId), v...))
	})
}

// UnionIdGT applies the GT predicate on the "unionId" field.
func UnionIdGT(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnionId), v))
	})
}

// UnionIdGTE applies the GTE predicate on the "unionId" field.
func UnionIdGTE(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnionId), v))
	})
}

// UnionIdLT applies the LT predicate on the "unionId" field.
func UnionIdLT(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnionId), v))
	})
}

// UnionIdLTE applies the LTE predicate on the "unionId" field.
func UnionIdLTE(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnionId), v))
	})
}

// UnionIdContains applies the Contains predicate on the "unionId" field.
func UnionIdContains(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUnionId), v))
	})
}

// UnionIdHasPrefix applies the HasPrefix predicate on the "unionId" field.
func UnionIdHasPrefix(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUnionId), v))
	})
}

// UnionIdHasSuffix applies the HasSuffix predicate on the "unionId" field.
func UnionIdHasSuffix(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUnionId), v))
	})
}

// UnionIdEqualFold applies the EqualFold predicate on the "unionId" field.
func UnionIdEqualFold(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUnionId), v))
	})
}

// UnionIdContainsFold applies the ContainsFold predicate on the "unionId" field.
func UnionIdContainsFold(v string) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUnionId), v))
	})
}

// AppTypeEQ applies the EQ predicate on the "app_type" field.
func AppTypeEQ(v AppType) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppType), v))
	})
}

// AppTypeNEQ applies the NEQ predicate on the "app_type" field.
func AppTypeNEQ(v AppType) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppType), v))
	})
}

// AppTypeIn applies the In predicate on the "app_type" field.
func AppTypeIn(vs ...AppType) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppType), v...))
	})
}

// AppTypeNotIn applies the NotIn predicate on the "app_type" field.
func AppTypeNotIn(vs ...AppType) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppType), v...))
	})
}

// MetaDataEQ applies the EQ predicate on the "meta_data" field.
func MetaDataEQ(v []byte) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMetaData), v))
	})
}

// MetaDataNEQ applies the NEQ predicate on the "meta_data" field.
func MetaDataNEQ(v []byte) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMetaData), v))
	})
}

// MetaDataIn applies the In predicate on the "meta_data" field.
func MetaDataIn(vs ...[]byte) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMetaData), v...))
	})
}

// MetaDataNotIn applies the NotIn predicate on the "meta_data" field.
func MetaDataNotIn(vs ...[]byte) predicate.Wechat {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Wechat(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMetaData), v...))
	})
}

// MetaDataGT applies the GT predicate on the "meta_data" field.
func MetaDataGT(v []byte) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMetaData), v))
	})
}

// MetaDataGTE applies the GTE predicate on the "meta_data" field.
func MetaDataGTE(v []byte) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMetaData), v))
	})
}

// MetaDataLT applies the LT predicate on the "meta_data" field.
func MetaDataLT(v []byte) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMetaData), v))
	})
}

// MetaDataLTE applies the LTE predicate on the "meta_data" field.
func MetaDataLTE(v []byte) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMetaData), v))
	})
}

// MetaDataIsNil applies the IsNil predicate on the "meta_data" field.
func MetaDataIsNil() predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMetaData)))
	})
}

// MetaDataNotNil applies the NotNil predicate on the "meta_data" field.
func MetaDataNotNil() predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMetaData)))
	})
}

// HasCustomers applies the HasEdge predicate on the "customers" edge.
func HasCustomers() predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomersTable, CustomersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomersWith applies the HasEdge predicate on the "customers" edge with a given conditions (other predicates).
func HasCustomersWith(preds ...predicate.Customer) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CustomersTable, CustomersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Wechat) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Wechat) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Wechat) predicate.Wechat {
	return predicate.Wechat(func(s *sql.Selector) {
		p(s.Not())
	})
}
