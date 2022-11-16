package overpass

import (
	"fmt"
	"reflect"
	"strings"
)

func NewSetting(key SettingType, value interface{}, options ...string) *Setting {
	if _, ok := value.(OutType); key == Out && !ok {
		panic(fmt.Errorf("key '%s' must be paired with OutType (e.g. XML, JSON, CSV), not '%s' (%s)", key, value, reflect.TypeOf(value)))
	}
	s := new(Setting)
	s.Key = key
	s.Value = fmt.Sprintf("%v", value)
	s.Options = ""
	if len(options) > 0 {
		s.Options = options[0]
	}

	if len(s.Options) == 0 && value == CSV {
		panic("CSV output format requires header options (i.e. which tags/headers to keep)")
	}

	return s
}

func NewSettingsStatement(settings ...Setting) *SettingsStatement {
	s := new(SettingsStatement)
	s.Settings = settings
	s.compiled = s.compile()
	return s
}

func (s *SettingsStatement) compile() string {
	var c strings.Builder

	for _, st := range s.Settings {
		c.WriteString(fmt.Sprintf("[%s:%s%s]", st.Key, st.Value, st.Options))
	}

	c.WriteString(";")

	return c.String()
}

func (s *SettingsStatement) GetCompiled() string {
	return s.compiled
}

func (s *SettingsStatement) GetSetName() string {
	return "_"
}
