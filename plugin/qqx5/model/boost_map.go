package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Narushio/scarlet-bot/helper/file"
	"github.com/xuri/excelize/v2"
)

func init() {
	initDuoBoostMaps()
}

type Mode int

type Pattern int

const (
	Solo Mode = iota
	Duo
)

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

type DuoBoostMap struct {
	Letter string    `json:"letter"`
	Title  string    `json:"title"`
	PhaseA BoostInfo `json:"phase_a"`
	PhaseB BoostInfo `json:"phase_b"`
	Memo   string    `json:"memo"`
}

type Base64Src string

type BoostInfo struct {
	Description string    `json:"description,omitempty"`
	Start       Base64Src `json:"start,omitempty"`
	End         Base64Src `json:"end,omitempty"`
}

var DuoBoostMaps []*DuoBoostMap

func initDuoBoostMaps() {
	jsonFN := "resource/qqx5/kym爆气笔记2023s10.json"
	xlsxFN := "resource/qqx5/kym爆气笔记2023s10.xlsx"
	if file.IsExist(jsonFN) {
		b, err := ioutil.ReadFile(jsonFN)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(b, &DuoBoostMaps)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	f, err := excelize.OpenFile(xlsxFN)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows("星动双排")
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		duoBurstTable := &DuoBoostMap{}
		colNames := []string{"D", "E", "G", "H"}
		for _, c := range colNames {
			p, err := f.GetPictures("星动双排", fmt.Sprintf("%s%d", c, i+1))
			if err != nil {
				log.Fatal(err)
			}

			if len(p) != 0 {
				switch c {
				case "E":
					duoBurstTable.PhaseA.End = toBase64Src(p[0].File)
				case "G":
					duoBurstTable.PhaseB.Start = toBase64Src(p[0].File)
				case "H":
					duoBurstTable.PhaseB.End = toBase64Src(p[0].File)
				default:
					duoBurstTable.PhaseA.Start = toBase64Src(p[0].File)
				}
			}
		}

		if row[0] == "" {
			duoBurstTable.Letter = DuoBoostMaps[len(DuoBoostMaps)-1].Letter
		} else {
			duoBurstTable.Letter = row[0]
		}
		duoBurstTable.Title = row[1]
		duoBurstTable.PhaseA.Description = row[2]
		duoBurstTable.PhaseB.Description = row[5]
		if len(row) > 8 {
			duoBurstTable.Memo = row[8]
		}
		DuoBoostMaps = append(DuoBoostMaps, duoBurstTable)
	}

	b, err := json.Marshal(DuoBoostMaps)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(jsonFN, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func toBase64Src(fileContents []byte) Base64Src {
	base64Str := base64.StdEncoding.EncodeToString(fileContents)
	return Base64Src(fmt.Sprintf("data:image/jpeg;base64, %s", base64Str))
}
