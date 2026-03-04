package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/demonyze/infernote/internal/constants"
	"github.com/demonyze/infernote/internal/fetcher"
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

	fmt.Println("Importing embedded chords-db data...")
	chordsDBData, errs := chordsDBImport.ReadFile("chords-db/guitar.json")
	if errs != nil {
		panic(errs)
	}
	fmt.Println("✅ Successfully imported embedded chords-db data")

	fmt.Println("Importing embedded chord.rocks data...")
	chordRocksImport, errs := chordRocksImport.ReadFile("greed/greed.json")
	if errs != nil {
		panic(errs)
	}
	fmt.Println("✅ Successfully imported embedded chord.rocks data")

	fmt.Println("Fetching languages...")
	languageImports, errs := fetcher.GetLanguages(params.Languages, constants.URL_GITHUB_LANG)
	if errs != nil {
		panic(errs)
	}
	fmt.Println("✅ Successfully fetched languages from ", constants.URL_GITHUB_LANG)

	// Importer
	chordRocksGuitarImporter := model.Import[model.ChordRocksGuitarImport]{
		Data: chordRocksImport,
	}
	chordDbGuitarImporter := model.Import[model.ChordsDbGuitarImport]{
		Data: chordsDBData,
	}

	// Exporter
	exporter := model.Export[model.Infernote]{
		FileName: params.FileName,
		Path:     params.FilePath,
	}

	runner := model.Runner{
		LanguageImports:          languageImports,
		ChordsDbGuitarImporter:   chordDbGuitarImporter,
		ChordRocksGuitarImporter: chordRocksGuitarImporter,

		InfernoteExporter: exporter,
	}

	err := generator.Run(runner)
	if err != nil {
		fmt.Println("🚫", err)
		return
	}

	fmt.Println("🎉 Files successfully exported")
}
