package Go_Runescape

//Author: Edwin Heerschap
//Contains functions to get information from the runescape grandexchange api.

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
)

//Category is an array containing information
//on the category. The array is indexed by the alphabet including a hashtag.
//Indexed as # = 0, a = 1, b = 2, c = 3 ...
type Category []categoryLetterItem

//categoryJson is the json object returned by Jagex when
//requesting category information.
type categoryJson struct {
	Types []interface{}        `json:"types"`
	Alpha []categoryLetterItem `json:"alpha"`
}

//categoryLetterItem is the object passed back in the alpha array
//that is passed back in json by the Runescape API.
//Example: {"letter":"#","items":0}
type categoryLetterItem struct {
	Letter string `json:"letter"`
	Items  int64  `json:"items"`
}

//GetItemCountForLetter returns the amount of items found in the Category starting
//with a specific character.
func (c *Category) GetItemCountForLetter(letter byte) (itemAmount int64, err error) {

	//TODO convert # and alphabet to number conversations in seperate function.

	if letter == '#' {

		return (*c)[0].Items, nil
	} else {
		num := int64(letter)

		if num > 64 && num < 91 {
			//Passed letter is a capital letter.
			//Converting number to 1 - 26 matching alphabet i.e 'c'= '3'
			num = num - 63
			return (*c)[num].Items, nil
		} else if num > 96 && num < 122 {
			//Passed letter is a lowercase letter
			//Converting number to 1 - 26 matching alphabet. i.e 'c' = 3
			num = num - 96
			return (*c)[num].Items, nil
		} else {
			err = errors.New("number passed into getItemCountForLetter(letter byte) method is not a letter")
			return -1, err
		}
	}
}

//GetCategory returns a Category for the passed ge_constant.
func GetCategory(geConstant string, HttpClient IHttpClient) (Category, error) {

	cj := categoryJson{}

	num, err := strconv.ParseInt(geConstant, 10, 64)
	//Ensuring passed rune is valid
	if (num < 0 && num > 37) || err != nil {
		return Category{}, errors.New("Go-Runescape: ge_constant passed into GetCategory(ge_constant string) must be between 1->37")
	}

	//Appending string
	stringWriter := bytes.NewBufferString("http://services.runescape.com/m=itemdb_rs/api/catalogue/category.json?category=")
	stringWriter.WriteString(geConstant)

	resp, err := HttpClient.Get(stringWriter.String())

	if err != nil {
		return Category{}, err
	}

	defer resp.Body.Close()

	//Reading bytes
	responseJson, err := ioutil.ReadAll(resp.Body)

	stringWriter.Reset()
	stringWriter.Write(responseJson)

	if err != nil {
		return Category{}, err
	}

	err = json.Unmarshal(responseJson, &cj)

	if err != nil {
		return Category{}, err
	}

	return cj.Alpha, err
}

//ItemJson defines the json object
//returned by jagex when requesting for item information.
type ItemJson struct {
	Item ItemDetail `json:"item"`
}

//ItemDetail contains information on a item.
type ItemDetail struct {
	Icon        string              `json:"icon"`
	IconLarge   string              `json:"icon_large"`
	Id          int64               `json:"id"`
	ItemType    string              `json:"type"`
	TypeIconURL string              `json:"typeIcon"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Current     timeTrendPrice      `json:"current"`
	Today       timeTrendPrice      `json:"today"`
	Day30       timeTrendPercentage `json:"day30"`
	Day90       timeTrendPercentage `json:"day90"`
	Day180      timeTrendPercentage `json:"day180"`
}

//Defines the trend of the price.
type timeTrendPrice struct {
	Trend string `json:"trend"`
	Price string `json:"price"`
}

//Defined the trend of the price as a percentage.
type timeTrendPercentage struct {
	Trend  string `json:"trend"`
	Change string `json:"change"`
}

//Gets grandexchage information from the passed item. If the passed item id is invalid
//a json error will occur and be passed back.
func GetItemDetail(itemID int64, HttpClient IHttpClient) (ItemDetail, error) {

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

//Information on the current price trend.
type TrendPrice struct {
	Trend string      `json:"trend"`
	Price interface{} `json:"price"`
}


//Information about item.
type Items struct {
	Icon        string     `json:"icon"`
	IconLarge   string     `json:"icon_lrage"`
	Id          int        `json:"id"`
	Type        string     `json:"type"`
	TypeIcon    string     `json:"typeIcon"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Current     TrendPrice `json:"current"`
	Today       TrendPrice `json:"today"`
	Members     string     `json:"members"`
}

//Information on an grandexchange category.
type ItemsCatalogue struct {
	Items []Items `json:"items"`
	Total int     `json:"total"`
}

//Gets information on the passed grandexchange category. It is advised that the user uses the ge_constants package to
//choose a grandexchange category. Only items starting with the passed letter will be returned. pageNo is the page number
//of the grandexchange catalogue to return for that category & letter.
func GetItemsCatalogue(geConstant string, letter byte, pageNo int, HttpClient IHttpClient) (c ItemsCatalogue, err error) {

	//Ensuring passed letter is within valid bounds
	num := int64(letter)

	if num > 64 && num < 91 {
		num = num + 32
	} else if num < 97 || num > 122 {
		return c, errors.New("Go-Runescape: GetItemsCatalogue(ge_constant string, letter byte, pageNo int), " +
			"parameter must be between A->Z (not case-sensitive).")
	}

	//Creating url string
	stringWrite := bytes.NewBufferString("http://services.runescape.com/m=itemdb_rs/api/catalogue/items.json?category=")
	stringWrite.WriteString(geConstant)
	stringWrite.WriteString("&alpha=")
	stringWrite.WriteByte(letter)
	stringWrite.WriteString("&page=")
	stringWrite.WriteString(strconv.FormatInt(int64(pageNo), 10))

	resp, err := HttpClient.Get(stringWrite.String())

	if err != nil {
		return c, err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(respBytes, &c)

	if err != nil {
		return c, err
	}

	return c, err

}
