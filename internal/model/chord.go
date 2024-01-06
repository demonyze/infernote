package model

type Chord struct {
	Id               string                   `json:"id"`
	Name             string                   `json:"name"`
	Type             ChordType                `json:"type"`
	AlternativeNames map[Language][]string    `json:"alternative_names"`
	Abbreviation     map[Language][]string    `json:"abbrevation"`
	Notes            []Note                   `json:"notes"`
	Root             Note                     `json:"root"`
	GuitarPositions  []ChordsDbGuitarPosition `json:"guitar_positions"`
}

// type GuitarVariation struct {
// 	Tuning   Tuning `json:"tuning"`
// 	Frets    []int  `json:"frets"`
// 	Fingers  []int  `json:"fingers"`
// 	Barres   []int  `json:"barres"`
// 	Capo     bool   `json:"capo"`
// 	BaseFret int    `json:"baseFret"`
// }

// {
// 	"frets": [-1, 1, 3, 3, 3, 1],
// 	"fingers": [0, 1, 2, 3, 4, 1],
// 	"barres": [1],
// 	"capo": true,
// 	"baseFret": 3,
// 	"midi": [48, 55, 60, 64, 67]
//   },
