package store

import (
	"store/db" //src/store/db our own created database.
)

/*
	Function to check the price of target item, this function takes in an
	id and returns a float acting as the Item's price and a boolean true if it was
	an success.
*/
func PriceCheck(itemId int) (float64, bool) {
	//Load the specified item from the db, if it doesn't exist it will be created.
	item := db.LoadItem(itemId)

	//If the item doesn't exist we return false
	if item == nil {
		return 0, false
	}

	//If everything goes as planned we return the price and true for success.
	return item.Price, true
}

/*
	Function to set the price of target item, this function takes in an
	id and a price. This will set a new price to a function with the given id and return
	the new price together with a boolean for success.
*/
func SetPrice(itemId int, itemPrice float64) (float64, bool) {
	//Load the specified item from the db, if it doesn't exist it will be created.
	item := db.LoadItem(itemId)

	//Set the loaded item's price to the price specified in the parameters.
	item.Price = itemPrice

	//Return the new price and the boolean true for success.
	return item.Price, true
}
