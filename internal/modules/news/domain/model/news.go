package model

import (
	"fmt"
	"gopkg.in/reform.v1"
)

// News corresponds to the News table
type News struct {
	ID      int64  `reform:"Id,pk"`
	Title   string `reform:"Title"`
	Content string `reform:"Content"`
}

var newsTable = &tableImpl{
	name:    "News",
	columns: []string{"Id", "Title", "Content"},
}

func (n *News) String() string {
	return fmt.Sprintf("News (Id: %d, Title: %s)", n.ID, n.Title)
}

func (n *News) Table() reform.Table {
	return newsTable
}

func (n *News) Values() []interface{} {
	return []interface{}{n.ID, n.Title, n.Content}
}

func (n *News) Pointers() []interface{} {
	return []interface{}{&n.ID, &n.Title, &n.Content}
}

func (n *News) PKValue() interface{} {
	return n.ID
}

func (n *News) PKPointer() interface{} {
	return &n.ID
}

func (n *News) View() reform.View {
	return newsTable
}

func (n *News) HasPK() bool {
	return n.ID > 0
}

func (n *News) SetPK(pk interface{}) {
	n.ID = pk.(int64)
}

type tableImpl struct {
	name    string
	columns []string
}

func (t *tableImpl) Name() string {
	return t.name
}

func (t *tableImpl) Columns() []string {
	return t.columns
}

func (t *tableImpl) NewStruct() reform.Struct {
	return &News{}
}

func (t *tableImpl) NewRecord() reform.Record {
	return &News{}
}

func (t *tableImpl) PKColumnIndex() uint {
	return 0
}

func (t *tableImpl) Schema() string {
	return ""
}
