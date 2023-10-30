package model

type DuoIdolBoostMap struct {
	Letter string    `json:"letter"`
	Title  string    `json:"title"`
	PhaseA BoostInfo `json:"phase_a"`
	PhaseB BoostInfo `json:"phase_b"`
	Memo   string    `json:"memo"`
}

var DuoIdolBoostMapList []*DuoIdolBoostMap
