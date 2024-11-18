package structs

import (
	"fmt"
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

func (f FilterBool) Matches(property RepoProperty) bool {
	if property.Type != "bool" {
		return false
	}

	return property.Value.(bool) == f.Value
}

func (f FilterBool) String() string {
	return fmt.Sprintf("%s = %s", f.Name, YesNo(f.Value))
}
