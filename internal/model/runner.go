package model

type Runner struct {
	LanguageImport LanguageImport

	// -- Importer --

	// Chords
	ChordRocksGuitarImporter Importer[ChordRocksGuitarImport]
	ChordsDbGuitarImporter   Importer[ChordsDbGuitarImport]

	// -- Exporter --
	InfernoteExporter Exporter[Infernote]
}
