package messages

import (
	"gh-hubbub/structs"
)

type RepoSelectMsg struct {
	Repository structs.RepositorySettings
	Width      int
	Height     int
}
