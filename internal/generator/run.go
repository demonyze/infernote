package generator

import (
	"github.com/demonyze/infernote/internal/model"
)

func Run(runner model.Runner) error {
	chordRocksImport, err := runner.ChordRocksGuitarImporter.Import()
	if err != nil {
		return err
	}

	chordsDbImport, err := runner.ChordsDbGuitarImporter.Import()
	if err != nil {
		return err
	}

	chordsMap := ChordsAsMap(CreateChordParams{
		chordsDbImport,
		chordRocksImport,
	})

	languageImport, err := runner.LanguageImporter.Import()
	if err != nil {
		return err
	}

	types := Types(languageImport.Types)

	infernote := model.Infernote{
		Language: languageImport.Language,
		Types:    types,
		Chords:   chordsMap,
	}

	exportError := runner.InfernoteExporter.Export(infernote)
	if exportError != nil {
		return exportError
	}

	return nil
}
