package model

type AlternativeName = string

const (
	Chord5th AlternativeName = "Chord5th"
	Chord7th AlternativeName = "Chord7th"
)

type AlternativeNamesType = map[Language][]string

// Abbrevation

type Abbreviation = string

const (
	Major   Abbreviation = "major"
	Minor   Abbreviation = "minor"
	Fith    Abbreviation = "fith"
	Seventh Abbreviation = "seventh"
)

type AbbreviationType = map[Language][]string

// ChorNames

type ChordName = string

const (
	AMajor ChordName = "amajor"
	AMinor ChordName = "aminor"
	A5th   ChordName = "a5th"
	A7th   ChordName = "a7th"
	// Minor   Abbreviation = "minor"
	// Fith    Abbreviation = "fith"
	// Seventh Abbreviation = "seventh"
)

type ChordNameType = map[Language]string
