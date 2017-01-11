package grand_exchange

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"errors"
	"io/ioutil"
	"strconv"
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
func GetCategory(ge_constant string) (c category, err error){

	cj := categoryJson{}

	num, err := strconv.ParseInt(ge_constant, 10, 64)
	//Ensuring passed rune is valid
	if (num < 0 && num > 37) || err != nil{
		return category{}, errors.New("Go-Runescape: ge_constant passed into GetCategory(ge_constant string) must be between 1->37")
	}

	//Appending string
	stringWriter := bytes.NewBufferString("http://services.runescape.com/m=itemdb_rs/api/catalogue/category.json?category=")
	stringWriter.WriteString(ge_constant)

	resp, err := http.Get(stringWriter.String())
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