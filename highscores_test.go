package Go_Runescape

import (
	"testing"
	"strconv"
	"github.com/kingpulse/Go-Runescape/highscore_constants"
	"net/http"
)

//TestPlayerHighScores is some basic tests to check if the player information is being correctly downloaded.
//(Primarily to test CI to be honest).
//Player names that are passed in as parameters to the GetPlayerHighscores function are
//#1 players as of 30/01/2018
func TestPlayerHighScores(t *testing.T) {

	//Creating http client to use for requests
	testClient := http.DefaultClient

	/*
	TESTING RS3 Player
	 */

	hs, err := GetPlayerHighscores("le me", highscore_constants.RS3PLAYER, testClient)
	levelsCheck(t, hs, "RS3")

	if err != nil {
		t.Error("Failed to GetPlayerHighscores. Error: " + err.Error())
	}

	/*
	TESTING OSRS Player
	 */

	hs, err = GetPlayerHighscores("Lynx Titan", highscore_constants.OSRSPLAYER, testClient)
	levelsCheck(t, hs, "OSRS")

	if err != nil {
		t.Error("Failed to GetPlayerHighscores. Error: " + err.Error())
	}

	//Testing failure during GET()
	failureGetClient := failGetHttpClient{}

	_, err = GetPlayerHighscores("Lynx Titan", highscore_constants.OSRSPLAYER, failureGetClient)

	if err == nil {
		t.Error("GetPlayerHighscores failed to return errors from GET request.")
	}

	ijClient := invalidJsonHttpClient{}

	_, err = GetPlayerHighscores("Lynx Titan", highscore_constants.OSRSPLAYER, ijClient)

	if err == nil {
		t.Error("GetPlayerHighscores failed to regonise error in json")
	}

}

func levelsCheck(t *testing.T, hs PlayerHighscores, playerType string) {
	if hs.Levels == nil || hs.Ranks == nil || hs.XP == nil {
		t.Error("Failed to get " + playerType + " player highscore_constants object.")
	}

	//Checking that levels are correct.
	for _, val := range hs.Levels {
		if val < 99 {
			t.Error("Incorrect levels are being grabbed for " + playerType + " players.")
		}
	}
}

func TestPlayerRanks(t *testing.T) {

	//Creating default http client to send requests.
	testClient := http.DefaultClient

	strengthRankings, err := GetRankings(highscore_constants.STRENGTH, 0, 5, testClient)

	if err != nil {
		t.Error("Failed to get player rankings for strength.")
	}

	if strengthRankings[0].Rank != strconv.Itoa(1) {
		t.Error("Failed to get top player ranking in strength.")
	}

	failureClient := failGetHttpClient{}

	_, err = GetRankings(highscore_constants.STRENGTH, 0, 5, failureClient)

	if err == nil {
		t.Error("GetRankings failed to return errors from GET request.")
	}

	ijClient := invalidJsonHttpClient{}

	_, err = GetRankings(highscore_constants.STRENGTH, 0, 5, ijClient)

	if err == nil {
		t.Error("GetRaknings failed to recognise json errors.")
	}

}
