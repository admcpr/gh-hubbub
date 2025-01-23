package filters

import (
	"gh-reponark/repos"
)

type Filter interface {
	GetName() string
	Matches(property repos.RepoProperty) bool
	String() string
}
