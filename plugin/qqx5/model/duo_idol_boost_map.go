package model

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"

	"github.com/Narushio/scarlet-bot/helper/file"
	"github.com/Narushio/scarlet-bot/helper/image"
)

type DuoIdolBoostMap struct {
	Letter string    `json:"letter"`
	Title  string    `json:"title"`
	PhaseA BoostInfo `json:"phase_a"`
	PhaseB BoostInfo `json:"phase_b"`
	Memo   string    `json:"memo"`
}

func (m *DuoIdolBoostMap) GetJsonPath() string {
	return "resources/qqx5/duo_idol_boost_maps.json"
}

func (m *DuoIdolBoostMap) GetExcelSheetName() string {
	return "星动双排"
}

func (m *DuoIdolBoostMap) GetTemplateUrl() string {
	return "http://localhost:8000/qqx5/duo-idol-boost-map.html"
}

func (m *DuoIdolBoostMap) GetTitle() string {
	return m.Title
}

var DuoIdolBoostMaps []*DuoIdolBoostMap

func (m *DuoIdolBoostMap) UnmarshalFromJson() error {
	maps, err := file.ReadJson[[]*DuoIdolBoostMap](m.GetJsonPath())
	if err != nil {
		return err
	}
	DuoIdolBoostMaps = maps
	return nil
}

func (m *DuoIdolBoostMap) UnmarshalFromExcel(f *excelize.File) error {
	rows, err := f.GetRows(m.GetExcelSheetName())
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
			p, err := f.GetPictures(m.GetExcelSheetName(), fmt.Sprintf("%s%d", c, i+1))
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
			boostMap.Letter = DuoIdolBoostMaps[len(DuoIdolBoostMaps)-1].Letter
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
		DuoIdolBoostMaps = append(DuoIdolBoostMaps, boostMap)
	}

	err = file.WriteJson(m.GetJsonPath(), DuoIdolBoostMaps)
	if err != nil {
		return err
	}

	return nil
}
