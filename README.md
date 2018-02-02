# Go-Runescape Library

[![Build Status](https://travis-ci.org/kingpulse/Go-Runescape.svg?branch=master)](https://travis-ci.org/kingpulse/Go-Runescape) [![Maintainability](https://api.codeclimate.com/v1/badges/df6f8708f7170d5e2019/maintainability)](https://codeclimate.com/github/kingpulse/Go-Runescape/maintainability) [![codecov](https://codecov.io/gh/kingpulse/Go-Runescape/branch/master/graph/badge.svg)](https://codecov.io/gh/kingpulse/Go-Runescape) TODO Add GODOC Badge

## Overview
Go-Runescape is a library for the Runescape API for Go. The library supports:

- Getting player Levels
- Getting a players Rank
- Getting item details from the grandexchange (RS3 Only)
- Getting items in categories from the grandexchange (RS3 Only)

If you are after information regarding the old-school grandexchange, read the note at the end.

This project will no longer be updated because the initial goals for the project have been met.

I apollogise if this does not meet the *effective go* standards. This was primarly built while I was initially learning go (: .

----

# Install

```
$ go get github.com/kingpulse/Go-Runescape
```

# Documentation

## IHttpClient

The IHttpClient is a interface which is used in all of the major functions. The purpose IHttpClient is to add flexibility
and testability to the code by allowing the user to define their own Get() method. This allows the user to use proxies, mock returns, etc
as long as the struct has a defined Get() method following:

```
func (t type) Get(url string) (*http.Response, error)
```

Here is an example of the user using the default http.Client struct to send requests for a players highscores

```
//Creating default http.Client
defaultClient := htttp.DefaultClient

//Sending request for player highscores.
hs, err := GetPlayerHighscores("le me", highscore_constants.RS3PLAYER, defaultClient)
```

In the rest of the examples the creation of a struct that inherits the IHttpClient interface will be neglected and
HttpClient will be used to fill the parameter where required.

## Constants

This package contains two sub-packages dedicated to constants. The constants are stored in sub-packages so the user can easily
identify which group they belong to without each constant being prefixed individually.

The two constants packages have constants for information regarding the grandexchange and highscores, they are stored in
ge_constants and highscore_constants respectively.

## Get a players highscores/levels

To get a players level, rank or experience you can you need to get a **PlayerHighscores** object and access them from the arrays. This is where the highscore_constants package comes in handy.

Example:
```
//Getting PlayerHighscores object for passed player and type.
//GetPlayerHighscores method takes the players name and the URL path to the endpoint you are after.
//These have been defined as constants in the Go-Runescape/highscores/highscore_constants package for you.
playerHighscore, err := highscores.GetPlayerHighscores("kingpulse", highscore_constants.RS3PLAYER, HttpClient)

//Getting mining level of the RS3 player "kingpulse" using the highscores_constants package.
miningLevel = playerHighscore.Levels[highscore_constants.MINING]

//Printing the mining level
fmt.Print(strconv.FormatInt(miningLevel, 10))

```

The same can be done for the players rank or experience in the specified skill. For example:
```
miningExperience := playerHighscore.XP[highscore_constants.MINING]
miningRank := playerHighscore.Ranks[highscore_constants.MINING]
```

## Get grandexchange item details

Getting details for an item requires the items id. These can be found [here](http://www.itemdb.biz/).
You can get details through a **ItemDetail** object by doing the following:
```
//Getting a ItemDetail object
tuna, err := grand_exchange.GetItemDetail(359, HttpClient)

if err != nil {
	//Handle
	return
}
//Displaying the description of a tuna.
fmt.Println(tuna.Description)
```

## Get grandexchange catelog information

You can get information on the grandexchange catelog page through a **itemsCatelog** object. Before you read the example it is important to note that Jagex and their infinite wisdom have decided to sometimes have price passed as a string or a number which has made it a pain. I've done a simple solution of having the price field as type **interface{}**. This means you will need to use [type assertion](https://golang.org/ref/spec#Type_assertions) to get the type you are after. Generally it is a string. If you have a better solution to this let me know (:
Example:
```
//Getting a itemsCatelog object of based on the first page of high level melee armor starting with the letter b.
catelogue, err := grand_exchange.GetItemsCatalogue(ge_constants.MELEE_ARMOUR_HIGH_LEVEL, 'b', 1, HttpClient)

if err != nil {
	//Handle
	return
}

//Printing the name of the first item on the page    
fmt.Println(catelogue.Items[0].Name)

//Printing todays price of the item.
fmt.Println(catelogue.Items[0].Today.Price.(string))
```

## Get number of items in category for letter

Lets jump right into the example:

```
//Getting category information for high level melee armor
category, err := grand_exchange.GetCategory(ge_constants.MELEE_ARMOUR_HIGH_LEVEL, HttpClient)

if err != nil {
    //Handle
    return
}

//Getting number of items in category starting with the letter b.
itemCount, err := category.GetItemCountForLetter('b')

if err != nil {
	//Handle
	return
}

//Printing out number of items.
fmt.Println(strconv.FormatInt(itemCount, 10))
```

### Note
Keep in mind Jagex returns very minimal information unfourtunately so not much information is availible from their API.
Currently this library does not have any support for the OSBuddy API. If you are looking for a GOOD grand exchange API, use OSBuddy's. It provides much more information than Jagex's.
