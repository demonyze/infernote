package main

import (
	"fmt"

	"github.com/demonyze/infernote/internal/generator"
	"github.com/demonyze/infernote/internal/model"
	"github.com/demonyze/infernote/internal/utils"
)

func main() {

	var fileName string = utils.GetArg(0, "chords.json")
	var path string = utils.GetArg(1, "sample")
	var language string = utils.GetArg(2, "en")

	fmt.Println("🔥 Generating chords... 🎵")

	// Importer
	chordRocksGuitarImporter := model.Import[model.ChordRocksGuitarImport]{
		Path: "resources/greed/output.json",
	}
	chordDbGuitarImporter := model.Import[model.ChordsDbGuitarImport]{
		Path: "resources/chords-db/guitar.json",
	}
	languageImporter := model.Import[model.LanguageImport]{
		Path: fmt.Sprintf("lang/%s.json", language),
	}

	// Exporter
	exporter := model.Export[model.Infernote]{
		FileName: fileName,
		Path:     path,
	}

	runner := model.Runner{
		ChordsDbGuitarImporter:   chordDbGuitarImporter,
		ChordRocksGuitarImporter: chordRocksGuitarImporter,
		LanguageImporter:         languageImporter,

		InfernoteExporter: exporter,
	}

	generator.Run(runner)

	fmt.Println("🎉 Files successfully exported")
}
