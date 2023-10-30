package model

type Pattern int

const (
	Idol Pattern = iota
	Pinball
	Bubble
	Crescent
)

func (p *Pattern) String() string {
	switch *p {
	case Idol:
		return "星动"
	case Pinball:
		return "弹珠"
	case Bubble:
		return "泡泡"
	case Crescent:
		return "弦月"
	default:
		return "未知"
	}
}
