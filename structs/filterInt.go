package structs

import (
	"fmt"
	"reflect"
)

type FilterInt struct {
	Tab  string
	Name string
	From int
	To   int
}

func NewFilterInt(tab, name string, from, to int) FilterInt {
	return FilterInt{Tab: tab, Name: name, From: from, To: to}
}

func (f FilterInt) GetTab() string {
	return f.Tab
}

func (f FilterInt) GetName() string {
	return f.Name
}

func (f FilterInt) Matches(setting Setting) bool {
	if setting.Type != reflect.TypeOf(f.From) {
		return false
	}

	value := setting.Value.(int)

	return value >= f.From && value <= f.To
}

func (f FilterInt) String() string {
	return fmt.Sprintf("%s > %s between %d and %d", f.Tab, f.Name, f.From, f.To)
}
