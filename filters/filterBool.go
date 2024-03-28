package filters

import (
	"fmt"
	"gh-hubbub/structs"
	"reflect"
)

type FilterBool struct {
	Tab   string
	Name  string
	Value bool
}

func NewFilterBool(tab, name string, value bool) FilterBool {
	return FilterBool{Tab: tab, Name: name, Value: value}
}

func (f FilterBool) GetTab() string {
	return f.Tab
}

func (f FilterBool) GetName() string {
	return f.Name
}

func (f FilterBool) Matches(setting structs.Setting) bool {
	if setting.Type != reflect.TypeOf(f.Value) {
		return false
	}

	return setting.Value.(bool) == f.Value
}

func (f FilterBool) String() string {
	return fmt.Sprintf("%s = %s", f.Name, structs.YesNo(f.Value))
}
