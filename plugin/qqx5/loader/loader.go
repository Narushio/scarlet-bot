package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/xuri/excelize/v2"

	"github.com/Narushio/scarlet-bot/helper/file"
	"github.com/Narushio/scarlet-bot/helper/image"
	"github.com/Narushio/scarlet-bot/plugin/qqx5/model"
)

const (
	BaseFilePath                 = "resource/qqx5/"
	ExcelFilePath                = BaseFilePath + "kym爆气笔记2023s10.xlsx"
	DuoIdolBoostMapListFilePath  = BaseFilePath + "duo_idol_boost_maps.json"
	SoloIdolBoostMapListFilePath = BaseFilePath + "solo_idol_boost_maps.json"
	SoloIdolBoostMapSheetName    = "星动"
	DuoIdolBoostMapSheetName     = "星动双排"
)

func readJson(jsonFilePath string, model interface{}) {
	b, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, model)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func writeJson(jsonFilePath string, model interface{}) {
	b, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(jsonFilePath, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func BoostMapExcel() {
	f, err := excelize.OpenFile(ExcelFilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	processDuoIdolBoostMap(f)
}

func processSoloIdolBoostMap(f *excelize.File) {
	if file.IsExist(SoloIdolBoostMapListFilePath) {
		readJson(SoloIdolBoostMapListFilePath, &model.SoloIdolBoostMapList)
	} else {
		rows, err := f.GetRows(SoloIdolBoostMapSheetName)
		if err != nil {
			log.Fatal(err)
		}

		for i, _ := range rows {
			if i == 0 {
				continue
			}
			boostMap := &model.SoloIdolBoostMap{}
			colNames := []string{"I", "J"}
			for _, c := range colNames {
				p, err := f.GetPictures(SoloIdolBoostMapSheetName, fmt.Sprintf("%s%d", c, i+1))
				if err != nil {
					log.Fatal(err)
				}

				if len(p) != 0 {
					switch c {
					case "J":
						boostMap.BoostInfo.End = image.ToBase64Src(p[0].File, image.Jpeg)
					default:
						boostMap.BoostInfo.Start = image.ToBase64Src(p[0].File, image.Jpeg)
					}
				}
			}

			// Todo: assign SoloIdolBoostMap
			//if row[0] == "" {
			//	boostMap.Letter = model.SoloIdolBoostMapList[len(model.SoloIdolBoostMapList)-1].Letter
			//} else {
			//	boostMap.Letter = row[0]
			//}
			//boostMap.Title = row[1]
			//boostMap.BoostInfo.Description = row[2]
			//boostMap.BoostInfo.Description = row[5]
			//if len(row) > 8 {
			//	boostMap.Memo = row[8]
			//}
			//model.SoloIdolBoostMapList = append(model.SoloIdolBoostMapList, boostMap)
		}

		writeJson(SoloIdolBoostMapListFilePath, model.SoloIdolBoostMapList)
	}
}

func processDuoIdolBoostMap(f *excelize.File) {
	if file.IsExist(DuoIdolBoostMapListFilePath) {
		readJson(DuoIdolBoostMapListFilePath, &model.DuoIdolBoostMapList)
	} else {
		rows, err := f.GetRows(DuoIdolBoostMapSheetName)
		if err != nil {
			log.Fatal(err)
		}

		for i, row := range rows {
			if i == 0 {
				continue
			}
			boostMap := &model.DuoIdolBoostMap{}
			colNames := []string{"D", "E", "G", "H"}
			for _, c := range colNames {
				p, err := f.GetPictures(DuoIdolBoostMapSheetName, fmt.Sprintf("%s%d", c, i+1))
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
				boostMap.Letter = model.DuoIdolBoostMapList[len(model.DuoIdolBoostMapList)-1].Letter
			} else {
				boostMap.Letter = row[0]
			}
			boostMap.Title = row[1]
			boostMap.PhaseA.Description = row[2]
			boostMap.PhaseB.Description = row[5]
			if len(row) > 8 {
				boostMap.Memo = row[8]
			}
			model.DuoIdolBoostMapList = append(model.DuoIdolBoostMapList, boostMap)
		}

		writeJson(DuoIdolBoostMapListFilePath, model.DuoIdolBoostMapList)
	}
}
