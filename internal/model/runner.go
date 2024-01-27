package model

type Runner struct {
	// Importer
	ChordRocksGuitarImporter Importer[ChordRocksGuitarImport]
	ChordsDbGuitarImporter   Importer[ChordsDbGuitarImport]

	// Exporter
	ChordsMapExporter Exporter[map[string]Chord]
}
