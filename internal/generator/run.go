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

	chordTypes := Types(runner.LanguageImports)
	chordsMap := ChordsAsMap(CreateChordParams{
		chordsDbImport,
		chordRocksImport,
		runner.LanguageImports,
	})

	languages := make([]string, 0, len(runner.LanguageImports))
	for _, imp := range runner.LanguageImports {
		languages = append(languages, imp.Language)
	}

	infernote := model.Infernote{
		Languages: languages,
		Types:     chordTypes,
		Chords:    chordsMap,
	}

	exportError := runner.InfernoteExporter.Export(infernote)
	if exportError != nil {
		return exportError
	}

	return nil
}
