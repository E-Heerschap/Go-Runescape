package Go_Runescape

import (
	"net/url"
	"net/http"
	"strings"
	"bufio"
	"strconv"
	"fmt"
	"github.com/kingpulse/Go-Runescape"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

//This will store the highscore_constants information of a player.
type PlayerHighscores struct {

	//Maps are being used so
	Levels []int64
	XP []int64
	Ranks []int64

}

//GetPlayerHighscores gets a PlayerHighscores object relating to the RS3 player from the name passed.
func GetPlayerHighscores (playerName string, highscoreType string, httpClient Go_Runescape.IHttpClient) (rsph PlayerHighscores, err error) {

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

type rank struct {

	Name string `json:"name"`
	Score string `json:"score"`
	Rank string `json:"rank"`

}

func GetRankings(skill int64, category int64, amountOfPlayers int64, HttpClient Go_Runescape.IHttpClient) (rankings []rank, err error)  {

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
