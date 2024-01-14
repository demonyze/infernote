package generator

import (
	"fmt"

	"github.com/demonyze/infernote/internal/model"
	"github.com/demonyze/infernote/internal/utils"
)

func Run(fileName string, path string, output string) {
	chordsData, err := utils.Import[model.ChordsDbGuitarImport]("assets/chords-db/guitar.json")
	if err != nil {
		fmt.Println("🚫", err)
		return
	}

	switch output {
	case "array":
		{
			chordsArray := ChordsAsArray(chordsData)
			exportError := utils.Export(utils.ExportParams[[]model.Chord]{
				FileName: fileName,
				Path:     path,
				Data:     chordsArray,
			})
			if exportError != nil {
				fmt.Println("🚫", exportError)
				return
			}
		}
	default:
		{
			chordsMap := ChordsAsMap(chordsData)
			exportError := utils.Export(utils.ExportParams[map[string]model.Chord]{
				FileName: fileName,
				Path:     path,
				Data:     chordsMap,
			})
			if exportError != nil {
				fmt.Println("🚫", exportError)
				return
			}
		}

	}
}
