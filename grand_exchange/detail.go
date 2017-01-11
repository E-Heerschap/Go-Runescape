package grand_exchange

import (
	"bytes"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type ItemDetail struct {
	Item item `json:"item"`
	Current timeTrendPrice `json:"current"`
	Today timeTrendPrice `json:"today"`
	Day30 timeTrendPercentage `json:"day30"`
	Day90 timeTrendPercentage `json:"day90"`
	Day180 timeTrendPercentage `json:"day180"`
}

type item struct {
	Icon string `json:"icon"`
	Icon_large string `json:"icon_large"`
	Id int64 `json:"id"`
	ItemType string `json:"type"`
	TypeIconURL string `json:"typeIcon"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type timeTrendPrice struct {
	Trend string `json:"trend"`
	Price string `json:"price"`
}

type timeTrendPercentage struct {
	Trend string `json:"trend"`
	Change string `json:"change"`
}

func GetItemDetail(itemID int, e error) (itemDetail ItemDetail, err error){

	//Creating URL for request.
	stringWriter := bytes.NewBufferString("http://services.runescape.com/m=itemdb_rs/api/catalogue/detail.json?item=")
	itemIdString := strconv.FormatInt(itemDetail, 10)
	stringWriter.WriteString(itemIdString)

	resp, err := http.Get(stringWriter.String())

	if err != nil {
		return itemDetail, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBytes, itemDetail)

	return itemDetail, err
}


