package model

import (
	"encoding/base64"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func init() {
	initDuoBurstTableList()
}

type DuoBurstTable struct {
	Letter string    `json:"letter"`
	Title  string    `json:"title"`
	PhaseA BurstInfo `json:"phase_a"`
	PhaseB BurstInfo `json:"phase_b"`
	Memo   string    `json:"memo"`
}

type Base64Src string

type BurstInfo struct {
	Description string    `json:"description"`
	Start       Base64Src `json:"start"`
	End         Base64Src `json:"end"`
}

var DuoBurstTableList []*DuoBurstTable

func initDuoBurstTableList() {
	f, err := excelize.OpenFile("resource/kym爆气笔记2023s10.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows("星动双排")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		duoBurstTable := &DuoBurstTable{}
		colNames := []string{"D", "E", "G", "H"}
		for _, c := range colNames {
			p, err := f.GetPictures("星动双排", fmt.Sprintf("%s%d", c, i+1))
			if err != nil {
				fmt.Println(err)
				continue
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
		duoBurstTable.Letter = row[0]
		duoBurstTable.Title = row[1]
		duoBurstTable.PhaseA.Description = row[2]
		duoBurstTable.PhaseB.Description = row[5]
		if len(row) > 8 {
			duoBurstTable.Memo = row[8]
		}
		DuoBurstTableList = append(DuoBurstTableList, duoBurstTable)
	}
}

func toBase64Src(fileContents []byte) Base64Src {
	base64Str := base64.StdEncoding.EncodeToString(fileContents)
	return Base64Src(fmt.Sprintf("data:image/jpeg;base64, %s", base64Str))
}
