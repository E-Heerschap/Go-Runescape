package Go_Runescape

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"errors"
	"io/ioutil"
	"strconv"
	"github.com/kingpulse/Go-Runescape"
)

type category struct {
	alpha []categoryLetterItem
}

type categoryJson struct {
	Types []interface{} `json:"types"`
	Alpha []categoryLetterItem `json:"alpha"`
}

//categoryLetterItem is the object passed back in the alpha array
//that is passed back in json by the Runescape API.
//Example: {"letter":"#","items":0}
type categoryLetterItem struct {
	Letter string `json:"letter"`
	Items int64 `json:"items"`
}

//GetItemCountForLetter returns the amount of items found in the category starting
//with a specific character.
func (c *category) GetItemCountForLetter(letter byte) (itemAmount int64, err error){

	if letter == '#' {
		return c.alpha[0].Items, nil
	}else{
		num := int64(letter)

		if num > 64 && num < 91 {
			//Passed letter is a capital letter.
			//Converting number to 1 - 26 matching alphabet i.e 'c'= '3'
			num = num - 63
			return c.alpha[num].Items, nil
		}else if num > 96 && num < 122 {
			//Passed letter is a lowercase letter
			//Converting number to 1 - 26 matching alphabet. i.e 'c' = 3
			num = num - 96
			return c.alpha[num].Items, nil
		}else{
			err = errors.New("Number passed into getItemCountForLetter(letter byte) method is not a letter.")
			return -1, err
		}
	}
}

//GetCategory returns a Category for the passed ge_constant.
func GetCategory(ge_constant string, HttpClient Go_Runescape.HttpClientWrap) (c category, err error){

	cj := categoryJson{}

	num, err := strconv.ParseInt(ge_constant, 10, 64)
	//Ensuring passed rune is valid
	if (num < 0 && num > 37) || err != nil{
		return category{}, errors.New("Go-Runescape: ge_constant passed into GetCategory(ge_constant string) must be between 1->37")
	}

	//Appending string
	stringWriter := bytes.NewBufferString("http://services.runescape.com/m=itemdb_rs/api/catalogue/category.json?category=")
	stringWriter.WriteString(ge_constant)

	resp, err := HttpClient.Get(stringWriter.String())
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Go-Runescape: An error occoured when sending get request.")
		return category{}, err
	}

	//Reading bytes
	responseJson, err := ioutil.ReadAll(resp.Body)

	stringWriter.Reset()
	stringWriter.Write(responseJson)

	if err != nil {
		fmt.Println("Go-Runescape: An error occoured when reading json from Runescape's API")
		return category{}, err
	}

	err = json.Unmarshal(responseJson, &cj)

	if err != nil {
		fmt.Println("Go-Runescape: An error occoured when parsing json from Runescape's API")
		return category{}, err
	}

	c = category{
		alpha: cj.Alpha,
	}

	return c, err
}

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


func GetItemsCatalogue(ge_constant string, letter byte, pageNo int, HttpClient Go_Runescape.HttpClientWrap) (c itemsCatalogue, err error){


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

	resp, err := HttpClient.Get(stringWrite.String())

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
