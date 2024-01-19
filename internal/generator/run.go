package generator

import (
	"fmt"

	"github.com/demonyze/infernote/internal/model"
	"github.com/demonyze/infernote/internal/utils"
)

func Run(fileName string, path string, output string) {
	chordRocksImport, err := utils.Import[model.ChordRocksGuitarImport]("resources/greed/output.json")
	if err != nil {
		fmt.Println("🚫", err)
		return
	}

	chordsDbImport, err := utils.Import[model.ChordsDbGuitarImport]("resources/chords-db/guitar.json")
	if err != nil {
		fmt.Println("🚫", err)
		return
	}

	switch output {
	case "array":
		{
			chordsArray := ChordsAsArray(CreateChordParams{
				chordsDbImport,
				chordRocksImport,
			})
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
			chordsMap := ChordsAsMap(CreateChordParams{
				chordsDbImport,
				chordRocksImport,
			})
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
