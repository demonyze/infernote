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

	fmt.Println("🔥 Generating chords... 🎵")

	// Importer
	chordRocksGuitarImporter := model.Import[model.ChordRocksGuitarImport]{
		Path: "resources/greed/output.json",
	}
	chordDbGuitarImporter := model.Import[model.ChordsDbGuitarImport]{
		Path: "resources/chords-db/guitar.json",
	}

	chordsMapExporter := model.Export[map[string]model.Chord]{
		FileName: fileName,
		Path:     path,
	}

	runner := model.Runner{
		ChordsDbGuitarImporter:   chordDbGuitarImporter,
		ChordRocksGuitarImporter: chordRocksGuitarImporter,

		ChordsMapExporter: chordsMapExporter,
	}

	generator.Run(runner)

	fmt.Println("🎉 Files successfully exported")
}
