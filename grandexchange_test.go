package Go_Runescape

import (
	"testing"
	"github.com/kingpulse/Go-Runescape/ge_constants"
	"net/http"
)

//TestGetItemsCatalogue tests the GetItemsCatalogue function.
func TestGetItemsCatalogue(t *testing.T) {

	//Creating working http client.
	httpClient := &http.Client{}

	ci, err := GetItemsCatalogue(ge_constants.POTIONS, 'p', 1, httpClient)

	if err != nil {
		t.Error("Failed to get items catalogue. Error: " + err.Error())
	}

	if ci.Items == nil {
		t.Error("Failed to get items catalogue. Items array is a nil pointer.")
	}

	if ci.Total <= 0 {
		t.Error("Failed to get items catalogue. Total count is incorrect.")
	}

	if len(ci.Items) <= 0 {
		t.Error("Failed to get items catalogue. Items array has 0 length")
	}

	//Creating failure http client (will intentionally return an error)
	failureClient := notNilHttpClient{}

	_, err = GetItemsCatalogue(ge_constants.POTIONS, 'p', 1, failureClient)

	if err == nil {
		t.Error("GetItemsCatalogue failed to recognise errors from GET request.")
	}

	ijClient := invalidJsonHttpClient{}

	_, err = GetItemsCatalogue(ge_constants.POTIONS, 'p', 1, ijClient)

	if err == nil {
		t.Error("GetItemsCatalogue failed to recognise invalid json.")
	}

}

//TestGetITemDetail tests the GetItemDetail function.
func TestGetItemDetail(t *testing.T) {

	testClient := &http.Client{}

	//Getting Rune Scimitar item detail.
	item, err := GetItemDetail(1333, testClient)

	if err != nil {
		t.Error("Failed to get item detail. Error: " + err.Error())
	}

	if item.Name != "Rune scimitar" {
		t.Error("Failed to get item detail. Name is incorrect.")
	}

	if item.Id != 1333 {
		t.Error("Failed to get item detail. Item id is incorrect.")
	}

	//Checking for valid TimeTrendPrice objects.
	checkTimeTrendPrice(t, &item.Current)
	checkTimeTrendPrice(t, &item.Today)

	//Checking for valid TimeTrendPercentage objects.
	checkTimeTrendPercentage(t, &item.Day30)
	checkTimeTrendPercentage(t, &item.Day90)
	checkTimeTrendPercentage(t, &item.Day180)

	failureClient := notNilHttpClient{}

	_, err = GetItemDetail(1333, failureClient)

	if err == nil {
		t.Error("GetItemDetail failed to recognize error from GET request.")
	}

	ijClient := invalidJsonHttpClient{}

	_, err = GetItemDetail(1333, ijClient)

	if err == nil {
		t.Error("GetItemDetail failed to recognize invaid JSON")
	}
}

//checkTimeTrendPrice checks if string fields in timeTrendPrice objects are not equal to ""
func checkTimeTrendPrice(t *testing.T, tdp *timeTrendPrice) {

	if tdp.Price == "" {
		t.Error("Invalid Price for timeTrendPrice")
	}

	if tdp.Trend == "" {
		t.Error("Invalid Trend for timeTrendPrice")
	}

}

//checkTimeTrendPercentage checks if string fields in timeTrendPercentage objects are not equal to ""
func checkTimeTrendPercentage(t *testing.T, ttp *timeTrendPercentage) {

	if ttp.Trend == "" {
		t.Error("Invalid Trend for timeTrendPercentage")
	}

	if ttp.Change == "" {
		t.Error("Invalid Change for timeTrendPercentage")
	}

}

func TestGetCategory(t *testing.T) {

	//Getting default http client.
	testClient := &http.Client{}

	newCategory, err := GetCategory(ge_constants.MELEE_WEAPONS_HIGH_LEVEL, testClient)

	if err != nil {
		t.Error("Failed to get Category. Error: " + err.Error())
	}

	if newCategory.alpha == nil {
		t.Error("Failed to get Category. alpha array is nil.")
	}

	categoryCount, err := newCategory.GetItemCountForLetter('s')

	if err != nil {
		t.Error("Failed to get item count from Category. Error: " + err.Error())
	}

	if categoryCount < 1 {
		t.Error("Failed to get correct item count from Category.")
	}

	for _, categoryLetter := range newCategory.alpha {

		if categoryLetter.Letter == "" {
			t.Error("Failed to get correct Category. Invalid Category letter contained in alpha array.")
		}

		if categoryLetter.Items < 0 {
			t.Error("Failed to get correct Category. At least one Category items is less than 0.")
		}

	}

	failureClient := notNilHttpClient{}

	_, err = GetCategory(ge_constants.MELEE_WEAPONS_HIGH_LEVEL, failureClient)

	if err == nil {
		t.Error("GetCategory failed to recognise errors from a GET request.")
	}

	ijClient := invalidJsonHttpClient{}

	_, err = GetCategory(ge_constants.MELEE_WEAPONS_HIGH_LEVEL, ijClient)

	if err == nil {
		t.Error("GetCategory failed to recognise errors from invalid json.")
	}
}
