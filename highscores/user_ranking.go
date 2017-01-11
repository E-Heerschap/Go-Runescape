package highscores

import (
	"bytes"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type rank struct {

	name string `json:"name"`
	score int `json:"score"`
	rank int `json:"rank"`

}

func getRankings(skill int64, category int64, amountOfPlayers int64) (rankings []rank, err error)  {

	stringWriter := bytes.NewBufferString("http://services.runescape.com/m=hiscore/ranking.json?table=")
	stringWriter.WriteString(strconv.FormatInt(skill, 10))
	stringWriter.WriteString("&category=")
	stringWriter.WriteString(strconv.FormatInt(category, 10))
	stringWriter.WriteString("&size=")
	stringWriter.WriteString(strconv.FormatInt(amountOfPlayers, 10))

	resp, err := http.Get(stringWriter.String())

	if resp != nil {
		return rankings, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return rankings, err
	}

	err = json.Unmarshal(respBytes, rankings)

	return rankings, err
}