package data

import "github.com/demonyze/infernote/internal/model"

var ChordNames = map[model.ChordType]map[model.Language]string{
	model.Major: {
		model.EN: "Major",
		model.DE: "Dur",
	},
	model.Minor: {
		model.EN: "Minor",
		model.DE: "Moll",
	},
	model.Fith: {
		model.EN: "5th",
		model.DE: "Quint",
	},
	model.Seventh: {
		model.EN: "7th",
		model.DE: "Septakkord",
	},
	model.Major7th: {
		model.EN: "Major 7th",
		model.DE: "Dur-Septakkord",
	},
	model.Minor7th: {
		model.EN: "Minor 7th",
		model.DE: "Moll-Septakkord",
	},
	model.Minor7thFlat5th: {
		model.EN: "Minor 7th Flat 5th",
		model.DE: "Halbvermindert-Septakkord",
	},
	model.Diminished: {
		model.EN: "Diminished",
		model.DE: "Vermindert",
	},
	model.Augmented: {
		model.EN: "Augmented",
		model.DE: "Erhöht",
	},
	model.Suspended2nd: {
		model.EN: "Suspended 2nd",
		model.DE: "Suspendiert 2",
	},
	model.Minor6th: {
		model.EN: "Minor 6th",
		model.DE: "Moll Sextakkord",
	},
	model.Sixth: {
		model.EN: "Sixth",
		model.DE: "Sextakkord",
	},
	model.Suspended4th: {
		model.EN: "Suspended 4th",
		model.DE: "Suspendiert 4",
	},
	model.Ninth: {
		model.EN: "Ninth",
		model.DE: "None-Akkord",
	},
	model.Major9th: {
		model.EN: "Major 9th",
		model.DE: "Dur-9-Akkord",
	},
	model.Minor9th: {
		model.EN: "Minor 9th",
		model.DE: "Moll-9-Akkord",
	},
	model.SixthAdd9: {
		model.EN: "Sixth Add 9",
		model.DE: "Sextakkord mit None",
	},
	model.Eleventh: {
		model.EN: "Eleventh",
		model.DE: "11-Akkord",
	},
	model.SeventhFlat5: {
		model.EN: "Seventh Flat 5th",
		model.DE: "Septakkord mit verminderte None",
	},
	model.MinorMajor7th: {
		model.EN: "Minor Major 7th",
		model.DE: "Moll-Dur-Septakkord",
	},
	model.SeventhSuspended4th: {
		model.EN: "Seventh Suspended 4th",
		model.DE: "Suspendiert 4 Septakkord",
	},
	model.Diminished7th: {
		model.EN: "Diminished 7th",
		model.DE: "Vermindert 7 Septakkord",
	},
	model.SeventhSharp5: {
		model.EN: "Seventh Sharp 5th",
		model.DE: "Septakkord mit Erhöhte None",
	},
	model.MinorMajor9th: {
		model.EN: "Minor Major 9th",
		model.DE: "Moll-Dur-9-Akkord",
	},
	model.Minor11th: {
		model.EN: "Minor 11th",
		model.DE: "Moll-11-Akkord",
	},
	model.Major11th: {
		model.EN: "Major 11th",
		model.DE: "Dur-11-Akkord",
	},
}
