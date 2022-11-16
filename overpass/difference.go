package overpass

import (
	"fmt"
	"reflect"
	"strings"
)

func NewDifferenceStatement(resultSetName string, firstStatement Statement, secondStatement Statement) *DifferenceStatement {
	s := new(DifferenceStatement)
	if resultSetName != "" {
		s.SetName = resultSetName
	}
	s.FirstStatement = firstStatement
	s.SecondStatement = secondStatement
	s.compiled = s.compile()
	return s
}

func (s *DifferenceStatement) compile() string {
	_, isUnionFirst := s.FirstStatement.(*UnionStatement)
	_, isDiffFirst := s.FirstStatement.(*DifferenceStatement)

	if isUnionFirst || isDiffFirst {
		panic(fmt.Errorf("%s cannot wrap %s", reflect.TypeOf(s), reflect.TypeOf(s.FirstStatement)))
	}

	_, isUnionSecond := s.SecondStatement.(*UnionStatement)
	_, isDiffSecond := s.SecondStatement.(*DifferenceStatement)

	if isUnionSecond || isDiffSecond {
		panic(fmt.Errorf("%s cannot wrap %s", reflect.TypeOf(s), reflect.TypeOf(s.FirstStatement)))
	}

	var c strings.Builder

	c.WriteString(
		fmt.Sprintf(
			"(%s - %s)->.%s;",
			s.FirstStatement.GetCompiled(),
			s.SecondStatement.GetCompiled(),
			s.SetName,
		),
	)

	return c.String()
}

func (s *DifferenceStatement) GetCompiled() string {
	return s.compiled
}

func (s *DifferenceStatement) GetSetName() string {
	return s.SetName
}
