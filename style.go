package doc

import (
	"fmt"
	"strings"
)

type Style struct {
	data  []string
	space int
}

func (s *Style) Add(str string) *Style {
	data := strings.Split(str, "\n")
	str = ""
	for _, datum := range data {
		str += delimiter(s.space) + datum + "\n"
	}
	s.data = append(s.data, str)
	return s
}

func (s *Style) Template() (str string, err error) {
	return fmt.Sprintf("<style>\r\n%s%s</style>\n", strings.Join(s.data, "\r\n"), delimiter(s.space)), nil
}

func (s *Style) Tag() string {
	str, _ := s.Template()
	return str
}

func NewStyle(space int) *Style {
	return &Style{space: space, data: []string{}}
}
