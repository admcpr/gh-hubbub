package structs

type Filter interface {
	GetTab() string
	GetName() string
	Matches(setting Setting) bool
	String() string
}
