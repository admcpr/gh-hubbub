package structs

type Filter interface {
	GetName() string
	Matches(property RepoProperty) bool
	String() string
}
