package filters

import "gh-hubbub/structs"

type Filter interface {
	GetName() string
	Matches(property structs.RepoProperty) bool
	String() string
}
