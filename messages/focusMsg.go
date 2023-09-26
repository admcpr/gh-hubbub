package messages

import "gh-hubbub/consts"

type FocusMsg struct {
	Focus consts.Focus
}

func NewFocusMsg(focus consts.Focus) FocusMsg {
	return FocusMsg{Focus: focus}
}
