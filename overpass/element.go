package overpass

import (
	"fmt"
	"strings"
)

func NewElementStatement(elementType ElementType, tagFilters []TagFilter, elementFilter ...*ElementFilter) *ElementStatement {
	s := new(ElementStatement)
	s.ElementType = elementType
	s.TagFilters = tagFilters

	if len(elementFilter) > 0 {
		s.ElementFilter = elementFilter[0]
	}

	s.compiled = s.compile()
	return s
}

func (s *ElementStatement) compile() string {
	var c strings.Builder
	var _c string

	c.WriteString(string(s.ElementType))

	for _, f := range s.TagFilters {
		if f.TagFilterType == Exists || f.TagFilterType == NotExists {
			_c = fmt.Sprintf("[%s\"%s\"]", f.TagFilterType, f.Key)
		} else {
			_c = fmt.Sprintf("[\"%s\"%s\"%s\"]", f.Key, f.TagFilterType, f.Value)
		}
		c.WriteString(_c)
	}

	if s.ElementFilter != nil {
		c.WriteString(fmt.Sprintf("(%s)", s.ElementFilter.FilterStr))
	}

	c.WriteString(";")

	return c.String()
}

func (s *ElementStatement) GetCompiled() string {
	return s.compiled
}

func (s *ElementStatement) GetSetName() string {
	return ""
}
