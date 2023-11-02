package model

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"

	"github.com/Narushio/scarlet-bot/helper/file"
	"github.com/Narushio/scarlet-bot/helper/image"
)

const DuoIdolBoostMapJsonPath = "resource/qqx5/duo_idol_boost_maps.json"

type DuoIdolBoostMap struct {
	Letter string    `json:"letter"`
	Title  string    `json:"title"`
	PhaseA BoostInfo `json:"phase_a"`
	PhaseB BoostInfo `json:"phase_b"`
	Memo   string    `json:"memo"`
}

var DuoIdolBoostMapList []*DuoIdolBoostMap

func (m *DuoIdolBoostMap) UnmarshalFromJson() error {
	list, err := file.ReadJson[[]*DuoIdolBoostMap](DuoIdolBoostMapJsonPath)
	if err != nil {
		return err
	}
	DuoIdolBoostMapList = list
	return nil
}

func (m *DuoIdolBoostMap) UnmarshalFromExcel(f *excelize.File) error {
	rows, err := f.GetRows("星动双排")
	if err != nil {
		return err
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		boostMap := &DuoIdolBoostMap{}
		colNames := []string{"D", "E", "G", "H"}
		for _, c := range colNames {
			p, err := f.GetPictures("星动双排", fmt.Sprintf("%s%d", c, i+1))
			if err != nil {
				log.Fatal(err)
			}

			if len(p) != 0 {
				switch c {
				case "E":
					boostMap.PhaseA.End = image.ToBase64Src(p[0].File, image.Jpeg)
				case "G":
					boostMap.PhaseB.Start = image.ToBase64Src(p[0].File, image.Jpeg)
				case "H":
					boostMap.PhaseB.End = image.ToBase64Src(p[0].File, image.Jpeg)
				default:
					boostMap.PhaseA.Start = image.ToBase64Src(p[0].File, image.Jpeg)
				}
			}
		}

		if row[0] == "" {
			boostMap.Letter = DuoIdolBoostMapList[len(DuoIdolBoostMapList)-1].Letter
		} else {
			boostMap.Letter = row[0]
		}
		boostMap.Title = row[1]
		boostMap.PhaseA.Description = row[2]
		boostMap.PhaseB.Description = row[5]
		boostEndColumnIndex := 8
		if len(row) > boostEndColumnIndex {
			boostMap.Memo = row[8]
		}
		DuoIdolBoostMapList = append(DuoIdolBoostMapList, boostMap)
	}

	err = file.WriteJson(DuoIdolBoostMapJsonPath, DuoIdolBoostMapList)
	if err != nil {
		return err
	}

	return nil
}
