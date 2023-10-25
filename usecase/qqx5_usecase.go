package usecase

import (
	"encoding/base64"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func init() {
	f, err := excelize.OpenFile("resources/kym爆气笔记2023s10.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows("星动双排")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

type DuoRankedSongBurstMeter struct {
	Letter string
	Title  string
	PhaseA BurstInfo
	PhaseB BurstInfo
	Memo   string
}

type BurstInfo struct {
	Description string
	Start       base64.Encoding
	End         base64.Encoding
}
