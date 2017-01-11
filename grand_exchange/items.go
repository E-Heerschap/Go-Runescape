package grand_exchange

import (
	"errors"
	"net/http"
	"bytes"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type trendPrice struct {
	Trend string `json:"trend"`
	Price interface{} `json:"price"`
}

type items struct {
	Icon string `json:"icon"`
	Icon_large string `json:"icon_lrage"`
	Id int `json:"id"`
	Type string `json:"type"`
	TypeIcon string `json:"typeIcon"`
	Name string `json:"name"`
	Description string `json:"description"`
	Current trendPrice `json:"current"`
	Today trendPrice `json:"today"`
	Members string `json:"members"`
}

type itemsCatalogue struct {
	Items []items `json:"items"`
	Total int `json:"total"`
}


func GetItemsCatalogue(ge_constant string, letter byte, pageNo int) (c itemsCatalogue, err error){


	//Ensuring passed letter is within valid bounds
	num := int64(letter)

	if num > 64 && num < 91 {
		num = num + 32
	}else if num < 97 || num > 122 {
		return c, errors.New("Go-Runescape: GetItemsCatalogue(ge_constant string, letter byte, pageNo int), " +
			"parameter must be between A->Z (not case-sensitive).")
	}

	//Creating url string
	stringWrite := bytes.NewBufferString("http://services.runescape.com/m=itemdb_rs/api/catalogue/items.json?category=")
	stringWrite.WriteString(ge_constant)
	stringWrite.WriteString("&alpha=")
	stringWrite.WriteByte(letter)
	stringWrite.WriteString("&page=")
	stringWrite.WriteString(strconv.FormatInt(int64(pageNo), 10))

	resp, err := http.Get(stringWrite.String())

	if err != nil {
		return c, err;
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBytes, &c)

	if err != nil {
		fmt.Println("Go-Runescape: Failed to unmarshal JSON while getting item catalouge.")
		return c, err
	}

	return c, err

}
