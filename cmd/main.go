package main

import (
	"fmt"

	"github.com/demonyze/infernote/internal/data"
	"github.com/demonyze/infernote/internal/model"
	"github.com/demonyze/infernote/internal/utils"
)

func main() {

	var fileName string = utils.GetArg(0, "chords.json")
	var path string = utils.GetArg(1, "sample")

	fmt.Println("🔥 Generating chords... 🎵")

	chordsData, err := utils.Import[model.ChordsDbGuitarImport]("assets/chords-db/guitar.json")
	if err != nil {
		fmt.Println("🚫", err)
		return
	}

	chordsMap := data.ChordsAsMap(chordsData)

	exportError := utils.Export(utils.ExportParams[map[string]model.Chord]{
		FileName: fileName,
		Path:     path,
		Data:     chordsMap,
	})
	if exportError != nil {
		fmt.Println("🚫", exportError)
		return
	}

	fmt.Println("🎉 Files successfully exported")
}
