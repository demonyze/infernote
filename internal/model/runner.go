package model

type Runner struct {
	// Importer
	ChordRocksGuitarImporter Importer[ChordRocksGuitarImport]
	ChordsDbGuitarImporter   Importer[ChordsDbGuitarImport]

	// Exporter
	ChordsArrayExporter Exporter[[]Chord]
	ChordsMapExporter   Exporter[map[string]Chord]
}
