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
	var output string = utils.GetArg(2, "map")

	fmt.Println("🔥 Generating chords... 🎵")

	// Importer
	chordRocksGuitarImporter := model.Importer[model.ChordRocksGuitarImport]{
		Path: "resources/greed/output.json",
	}
	chordDbGuitarImporter := model.Importer[model.ChordsDbGuitarImport]{
		Path: "resources/chords-db/guitar.json",
	}

	// Exporter
	chordsArrayExporter := model.Export[[]model.Chord]{
		FileName: fileName,
		Path:     path,
	}
	chordsMapExporter := model.Export[map[string]model.Chord]{
		FileName: fileName,
		Path:     path,
	}

	runner := model.Runner{
		ChordsDbGuitarImporter:   chordDbGuitarImporter,
		ChordRocksGuitarImporter: chordRocksGuitarImporter,

		ChordsArrayExporter: chordsArrayExporter,
		ChordsMapExporter:   chordsMapExporter,
	}

	generator.Run(output, runner)

	fmt.Println("🎉 Files successfully exported")
}
