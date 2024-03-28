package filters

import (
	"gh-hubbub/consts"
)

type FilterMsg struct {
	Action consts.FilterAction
	Filter Filter
}

func NewAddFilterMsg(filter Filter) FilterMsg {
	return FilterMsg{Action: consts.FilterAdd, Filter: filter}
}

func NewConfirmFilterMsg(filter Filter) FilterMsg {
	return FilterMsg{Action: consts.FilterConfirm, Filter: filter}
}

func NewDeleteFilterMsg(filter Filter) FilterMsg {
	return FilterMsg{Action: consts.FilterDelete, Filter: filter}
}

func NewCancelFilterMsg(filter Filter) FilterMsg {
	return FilterMsg{Action: consts.FilterCancel, Filter: filter}
}
