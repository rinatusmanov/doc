package doc

import (
	"fmt"
)

func NewHr(space int) *Hr {
	return &Hr{space: space, attr: make(map[string]string)}
}

type Hr struct {
	attr  map[string]string
	space int
}

func (hr *Hr) SetAttr(key string, val string) *Hr {
	hr.attr[key] = val
	return hr
}

func (hr *Hr) Template() (str string, err error) {
	attr := ""
	for key, val := range hr.attr {
		attr += fmt.Sprintf(" %s=\"%s\"", key, val)
	}
	return fmt.Sprintf("\r\n%s<hr%s/>\r\n", delimiter(hr.space), attr), nil
}

func (hr *Hr) Tag() string {
	str, _ := hr.Template()
	return str
}
