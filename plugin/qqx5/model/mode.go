package model

type Mode int

const (
	Solo Mode = iota
	Duo
)

func (m *Mode) String() string {
	switch *m {
	case Solo:
		return "单排"
	case Duo:
		return "双排"
	default:
		return "未知"
	}
}
