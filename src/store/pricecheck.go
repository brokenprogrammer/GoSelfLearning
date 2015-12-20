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
	//Load the specified item from the db, if it doesn't exist this function will not work.
	item := db.LoadItem(itemId)

	//If the item doesn't exist we return false
	if item == nil {
		//return out of this function with 0 and false
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
	//Load the specified item from the db, if it doesn't exist this function will not work.
	item := db.LoadItem(itemId)

	//If the item is nil (Doesn't exist)
	if item == nil {
		//return out of this function with a 0 and false.
		return 0, false
	}

	//Set the loaded item's price to the price specified in the parameters.
	item.Price = itemPrice

	//Return the new price and the boolean true for success.
	return item.Price, true
}

/*
	Function that is used to add a new item to the ItemDB. If an item isn't created all the other functions
	will not work since you cannit pricecheck an item that isn't created. Thats why this is used to add an item
	to our storage.
	This function is really similar to the function inside db.go that is getting called, this function will though
	work like a validator or a formatter checking so that the Item can be added or not while the function in db.go will
	work just as expected working with models just placing it in the db.
*/
func AddItem(name string, price float64, popular bool) (int, string, float64, bool) {
	//Creating a new item using the db.CreateItem function, the function returns an id as well as the item created
	id, item := db.CreateItem(name, price, popular)

	//Return the id of the newly created item as well as the data for the item.
	return id, item.Name, item.Price, item.Popular
}
