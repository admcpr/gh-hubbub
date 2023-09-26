package consts

type FilterAction int

const (
	FilterAdd FilterAction = iota
	FilterConfirm
	FilterDelete
	FilterCancel
)
