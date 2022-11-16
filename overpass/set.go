package overpass

import "fmt"

func NewSetStatement(setName string) *SetStatement {
	s := new(SetStatement)
	s.SetName = setName
	s.compiled = s.compile()
	return s
}

func (s *SetStatement) compile() string {
	return fmt.Sprintf(".%s;", s.SetName)
}

func (s *SetStatement) GetCompiled() string {
	return s.compiled
}

func (s *SetStatement) GetSetName() string {
	return s.SetName
}
