package doc

import (
	"fmt"
	"html"
	"strings"
)

type Paragraph struct {
	data  []string
	attr  map[string]string
	space int
}

func (p *Paragraph) SetAttr(key string, val string) *Paragraph {
	p.attr[key] = val
	return p
}

func (p *Paragraph) Add(str string) *Paragraph {
	p.data = append(p.data, html.EscapeString(str))
	return p
}

func (p *Paragraph) Template() (str string, err error) {
	attr := ""
	for key, val := range p.attr {
		attr += fmt.Sprintf(" %s=\"%s\"", key, val)
	}
	return fmt.Sprintf("\r\n%s<p%s>%s</p>\r\n", delimiter(p.space), attr, strings.Join(p.data, "\r\n")), nil
}

func (p *Paragraph) Tag() string {
	str, _ := p.Template()
	return str
}

func NewParagraph(space int) *Paragraph {
	return &Paragraph{space: space, data: []string{}, attr: make(map[string]string)}
}
