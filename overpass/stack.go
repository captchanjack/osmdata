package overpass

import (
	"os"
	"strings"
)

func NewStackStatement(statements ...Statement) *StackStatement {
	s := new(StackStatement)
	s.Statements = statements
	s.compiled = s.compile()
	return s
}

func (s *StackStatement) compile() string {
	var c strings.Builder

	for _, st := range s.Statements {
		c.WriteString(st.GetCompiled())
	}

	return c.String()
}

func (s *StackStatement) GetCompiled() string {
	return s.compiled
}

func (s *StackStatement) GetSetName() string {
	return ""
}

// Append statements and recompiles the query string
func (s *StackStatement) Append(statements ...Statement) {
	s.Statements = append(s.Statements, statements...)
	s.compiled = s.compile()
}

// Executes the query string
// Specify maxAttempts to retry when rate limited, defaults to 1
func (s *StackStatement) Execute(maxAttempts ...int) (string, error) {
	return QueryOverpass(s.GetCompiled(), maxAttempts...)
}

// Executes the query string and exports the data to file on disk, e.g. ./test.osm
// Specify maxAttempts to retry when rate limited, defaults to 1
func (s *StackStatement) ExecuteAndExport(filename string, maxAttempts ...int) (string, error) {
	resp, err := QueryOverpassBytes(s.GetCompiled(), maxAttempts...)

	if err != nil {
		return "", err
	}

	os.WriteFile(filename, resp, 0644)

	return string(resp), err
}
