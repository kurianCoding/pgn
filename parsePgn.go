package pgn

import "strings"

// this function is written to parse pgn into a struct
/* Example pgn
1. d4 {variation} e6 {variaion} 2. bf4 {variation} nc6 {variation}
*/
type PGN struct {
	Meta  map[string]string `json:"metadata"`
	Moves []struct {
		White string `json:"white"`
		Black string `json:"black"`
	} `json:"moves"`
}

func Parse(pgnbytes []byte) *PGN {
	pgnstring := strings.Split(string(pgnbytes), "\n")
	meta := pgnstring[0 : len(pgnstring)-2]

	getmeta := make(map[string]string, 0)
	for i := 0; i < len(meta)-1; i++ {
		line := meta[i]
		key, value := strings.Split(line, " ")[0], strings.Split(line, " ")[1]
		getmeta[key] = value
	}
	return &PGN{Meta: getmeta}
}
