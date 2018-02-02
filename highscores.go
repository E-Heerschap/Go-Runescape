package Go_Runescape

//Contains functions and structs to get information from the Runescape highscores api.
//Author: Edwin Heerschap

import (
	"net/url"
	"strings"
	"bufio"
	"strconv"

	"bytes"
	"io/ioutil"
	"encoding/json"
)

//This will store the highscore_constants information of a player.
//The skill indexes can be found in the highscore_constants package.
//For example Levels[highscore_constants.MINING] would give the mining level of the player.
type PlayerHighscores struct {

	Levels []int64
	XP     []int64
	Ranks  []int64

}

//GetPlayerHighscores gets a players highscores information. The highscoreType is the type of scoreboard to search.
//Acceptable values are stored as constants in the highscore_constants package. These constants are:
//
//highscore_constants.RS3PLAYER
//
//highscore_constants.RS3IRONMAN
//
//highscore_constants.RS3HARDCOREIRONMAN
//
//highscore_constants.OSRSPLAYER
//
//highscore_constants.OSRSIRONMAN
//
//highscore_constants.OSRSULTIMATEIRONMAN
func GetPlayerHighscores(playerName string, highscoreType string, httpClient IHttpClient) (rsph PlayerHighscores, err error) {

	Url, _ := url.Parse("http://services.runescape.com/")

	Url.Path += highscoreType
	parameters := url.Values{}
	parameters.Add("player", playerName)
	Url.RawQuery = parameters.Encode()

	resp, err := httpClient.Get(Url.String())

	if err != nil {
		return rsph, err
	}

	scanner := bufio.NewScanner(resp.Body)

	rsph.Levels = make([]int64, 0)
	rsph.Ranks = make([]int64, 0)
	rsph.XP = make([]int64, 0)

	for scanner.Scan() {
		str := strings.Trim(scanner.Text(), " \n")
		tokens := strings.Split(str, ",")
		if len(tokens) == 3 {
			rank, _ := strconv.ParseInt(tokens[0], 10, 64)
			rsph.Ranks = append(rsph.Ranks, rank)
			level, _ := strconv.ParseInt(tokens[1], 10, 64)
			rsph.Levels = append(rsph.Levels, level)
			XP, _ := strconv.ParseInt(tokens[2], 10, 64)
			rsph.XP = append(rsph.XP, XP)
		}

	}

	return rsph, nil
}

//Contains the list of ranked players on the highscores.
type Rank struct {
	Name  string `json:"name"`
	Score string `json:"score"`
	Rank  string `json:"rank"`
}

//GetRankings gets the rankings for a passed skill. The user is advised to use the highscore_constants package to specify
//the skill.
func GetRankings(skill int64, category int64, amountOfPlayers int64, HttpClient IHttpClient) (rankings []Rank, err error) {

	stringWriter := bytes.NewBufferString("http://services.runescape.com/m=hiscore/ranking.json?table=")
	stringWriter.WriteString(strconv.FormatInt(skill, 10))
	stringWriter.WriteString("&category=")
	stringWriter.WriteString(strconv.FormatInt(category, 10))
	stringWriter.WriteString("&size=")
	stringWriter.WriteString(strconv.FormatInt(amountOfPlayers, 10))

	resp, err := HttpClient.Get(stringWriter.String())

	if err != nil {
		return rankings, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return rankings, err
	}

	err = json.Unmarshal(respBytes, &rankings)

	return rankings, err
}
