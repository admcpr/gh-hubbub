package filters

import (
	"fmt"
	"gh-hubbub/structs"
	"time"
)

type FilterDate struct {
	Name string
	From time.Time
	To   time.Time
}

func NewFilterDate(name string, from, to time.Time) FilterDate {
	return FilterDate{Name: name, From: from, To: to}
}

func (f FilterDate) GetName() string {
	return f.Name
}

func (f FilterDate) Matches(property structs.RepoProperty) bool {
	if property.Type != "time.Time" {
		return false
	}

	date := property.Value.(time.Time)

	return (date.After(f.From) || date.Equal(f.From)) && (date.Before(f.To) || date.Equal(f.To))
}

func (f FilterDate) String() string {
	return fmt.Sprintf("%s between %s and %s", f.Name, f.From.Format("2006-01-02"), f.To.Format("2006-01-02"))
}
