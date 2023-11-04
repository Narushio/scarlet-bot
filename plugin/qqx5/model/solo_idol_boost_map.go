package model

import (
	"fmt"

	"github.com/xuri/excelize/v2"

	"github.com/Narushio/scarlet-bot/helper/file"
	"github.com/Narushio/scarlet-bot/helper/image"
)

const SoloIdolBoostMapJsonPath = "resources/qqx5/solo_idol_boost_maps.json"

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

func (m *SoloIdolBoostMap) UnmarshalFromJson() error {
	list, err := file.ReadJson[[]*SoloIdolBoostMap](SoloIdolBoostMapJsonPath)
	if err != nil {
		return err
	}
	SoloIdolBoostMapList = list
	return nil
}

func (m *SoloIdolBoostMap) UnmarshalFromExcel(f *excelize.File) error {
	rows, err := f.GetRows("星动")
	if err != nil {
		return err
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		boostMap := &SoloIdolBoostMap{}
		boostPicColNameList := []string{"I", "J"}
		for _, n := range boostPicColNameList {
			p, err := f.GetPictures("星动", fmt.Sprintf("%s%d", n, i+1))
			if err != nil {
				return err
			}

			if len(p) != 0 {
				switch n {
				case "J":
					boostMap.BoostInfo.End = image.ToBase64Src(p[0].File, image.Jpeg)
				default:
					boostMap.BoostInfo.Start = image.ToBase64Src(p[0].File, image.Jpeg)
				}
			}
		}

		if row[0] == "" {
			boostMap.Letter = SoloIdolBoostMapList[len(SoloIdolBoostMapList)-1].Letter
		} else {
			boostMap.Letter = row[0]
		}
		boostMap.Title = row[1]
		boostMap.Rating = row[2]
		boostMap.MapRating = row[3]
		boostMap.MapDifficulty = row[4]
		boostMap.BoostDifficulty = row[5]
		boostMap.Recommend = row[6]
		boostInfoColumnIndex := 8
		if len(row) > boostInfoColumnIndex {
			boostMap.BoostInfo.Description = row[7]
			boostMap.Memo = row[10]
			boostMap.TheoryScore = row[11]
			boostMap.HalfCombo = row[12]
			boostMap.Comment = row[13]
		}
		SoloIdolBoostMapList = append(SoloIdolBoostMapList, boostMap)
	}

	err = file.WriteJson(SoloIdolBoostMapJsonPath, SoloIdolBoostMapList)
	if err != nil {
		return err
	}

	return nil
}
