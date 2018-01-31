# Go-Runescape Library

[![Build Status](https://travis-ci.org/kingpulse/Go-Runescape.svg?branch=master)](https://travis-ci.org/kingpulse/Go-Runescape)

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

# Documentation

## Get a players highscores/levels
This can be done using the packages:
```
Go-Runescape/highscores
Go-Runescape/highscores/highscore_constants
```

To get a players level, rank or experience you can you need to get a **PlayerHighscores** object and access them from the arrays. This is where the highscore_constants package comes in handy.

Example:
```
//Getting PlayerHighscores object for passed player and type.
//GetPlayerHighscores method takes the players name and the URL path to the endpoint you are after.
//These have been defined as constants in the Go-Runescape/highscores/highscore_constants package for you.
playerHighscore := highscores.GetPlayerHighscores("kingpulse", highscore_constants.RS3PLAYER)

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
This is done using the packages:
```
Go-Runescape/grand_exchange
Go-Runescape/grand_exchange/ge_constants
```

Getting details for an item requires the items id. These can be found [here](http://www.itemdb.biz/).
You can get details through a **ItemDetail** object by doing the following:
```
//Getting a ItemDetail object
tuna, err := grand_exchange.GetItemDetail(359)

if err != nil {
	//Handle
	return
}
//Displaying the description of a tuna.
fmt.Println(tuna.Description)
```
If you want to look at which fields are in the **ItemDetail** object it can be found in:
```
Go-Runescape/grand_exchange/detail.go
```

## Get grandexchange catelog information
This is done using the packages:
```
Go-Runescape/grand_exchange
Go-Runescape/grand_exchange/ge_constants
```
You can get information on the grandexchange catelog page through a **itemsCatelog** object. Before you read the example it is important to note that Jagex and their infinite wisdom have decided to sometimes have price passed as a string or a number which has made it a pain. I've done a simple solution of having the price field as type **interface{}**. This means you will need to use [type assertion](https://golang.org/ref/spec#Type_assertions) to get the type you are after. Generally it is a string. If you have a better solution to this let me know (:
Example:
```
//Getting a itemsCatelog object of based on the first page of high level melee armor starting with the letter b.
catelogue, err := grand_exchange.GetItemsCatalogue(ge_constants.MELEE_ARMOUR_HIGH_LEVEL, 'b', 1)

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
This is done using the packages:
```
Go-Runescape/grand_exchange
Go-Runescape/grand_exchange/ge_constants
```
Lets jump right into the example:
```
//Getting category information for high level melee armor
category, err := grand_exchange.GetCategory(ge_constants.MELEE_ARMOUR_HIGH_LEVEL)

if err != nil {
    //Handle
    return
}

//Getting number of items in category starting with the letter b.
//Unfourtunately this is only useful thing to do with the information Jagex returns.
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

