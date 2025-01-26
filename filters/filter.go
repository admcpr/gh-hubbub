package filters

import (
	"gh-reponark/repo"
)

type Filter interface {
	GetName() string
	Matches(property repo.RepoProperty) bool
	String() string
}
