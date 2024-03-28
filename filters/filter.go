package filters

import "gh-hubbub/structs"

type Filter interface {
	GetTab() string
	GetName() string
	Matches(setting structs.Setting) bool
	String() string
}
