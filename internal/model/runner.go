package model

type Runner struct {
	// -- Importer --

	// Chords
	ChordRocksGuitarImporter Importer[ChordRocksGuitarImport]
	ChordsDbGuitarImporter   Importer[ChordsDbGuitarImport]

	//Language
	LanguageImporter Importer[LanguageImport]

	// -- Exporter --
	InfernoteExporter Exporter[Infernote]
}
