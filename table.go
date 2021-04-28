package doc

import (
	"fmt"
	"html"
	"text/template"

	_ "embed"
)

type CellTable struct {
	tag     string
	content string
	attr    map[string]string
}

func NewTdTable(tag string) *CellTable {
	return &CellTable{tag: tag, content: "", attr: make(map[string]string)}
}

func (t *CellTable) Tag() (result string) {
	return t.tag
}

func (t *CellTable) Content() (result string) {
	return t.content
}

func (t *CellTable) SetContent(content string) *CellTable {
	t.content = html.EscapeString(content)
	return t
}

func (t *CellTable) Attr() (result string) {
	for key, val := range t.attr {
		result += fmt.Sprintf(" %s=\"%s\"", key, val)
	}
	return
}

func (t *CellTable) SetAttr(key string, val string) *CellTable {
	t.attr[key] = val
	return t
}

type TrTable struct {
	tdSlice []*CellTable
	attr    map[string]string
}

func NewTrTable() *TrTable {
	return &TrTable{tdSlice: []*CellTable{}, attr: make(map[string]string)}
}

func (t *TrTable) Attr() (result string) {
	for key, val := range t.attr {
		result += fmt.Sprintf(" %s=\"%s\"", key, val)
	}
	return
}

func (t *TrTable) SetAttr(key string, val string) *TrTable {
	t.attr[key] = val
	return t
}

func (t *TrTable) TdSlice() []*CellTable {
	return t.tdSlice
}

func (t *TrTable) Td() (td *CellTable) {
	td = NewTdTable("td")
	t.tdSlice = append(t.tdSlice, td)
	return td
}

func (t *TrTable) Th() (td *CellTable) {
	td = NewTdTable("th")
	t.tdSlice = append(t.tdSlice, td)
	return td
}

type TableBody struct {
	trSlice []*TrTable
	attr    map[string]string
}

func (t *TableBody) Attr() (result string) {
	for key, val := range t.attr {
		result += fmt.Sprintf(" %s=\"%s\"", key, val)
	}
	return
}

func (t *TableBody) SetAttr(key string, val string) *TableBody {
	t.attr[key] = val
	return t
}

func (t *TableBody) TrSlice() []*TrTable {
	return t.trSlice
}

func (t *TableBody) Tr() (tr *TrTable) {
	tr = NewTrTable()
	t.trSlice = append(t.trSlice, tr)
	return tr
}

func NewTableBody() *TableBody {
	return &TableBody{
		trSlice: []*TrTable{},
		attr:    make(map[string]string),
	}
}

func NewTable(space int) *Table {
	return &Table{
		uuid:    random(),
		space:   space,
		caption: "",
		attr:    make(map[string]string),
		tHead:   NewTableBody(),
		tFoot:   NewTableBody(),
		tBody:   NewTableBody(),
	}
}

type Table struct {
	uuid    string
	space   int
	caption string
	attr    map[string]string
	tHead   *TableBody
	tFoot   *TableBody
	tBody   *TableBody
}

func (t *Table) UUID() string {
	return t.uuid
}

func (t *Table) Head() *TableBody {
	return t.tHead
}

func (t *Table) Foot() *TableBody {
	return t.tFoot
}

func (t *Table) Body() *TableBody {
	return t.tBody
}

func (t *Table) Delimiter() string {
	return delimiter(t.space)
}

func (t *Table) Caption() string {
	if t.caption != "" {
		return fmt.Sprintf(
			"<caption>%s</caption>",
			t.caption)
	}
	return ""
}

func (t *Table) SetCaption(caption string) *Table {
	t.caption = html.EscapeString(caption)
	return t
}

func (t *Table) Attr() (result string) {
	for key, val := range t.attr {
		result += fmt.Sprintf(" %s=\"%s\"", key, val)
	}
	return
}

func (t *Table) SetAttr(key string, val string) *Table {
	t.attr[key] = val
	return t
}

//go:embed tml/table.gohtml
var tableTml string
var tableTmlParsed = template.Must(template.New("table").Parse(tableTml))

func (t *Table) Template() (str string, err error) {
	writer := &tWriter{}
	err = tableTmlParsed.Execute(writer, t)
	return writer.String(), err
}

func (t *Table) Tag() string {
	str, _ := t.Template()
	return str
}
