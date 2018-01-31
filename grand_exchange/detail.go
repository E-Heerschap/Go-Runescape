package grand_exchange

import (
	"bytes"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/kingpulse/Go-Runescape"
)

type ItemJson struct {
	Item ItemDetail `json:"item"`
}

type ItemDetail struct {
	Icon string `json:"icon"`
	Icon_large string `json:"icon_large"`
	Id int64 `json:"id"`
	ItemType string `json:"type"`
	TypeIconURL string `json:"typeIcon"`
	Name string `json:"name"`
	Description string `json:"description"`
	Current timeTrendPrice `json:"current"`
	Today timeTrendPrice `json:"today"`
	Day30 timeTrendPercentage `json:"day30"`
	Day90 timeTrendPercentage `json:"day90"`
	Day180 timeTrendPercentage `json:"day180"`
}

type timeTrendPrice struct {
	Trend string `json:"trend"`
	Price string `json:"price"`
}

type timeTrendPercentage struct {
	Trend string `json:"trend"`
	Change string `json:"change"`
}

func GetItemDetail(itemID int64, HttpClient Go_Runescape.HttpClientWrap) (ItemDetail, error){

	//Creating URL for request.
	stringWriter := bytes.NewBufferString("http://services.runescape.com/m=itemdb_rs/api/catalogue/detail.json?item=")
	itemIdString := strconv.FormatInt(itemID, 10)
	stringWriter.WriteString(itemIdString)

	resp, err := HttpClient.Get(stringWriter.String())

	if err != nil {
		return ItemDetail{}, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return ItemDetail{}, err
	}

	item := ItemJson{}

	err = json.Unmarshal(respBytes, &item)

	return item.Item, err
}


