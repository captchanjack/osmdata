package overpass

import "fmt"

func NewOutStatement(verbosityType VebosityType) *OutStatement {
	s := new(OutStatement)
	s.VebosityType = verbosityType
	s.compiled = s.compile()
	return s
}

func (s *OutStatement) compile() string {
	return fmt.Sprintf("out %s;", s.VebosityType)
}

func (s *OutStatement) GetCompiled() string {
	return s.compiled
}

func (s *OutStatement) GetSetName() string {
	return ""
}
