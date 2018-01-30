package highscores

import (
	"testing"
	"strconv"
	"github.com/kingpulse/Go-Runescape/highscores/highscore_constants"
)

//TestPlayerHighScores is some basic tests to check if the player information is being correctly downloaded.
//(Primarily to test CI to be honest).
//Player names that are passed in as parameters to the GetPlayerHighscores function are
//#1 players as of 30/01/2018
func TestPlayerHighScores(t *testing.T){

	/*
	TESTING RS3 Player
	 */

	hs := GetPlayerHighscores("le me", highscore_constants.RS3PLAYER)
	levelsCheck(t, hs, "RS3")


	/*
	TESTING OSRS Player
	 */

	 hs = GetPlayerHighscores("Lynx Titan", highscore_constants.OSRSPLAYER)
	levelsCheck(t, hs, "OSRS")
}

func levelsCheck(t *testing.T, hs PlayerHighscores, playerType string){
	if (hs.Levels == nil || hs.Ranks == nil || hs.XP == nil) {
		t.Error("Failed to get " + playerType + " player highscores object.")
	}

	//Checking that levels are correct.
	for _, val := range hs.Levels {
		if val < 99 {
			t.Error("Incorrect levels are being grabbed for " + playerType + " players.")
		}
	}
}

func TestPlayerRanks(t *testing.T){

	strengthRankings, err := getRankings(highscore_constants.STRENGTH, 0, 5)

	if err != nil {
		t.Error("Failed to get player rankings for strength.")
	}

	if strengthRankings[0].Rank != strconv.Itoa(1) {
		t.Error("Failed to get top player ranking in strength.")
	}

}