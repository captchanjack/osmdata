package overpass

import (
	"fmt"
	"reflect"
	"strings"
)

func NewUnionStatement(resultSetName string, statements ...Statement) *UnionStatement {
	s := new(UnionStatement)
	if resultSetName != "" {
		s.SetName = resultSetName
	}
	s.Statements = statements
	s.compiled = s.compile()
	return s
}

func (s *UnionStatement) compile() string {
	var c strings.Builder

	c.WriteString("(")

	for _, _s := range s.Statements {
		_, isUnion := _s.(*UnionStatement)
		_, isDiff := _s.(*DifferenceStatement)

		if isUnion || isDiff {
			panic(fmt.Errorf("%s cannot wrap %s", reflect.TypeOf(s), reflect.TypeOf(_s)))
		}

		c.WriteString(_s.GetCompiled())
	}

	c.WriteString(fmt.Sprintf(")->.%s;", s.SetName))

	return c.String()
}

func (s *UnionStatement) GetCompiled() string {
	return s.compiled
}

func (s *UnionStatement) GetSetName() string {
	return s.SetName
}

// Append statements and recompiles the query string
func (s *UnionStatement) Append(statements ...Statement) {
	s.Statements = append(s.Statements, statements...)
	s.compiled = s.compile()
}
