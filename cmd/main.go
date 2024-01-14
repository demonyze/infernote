package main

import (
	"fmt"

	"github.com/demonyze/infernote/internal/generator"
	"github.com/demonyze/infernote/internal/utils"
)

func main() {

	var fileName string = utils.GetArg(0, "chords.json")
	var path string = utils.GetArg(1, "sample")
	var output string = utils.GetArg(2, "map")

	fmt.Println("🔥 Generating chords... 🎵")

	generator.Run(fileName, path, output)

	fmt.Println("🎉 Files successfully exported")
}
