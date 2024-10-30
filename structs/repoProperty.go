package structs

import (
	"gh-hubbub/queries"
	"reflect"
	"regexp"
	"strings"
)

type RepoProperties struct {
	Properties     []RepoProperty
	PropertyGroups map[string][]RepoProperty
}

func NewRepoProperties(r queries.Repository) RepoProperties {
	propertyGroups := make(map[string][]RepoProperty)
	properties := ToProperties(r)
	for _, p := range properties {
		propertyGroups[p.Group] = append(propertyGroups[p.Group], p)
	}
	return RepoProperties{Properties: properties, PropertyGroups: propertyGroups}
}

type RepoProperty struct {
	Name        string
	Group       string
	Value       interface{}
	Type        string
	Description string
}

func NewRepoProperty(name string, group string, value interface{}, typeStr string, description string) RepoProperty {
	return RepoProperty{Name: name, Group: group, Value: value, Type: typeStr, Description: description}
}

func ToProperties(r queries.Repository) []RepoProperty {
	var properties []RepoProperty
	t := reflect.TypeOf(r)
	v := reflect.ValueOf(r)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		name := splitCamelCase(field.Name)
		desc := field.Tag.Get("desc")
		group := field.Tag.Get("group")
		typeStr := field.Type.String()

		properties = append(properties, NewRepoProperty(name, group, v.Field(i).Interface(), typeStr, desc))
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
