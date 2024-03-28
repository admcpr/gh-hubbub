package structs

import (
	"fmt"
	"reflect"
	"time"
)

type Setting struct {
	Name  string
	Value interface{}
	Type  reflect.Type
}

func (s Setting) String() string {
	switch value := s.Value.(type) {
	case bool:
		return YesNo(value)
	case string:
		return value
	case time.Time:
		return value.Format("2006/01/02")
	case int:
		return fmt.Sprint(value)
	default:
		return "Unknown type"
	}
}

func NewSetting(name string, value interface{}) Setting {
	return Setting{Name: name, Value: value, Type: reflect.TypeOf(value)}
}
