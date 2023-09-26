package consts

type Focus int

const (
	FocusList Focus = iota
	FocusTabs
	FocusFilter
)

func (f Focus) Next() Focus {
	switch f {
	case FocusList:
		return FocusTabs
	default:
		return FocusFilter
	}
}

func (f Focus) Prev() Focus {
	switch f {
	case FocusFilter:
		return FocusTabs
	default:
		return FocusList
	}
}
