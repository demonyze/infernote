package model

// Abbrevation

type ChordType = string

const (
	Major               ChordType = "Major"
	Minor               ChordType = "Minor"
	Fith                ChordType = "Fith"
	Seventh             ChordType = "Seventh"
	Major7th            ChordType = "Major7th"
	Minor7th            ChordType = "Minor7th"
	Minor7thFlat5th     ChordType = "Minor7thFlat5th"
	Diminished          ChordType = "Diminished"
	Augmented           ChordType = "Augmented"
	Suspended2nd        ChordType = "Suspended2nd"
	Minor6th            ChordType = "Minor6th"
	Sixth               ChordType = "Sixth"
	Suspended4th        ChordType = "Suspended4th"
	Ninth               ChordType = "Ninth"
	Major9th            ChordType = "Major9th"
	Minor9th            ChordType = "Minor9th"
	SixthAdd9           ChordType = "SixthAdd9"
	Eleventh            ChordType = "Eleventh"
	SeventhFlat5        ChordType = "SeventhFlat5"
	MinorMajor7th       ChordType = "MinorMajor7th"
	SeventhSuspended4th ChordType = "SeventhSuspended4th"
	Diminished7th       ChordType = "Diminished7th"
	SeventhSharp5       ChordType = "SeventhSharp5"
	MinorMajor9th       ChordType = "MinorMajor9th"
	Minor11th           ChordType = "Minor11th"
	Major11th           ChordType = "Major11th"
)

type AbbreviationType = map[Language][]string

// ChorNames

type ChordNameType = map[Language]string
