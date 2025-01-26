package filters

import (
	"fmt"
	"gh-reponark/repo"
)

type IntFilter struct {
	Name string
	From int
	To   int
}

func NewIntFilter(name string, from, to int) IntFilter {
	return IntFilter{Name: name, From: from, To: to}
}

func (f IntFilter) GetName() string {
	return f.Name
}

func (f IntFilter) Matches(property repo.RepoProperty) bool {
	if property.Type != "int" {
		return false
	}

	value := property.Value.(int)

	return value >= f.From && value <= f.To
}

func (f IntFilter) String() string {
	return fmt.Sprintf("%s between %d and %d", f.Name, f.From, f.To)
}
