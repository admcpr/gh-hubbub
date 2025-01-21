package filters

import (
	"fmt"
	"gh-hubbub/structs"
)

type FilterBool struct {
	Name  string
	Value bool
}

func NewFilterBool(name string, value bool) FilterBool {
	return FilterBool{Name: name, Value: value}
}

func (f FilterBool) GetName() string {
	return f.Name
}

func (f FilterBool) Matches(property structs.RepoProperty) bool {
	if property.Type != "bool" {
		return false
	}

	return property.Value.(bool) == f.Value
}

func (f FilterBool) String() string {
	return fmt.Sprintf("%s = %s", f.Name, structs.YesNo(f.Value))
}
