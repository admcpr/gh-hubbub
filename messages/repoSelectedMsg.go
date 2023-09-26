package messages

import (
	"gh-hubbub/structs"
)

type RepoSelectMsg struct {
	Repository structs.Repository
	Width      int
	Height     int
}
