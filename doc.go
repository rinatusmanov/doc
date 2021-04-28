package doc

import (
	_ "embed"
	"os"
	"text/template"
)

type Doc struct {
	attr map[string]string
	head []Tag
	body []Tag
}

func (d *Doc) Head() []Tag {
	return d.head
}

func (d *Doc) Body() []Tag {
	return d.body
}

func NewDoc() *Doc {
	return &Doc{attr: make(map[string]string), head: []Tag{}, body: []Tag{}}
}

//go:embed tml/doc.gohtml
var docTml string
var docTmlParsed = template.Must(template.New("doc").Parse(docTml))

func (d *Doc) String() (str string, err error) {
	writer := &tWriter{}
	err = docTmlParsed.Execute(writer, d)
	return writer.String(), err
}

func (d *Doc) SaveFile(fileName string) error {
	str, err := d.String()
	if err != nil {
		return err
	}
	_ = os.Remove(fileName)
	file, errCreate := os.Create(fileName)
	if errCreate != nil {
		return errCreate
	}
	defer file.Close()
	_, errWrite := file.WriteString(str)
	if errWrite != nil {
		return errWrite
	}
	return nil
}

func (d *Doc) CreateTable() *Table {
	table := NewTable(2)
	d.body = append(d.body, table)
	return table
}

func (d *Doc) CreateStyle() *Style {
	data := NewStyle(2)
	d.head = append(d.head, data)
	return data
}
func (d *Doc) CreateParagraph() *Paragraph {
	data := NewParagraph(2)
	d.body = append(d.body, data)
	return data
}

func (d *Doc) CreateHr() *Hr {
	data := NewHr(2)
	d.body = append(d.body, data)
	return data
}
