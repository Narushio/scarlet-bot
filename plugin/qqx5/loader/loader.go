package loader

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"

	"github.com/Narushio/scarlet-bot/plugin/qqx5/model"
)

const BoostMapExcelPath = "resource/qqx5/boost_maps_2023s13.xlsx"

type BoostMap interface {
	UnmarshalFromJson() error
	UnmarshalFromExcel(f *excelize.File) error
}

func AllBoostMapList() error {
	f, err := excelize.OpenFile(BoostMapExcelPath)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	processBoostMap[*model.DuoIdolBoostMap](f, &model.DuoIdolBoostMap{})
	processBoostMap[*model.SoloIdolBoostMap](f, &model.SoloIdolBoostMap{})
	return nil
}

func processBoostMap[T BoostMap](f *excelize.File, m T) {
	err := m.UnmarshalFromJson()
	if err != nil {
		var pathErr *os.PathError
		if errors.As(err, &pathErr) {
			err := m.UnmarshalFromExcel(f)
			if err != nil {
				log.Fatalf(err.Error())
			}
		}
	}
}
