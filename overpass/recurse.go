package overpass

func NewRecurseStatement(recurseType RecurseType) *RecurseStatement {
	s := new(RecurseStatement)
	s.RecurseType = recurseType
	s.compiled = s.compile()
	return s
}

func (s *RecurseStatement) compile() string {
	return string(s.RecurseType) + ";"
}

func (s *RecurseStatement) GetCompiled() string {
	return s.compiled
}

func (s *RecurseStatement) GetSetName() string {
	return ""
}
