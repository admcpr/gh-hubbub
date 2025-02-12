package filters

import (
	"fmt"
	"gh-reponark/repo"
	"gh-reponark/shared"
)

type BoolFilter struct {
	Name  string
	Value bool
}

func NewBoolFilter(name string, value bool) BoolFilter {
	return BoolFilter{Name: name, Value: value}
}

func (f BoolFilter) GetName() string {
	return f.Name
}

func (f BoolFilter) Matches(property repo.RepoProperty) bool {
	if property.Type != "bool" {
		return false
	}

	return property.Value.(bool) == f.Value
}

func (f BoolFilter) String() string {
	return fmt.Sprintf("%s = %s", f.Name, shared.YesNo(f.Value))
}
