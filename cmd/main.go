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

	chords, err := data.GenerateChords()
	if err != nil {
		fmt.Println("🚫", err)
		return
	}

	exportError := utils.Export(utils.ExportParams[[]model.Chord]{
		FileName: fileName,
		Path:     path,
		Data:     chords,
	})
	if exportError != nil {
		fmt.Println("🚫", exportError)
		return
	}

	fmt.Println("🎉 Files successfully exported")
}
