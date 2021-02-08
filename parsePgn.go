package pgn

import (
	"log"
	"regexp"
	"strings"
)

// this function is written to parse pgn into a struct
/* Example pgn
[Event "Rated Blitz game"]
[Site "https://lichess.org/dDs8qbu2"]
[Date "2021.02.07"]
[White "unseenzulu"]
[Black "kurian"]
[Result "0-1"]
[UTCDate "2021.02.07"]
[UTCTime "07:14:52"]
[WhiteElo "1070"]
[BlackElo "1181"]
[WhiteRatingDiff "-4"]
[BlackRatingDiff "+4"]
[Variant "Standard"]
[TimeControl "300+0"]
[ECO "A40"]
[Opening "Horwitz Defense"]
[Termination "Normal"]

1. d4 e6 2. Bf4 Nc6 3. e3 d6 4. Bd3 Qf6 5. Nf3 Bd7 6. Nbd2 O-O-O 7. Qe2 d5 8. O-O-O h6 9. h3 g5 10. Be5 Nxe5 11. dxe5 Qe7 12. Nd4 f6 13. exf6 Nxf6 14. Bb5 Bxb5 15. Qxb5 Ne4 16. Qa5 Nxd2 17. Rxd2 Bg7 18. Nb5 b6 19. Qxa7 Qc5 20. Nd4 Bxd4 21. Rxd4 Rd6 22. Qa8+ Kd7 23. Qxh8 Rc6 24. c3 b5 25. Qh7+ Kd6 26. Qxh6 b4 27. Qxg5 bxc3 28. bxc3 Qxc3+ 29. Kb1 Rb6+ 30. Rb4 Rxb4# 0-1
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
	movslic := regexp.MustCompile(`[0-9]+\.`).Split(pgnstring[len(pgnstring)-1], 100)
	for _, val := range movslic {
		move := strings.Split(val, " ")
		if len(move) == 4 {
			log.Println(move[1], "---", move[2])
		} else {
			log.Println(len(move))
		}

	}
	return &PGN{Meta: getmeta}
}
