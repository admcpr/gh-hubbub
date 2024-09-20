package filters

type FilterAction int

const (
	FilterAdd FilterAction = iota
	FilterConfirm
	FilterDelete
	FilterCancel
)

type FilterMsg struct {
	Action FilterAction
	Filter Filter
}

func NewAddFilterMsg(filter Filter) FilterMsg {
	return FilterMsg{Action: FilterAdd, Filter: filter}
}

func NewConfirmFilterMsg(filter Filter) FilterMsg {
	return FilterMsg{Action: FilterConfirm, Filter: filter}
}

func NewDeleteFilterMsg(filter Filter) FilterMsg {
	return FilterMsg{Action: FilterDelete, Filter: filter}
}

func NewCancelFilterMsg(filter Filter) FilterMsg {
	return FilterMsg{Action: FilterCancel, Filter: filter}
}
