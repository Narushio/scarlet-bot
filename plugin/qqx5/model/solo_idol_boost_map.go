package model

type SoloIdolBoostMap struct {
	Letter          string    `json:"letter"`
	Title           string    `json:"title"`
	Rating          string    `json:"rating"`
	MapRating       string    `json:"map_rating"`
	MapDifficulty   string    `json:"map_difficulty"`
	BoostDifficulty string    `json:"boost_difficulty"`
	Recommend       string    `json:"recommend"`
	BoostInfo       BoostInfo `json:"boost_info"`
	Memo            string    `json:"memo"`
	TheoryScore     string    `json:"theory_score"`
	HalfCombo       string    `json:"half_combo"`
	Comment         string    `json:"comment"`
}

var SoloIdolBoostMapList []*SoloIdolBoostMap
