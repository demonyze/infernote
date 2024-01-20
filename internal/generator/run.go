package generator

import (
	"fmt"

	"github.com/demonyze/infernote/internal/model"
)

func Run(output string, runner model.Runner) error {
	chordRocksImport, err := runner.ChordRocksGuitarImporter.Import()
	if err != nil {
		fmt.Println("🚫", err)
		return err
	}

	chordsDbImport, err := runner.ChordsDbGuitarImporter.Import()
	if err != nil {
		fmt.Println("🚫", err)
		return err
	}

	switch output {
	case "array":
		{
			chordsArray := ChordsAsArray(CreateChordParams{
				chordsDbImport,
				chordRocksImport,
			})

			exportError := runner.ChordsArrayExporter.Export(chordsArray)
			if exportError != nil {
				fmt.Println("🚫", exportError)
				return exportError
			}
		}
	default:
		{
			chordsMap := ChordsAsMap(CreateChordParams{
				chordsDbImport,
				chordRocksImport,
			})

			exportError := runner.ChordsMapExporter.Export(chordsMap)
			if exportError != nil {
				fmt.Println("🚫", exportError)
				return exportError
			}
		}

	}

	return nil
}
