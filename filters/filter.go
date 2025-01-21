package filters

import (
	"gh-hubbub/repos"
)

type Filter interface {
	GetName() string
	Matches(property repos.RepoProperty) bool
	String() string
}
