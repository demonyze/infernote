package main

import (
	"fmt"

	"github.com/demonyze/infernote/internal/data"
)

func main() {
	fmt.Println("🔥 Generating chords... 🎵")
	err := data.ExportData("sample", "chords.json")
	if err != nil {
		fmt.Println("🚫", err)
		return
	}
	fmt.Println("🎉 Files succesfully exported")
}
