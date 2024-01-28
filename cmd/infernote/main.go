package main

import (
	"fmt"
	"os"

	"github.com/demonyze/infernote/internal/generator"
	"github.com/demonyze/infernote/internal/model"
	"github.com/demonyze/infernote/internal/utils"
)

func main() {

	args := os.Args[1:]
	params := utils.GetParams(args)

	fmt.Println("🔥 Generating chords... 🎵")

	// Importer
	chordRocksGuitarImporter := model.Import[model.ChordRocksGuitarImport]{
		Path: "resources/greed/output.json",
	}
	chordDbGuitarImporter := model.Import[model.ChordsDbGuitarImport]{
		Path: "resources/chords-db/guitar.json",
	}
	languageImporter := model.Import[model.LanguageImport]{
		Path: fmt.Sprintf("lang/%s.json", params.Language),
	}

	// Exporter
	exporter := model.Export[model.Infernote]{
		FileName: params.FileName,
		Path:     params.FilePath,
	}

	runner := model.Runner{
		ChordsDbGuitarImporter:   chordDbGuitarImporter,
		ChordRocksGuitarImporter: chordRocksGuitarImporter,
		LanguageImporter:         languageImporter,

		InfernoteExporter: exporter,
	}

	err := generator.Run(runner)
	if err != nil {
		fmt.Println("🚫", err)
		return
	}

	fmt.Println("🎉 Files successfully exported")
}
