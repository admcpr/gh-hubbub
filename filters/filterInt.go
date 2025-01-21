package filters

import (
	"fmt"
	"gh-hubbub/structs"
)

type FilterInt struct {
	Name string
	From int
	To   int
}

func NewFilterInt(name string, from, to int) FilterInt {
	return FilterInt{Name: name, From: from, To: to}
}

func (f FilterInt) GetName() string {
	return f.Name
}

func (f FilterInt) Matches(property structs.RepoProperty) bool {
	if property.Type != "int" {
		return false
	}

	value := property.Value.(int)

	return value >= f.From && value <= f.To
}

func (f FilterInt) String() string {
	return fmt.Sprintf("%s between %d and %d", f.Name, f.From, f.To)
}
