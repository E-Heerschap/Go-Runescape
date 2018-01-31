package highscores

import (
	"bytes"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/kingpulse/Go-Runescape"
)

type rank struct {

	Name string `json:"name"`
	Score string `json:"score"`
	Rank string `json:"rank"`

}

func GetRankings(skill int64, category int64, amountOfPlayers int64, HttpClient Go_Runescape.HttpClientWrap) (rankings []rank, err error)  {

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