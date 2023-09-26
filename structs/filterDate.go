package structs

import (
	"fmt"
	"reflect"
	"time"
)

type FilterDate struct {
	Tab  string
	Name string
	From time.Time
	To   time.Time
}

func NewFilterDate(tab, name string, from, to time.Time) FilterDate {
	return FilterDate{Tab: tab, Name: name, From: from, To: to}
}

func (f FilterDate) GetTab() string {
	return f.Tab
}

func (f FilterDate) GetName() string {
	return f.Name
}

func (f FilterDate) Matches(setting Setting) bool {
	if setting.Type != reflect.TypeOf(f.From) {
		return false
	}

	date := setting.Value.(time.Time)

	return date.After(f.From) && date.Before(f.To)
}

func (f FilterDate) String() string {
	return fmt.Sprintf("%s > %s between %s and %s", f.Tab, f.Name, f.From.Format("2006-01-02"), f.To.Format("2006-01-02"))
}
