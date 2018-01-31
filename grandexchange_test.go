package Go_Runescape

import(
	"testing"
	"github.com/kingpulse/Go-Runescape/ge_constants"
	"net/http"
)

//TestGetItemsCatalogue tests the GetItemsCatalogue function.
func TestGetItemsCatalogue(t *testing.T){

	ci, err := GetItemsCatalogue(ge_constants.POTIONS, 'p', 1)

	if err != nil {
		t.Error("Failed to get items catalogue. Error: " + err.Error())
	}

	if ci.Items == nil {
		t.Error("Failed to get items catalogue. Items array is a nil pointer.")
	}

	if (ci.Total <= 0) {
		t.Error("Failed to get items catalogue. Total count is incorrect.")
	}

	if len(ci.Items) <= 0 {
		t.Error("Failed to get items catalogue. Items array has 0 length")
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
		t.Error("Failed to get category. Error: " + err.Error())
	}

	if newCategory.alpha == nil {
		t.Error("Failed to get category. alpha array is nil.")
	}

	categoryCount, err := newCategory.GetItemCountForLetter('s')

	if err != nil {
		t.Error("Failed to get item count from category. Error: " + err.Error())
	}

	if categoryCount < 1 {
		t.Error("Failed to get correct item count from category.")
	}

	for _, categoryLetter := range newCategory.alpha {

		if categoryLetter.Letter == "" {
			t.Error("Failed to get correct category. Invalid category letter contained in alpha array.")
		}

		if categoryLetter.Items < 0 {
			t.Error("Failed to get correct category. At least one category items is less than 0.")
		}

	}
}