package model

type Runner struct {
	LanguageImports []LanguageImport

	// -- Importer --

	// Chords
	ChordRocksGuitarImporter Importer[ChordRocksGuitarImport]
	ChordsDbGuitarImporter   Importer[ChordsDbGuitarImport]

	// -- Exporter --
	InfernoteExporter Exporter[Infernote]
}
