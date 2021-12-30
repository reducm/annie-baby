// Code generated by entc, DO NOT EDIT.

package video

import (
	"annie-baby/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLink), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// OutputDir applies equality check predicate on the "output_dir" field. It's identical to OutputDirEQ.
func OutputDir(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutputDir), v))
	})
}

// Proxy applies equality check predicate on the "proxy" field. It's identical to ProxyEQ.
func Proxy(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProxy), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLink), v))
	})
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLink), v))
	})
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLink), v...))
	})
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLink), v...))
	})
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLink), v))
	})
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLink), v))
	})
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLink), v))
	})
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLink), v))
	})
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLink), v))
	})
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLink), v))
	})
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLink), v))
	})
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLink), v))
	})
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLink), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// OutputDirEQ applies the EQ predicate on the "output_dir" field.
func OutputDirEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOutputDir), v))
	})
}

// OutputDirNEQ applies the NEQ predicate on the "output_dir" field.
func OutputDirNEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOutputDir), v))
	})
}

// OutputDirIn applies the In predicate on the "output_dir" field.
func OutputDirIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOutputDir), v...))
	})
}

// OutputDirNotIn applies the NotIn predicate on the "output_dir" field.
func OutputDirNotIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOutputDir), v...))
	})
}

// OutputDirGT applies the GT predicate on the "output_dir" field.
func OutputDirGT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOutputDir), v))
	})
}

// OutputDirGTE applies the GTE predicate on the "output_dir" field.
func OutputDirGTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOutputDir), v))
	})
}

// OutputDirLT applies the LT predicate on the "output_dir" field.
func OutputDirLT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOutputDir), v))
	})
}

// OutputDirLTE applies the LTE predicate on the "output_dir" field.
func OutputDirLTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOutputDir), v))
	})
}

// OutputDirContains applies the Contains predicate on the "output_dir" field.
func OutputDirContains(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOutputDir), v))
	})
}

// OutputDirHasPrefix applies the HasPrefix predicate on the "output_dir" field.
func OutputDirHasPrefix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOutputDir), v))
	})
}

// OutputDirHasSuffix applies the HasSuffix predicate on the "output_dir" field.
func OutputDirHasSuffix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOutputDir), v))
	})
}

// OutputDirEqualFold applies the EqualFold predicate on the "output_dir" field.
func OutputDirEqualFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOutputDir), v))
	})
}

// OutputDirContainsFold applies the ContainsFold predicate on the "output_dir" field.
func OutputDirContainsFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOutputDir), v))
	})
}

// ProxyEQ applies the EQ predicate on the "proxy" field.
func ProxyEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProxy), v))
	})
}

// ProxyNEQ applies the NEQ predicate on the "proxy" field.
func ProxyNEQ(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProxy), v))
	})
}

// ProxyIn applies the In predicate on the "proxy" field.
func ProxyIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProxy), v...))
	})
}

// ProxyNotIn applies the NotIn predicate on the "proxy" field.
func ProxyNotIn(vs ...string) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProxy), v...))
	})
}

// ProxyGT applies the GT predicate on the "proxy" field.
func ProxyGT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProxy), v))
	})
}

// ProxyGTE applies the GTE predicate on the "proxy" field.
func ProxyGTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProxy), v))
	})
}

// ProxyLT applies the LT predicate on the "proxy" field.
func ProxyLT(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProxy), v))
	})
}

// ProxyLTE applies the LTE predicate on the "proxy" field.
func ProxyLTE(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProxy), v))
	})
}

// ProxyContains applies the Contains predicate on the "proxy" field.
func ProxyContains(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProxy), v))
	})
}

// ProxyHasPrefix applies the HasPrefix predicate on the "proxy" field.
func ProxyHasPrefix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProxy), v))
	})
}

// ProxyHasSuffix applies the HasSuffix predicate on the "proxy" field.
func ProxyHasSuffix(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProxy), v))
	})
}

// ProxyEqualFold applies the EqualFold predicate on the "proxy" field.
func ProxyEqualFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProxy), v))
	})
}

// ProxyContainsFold applies the ContainsFold predicate on the "proxy" field.
func ProxyContainsFold(v string) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProxy), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Video {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Video(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Video) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Video) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
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
func Not(p predicate.Video) predicate.Video {
	return predicate.Video(func(s *sql.Selector) {
		p(s.Not())
	})
}