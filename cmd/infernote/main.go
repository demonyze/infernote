package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/demonyze/infernote/internal/generator"
	"github.com/demonyze/infernote/internal/model"
	"github.com/demonyze/infernote/internal/utils"
)

//go:embed chords-db/guitar.json
var chordsDBImport embed.FS

//go:embed greed/greed.json
var chordRocksImport embed.FS

func main() {

	args := os.Args[1:]
	params := utils.GetParams(args)

	fmt.Println("🔥 Generating chords... 🎵")

	chordsDBData, errs := chordsDBImport.ReadFile("chords-db/guitar.json")
	if errs != nil {
		panic(errs)
	}
	chordRocksImport, errs := chordRocksImport.ReadFile("greed/greed.json")
	if errs != nil {
		panic(errs)
	}

	// Importer
	chordRocksGuitarImporter := model.Import[model.ChordRocksGuitarImport]{
		Data: chordRocksImport,
	}
	chordDbGuitarImporter := model.Import[model.ChordsDbGuitarImport]{
		Data: chordsDBData,
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
