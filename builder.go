package sqlc

import (
	"strconv"
	"strings"
)

type (
	Builder struct {
		filters       []filter
		order         string
		offset, limit int
	}

	filter struct {
		expression string
		args       []interface{}
	}
)

func (b *Builder) clone() *Builder {
	var cb Builder
	cb = *b
	return &cb
}

// Where set conditions of where in SELECT
// Where("user = ?","tom")
// Where("a = ? OR b = ?",1,2)
// Where("foo = $1","bar")
func (b *Builder) Where(query string, args ...interface{}) *Builder {
	b.filters = append(b.filters, filter{
		expression: query,
		args:       args,
	})

	return b
}

// Order sets columns of ORDER BY in SELECT.
// Order("name, age DESC")
func (b *Builder) Order(cols string) *Builder {
	b.order = cols
	return b
}

// Offset sets the offset in SELECT.
func (b *Builder) Offset(x int) *Builder {
	b.offset = x
	return b
}

// Limit sets the limit in SELECT.
func (b *Builder) Limit(x int) *Builder {
	b.limit = x
	return b
}

// Build returns compiled SELECT string and args.
func (b *Builder) Build(query string, args ...interface{}) (string, []interface{}) {
	var sb strings.Builder

	sb.WriteString(query)
	sb.WriteByte('\n')

	// append where conditions
	for idx, filter := range b.filters {
		if idx == 0 {
			sb.WriteString("WHERE ")
		} else {
			sb.WriteString("AND ")
		}

		sb.WriteByte('(')
		sb.WriteString(filter.expression)
		sb.WriteByte(')')
		sb.WriteByte('\n')

		args = append(args, filter.args...)
	}

	if b.order != "" {
		sb.WriteString("ORDER BY ")
		sb.WriteString(b.order)
		sb.WriteByte('\n')
	}

	if b.limit > 0 {
		sb.WriteString("LIMIT ")
		sb.WriteString(strconv.Itoa(b.limit))
		sb.WriteByte('\n')
	}

	if b.offset > 0 {
		sb.WriteString("OFFSET ")
		sb.WriteString(strconv.Itoa(b.offset))
		sb.WriteByte('\n')
	}

	return sb.String(), args
}
