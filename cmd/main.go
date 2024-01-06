package main

import (
	"fmt"

	"github.com/demonyze/infernote/internal/data"
	"github.com/demonyze/infernote/internal/model"
	"github.com/demonyze/infernote/internal/utils"
)

func main() {
	fmt.Println("🔥 Generating chords... 🎵")
	chords, err := data.GenerateChords()
	if err != nil {
		fmt.Println("🚫", err)
		return
	}
	exportError := utils.Export[[]model.Chord](utils.ExportParams[[]model.Chord]{
		Path:     "sample",
		FileName: "chords.json",
		Data:     chords,
	})
	if exportError != nil {
		fmt.Println("🚫", exportError)
		return
	}
	fmt.Println("🎉 Files succesfully exported")
}
