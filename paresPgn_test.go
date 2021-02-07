package pgn

import (
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	testinput := `[Event "Rated Blitz game"]
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
	[Annotator "lichess.org"]`
	meta := Parse([]byte(testinput))
	log.Println(meta)
}
