package structs

import (
	"fmt"
	"gh-hubbub/queries"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"
)

type RepoProperties struct {
	Name           string
	Url            string
	Properties     []RepoProperty
	PropertyGroups map[string][]RepoProperty
	Keys           []string
}

func NewRepoProperties(r queries.Repository) RepoProperties {
	propertyGroups := make(map[string][]RepoProperty)
	properties := ToProperties(r)
	for _, p := range properties {
		propertyGroups[p.Group] = append(propertyGroups[p.Group], p)
	}
	keys := make([]string, 0, len(propertyGroups))
	for k := range propertyGroups {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return RepoProperties{
		Name:           r.Name,
		Url:            r.Url,
		Properties:     properties,
		PropertyGroups: propertyGroups,
		Keys:           keys,
	}
}

type RepoProperty struct {
	Name        string
	Group       string
	Value       interface{}
	Type        string
	Description string
}

func (s RepoProperty) String() string {
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

func NewRepoProperty(name string, group string, value interface{}, typeStr string, description string) RepoProperty {
	return RepoProperty{Name: name, Group: group, Value: value, Type: typeStr, Description: description}
}

// func ToProperties(r queries.Repository) []RepoProperty {
// 	var properties []RepoProperty
// 	t := reflect.TypeOf(r)
// 	v := reflect.ValueOf(r)

// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)

// 		name := splitCamelCase(field.Name)
// 		desc := field.Tag.Get("desc")
// 		group := field.Tag.Get("group")
// 		typeStr := field.Type.String()

// 		properties = append(properties, NewRepoProperty(name, group, v.Field(i).Interface(), typeStr, desc))
// 	}

// 	return properties
// }

func ToProperties(r queries.Repository) []RepoProperty {
	var properties []RepoProperty
	t := reflect.TypeOf(r)
	v := reflect.ValueOf(r)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		properties = append(properties, processField(field, value)...)
	}

	return properties
}

func processField(field reflect.StructField, value reflect.Value) []RepoProperty {
	var properties []RepoProperty

	name := field.Tag.Get("name")
	if name == "" {
		name = splitCamelCase(field.Name)
	}
	desc := field.Tag.Get("desc")
	group := field.Tag.Get("group")
	typeStr := field.Type.String()

	if typeStr == "time.Time" || typeStr == "int" || typeStr == "string" || typeStr == "bool" {
		properties = append(properties, NewRepoProperty(name, group, value.Interface(), typeStr, desc))
	} else {
		switch field.Type.Kind() {
		case reflect.Struct:
			// Recursively process nested struct
			for i := 0; i < field.Type.NumField(); i++ {
				nestedField := field.Type.Field(i)
				nestedValue := value.Field(i)
				properties = append(properties, processField(nestedField, nestedValue)...)
			}
		case reflect.Ptr:
			// Handle pointer to struct
			if !value.IsNil() && value.Elem().Kind() == reflect.Struct {
				properties = append(properties, processField(field, value.Elem())...)
			} else {
				properties = append(properties, NewRepoProperty(name, group, value.Interface(), typeStr, desc))
			}
		}
	}

	return properties
}

func splitCamelCase(s string) string {
	// Add space before capital letters
	words := regexp.MustCompile(`([a-z0-9])([A-Z])`).ReplaceAllString(s, "$1 $2")
	// Handle consecutive capital letters (like "ID" in "UserID")
	words = regexp.MustCompile(`([A-Z])([A-Z][a-z])`).ReplaceAllString(words, "$1 $2")
	// Capitalize first letter and trim spaces
	return strings.TrimSpace(words)
}
