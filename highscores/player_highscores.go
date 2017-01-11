package highscores

import (
	"net/url"
	"net/http"
	"strings"
	"bufio"
	"strconv"
	"fmt"
)

//This will store the highscores information of a player.
type PlayerHighscores struct {

	//Maps are being used so
	Levels []int64
	XP []int64
	Ranks []int64

}

//GetPlayerHighscores gets a PlayerHighscores object relating to the RS3 player from the name passed.
func GetPlayerHighscores (playerName string, highscoreType string) (rsph PlayerHighscores) {

	Url, _ := url.Parse("http://services.runescape.com/")

	Url.Path += highscoreType
	parameters := url.Values{}
	parameters.Add("player", playerName)
	Url.RawQuery = parameters.Encode()

	resp, err := http.Get(Url.String())

	if err != nil {
		fmt.Println("Go-Runescape: Failed get request with url: " + Url.String())
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


	return rsph
}
