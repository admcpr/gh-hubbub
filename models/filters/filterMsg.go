package filters

import (
	"gh-hubbub/consts"
	"gh-hubbub/structs"
)

type FilterMsg struct {
	Action consts.FilterAction
	Filter structs.Filter
}

func NewAddFilterMsg(filter structs.Filter) FilterMsg {
	return FilterMsg{Action: consts.FilterAdd, Filter: filter}
}

func NewConfirmFilterMsg(filter structs.Filter) FilterMsg {
	return FilterMsg{Action: consts.FilterConfirm, Filter: filter}
}

func NewDeleteFilterMsg(filter structs.Filter) FilterMsg {
	return FilterMsg{Action: consts.FilterDelete, Filter: filter}
}

func NewCancelFilterMsg(filter structs.Filter) FilterMsg {
	return FilterMsg{Action: consts.FilterCancel, Filter: filter}
}
