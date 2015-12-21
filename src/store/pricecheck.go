package store

import (
	"errors" //Golangs error library
	"fmt"
	"store/db" //src/store/db our own created database.
)

/*
	Function to check the price of target item, this function takes in an
	id and returns a float acting as the Item's price and a boolean true if it was
	an success.
*/
func PriceCheck(itemId int) (float64, error) {
	//Load the specified item from the db, if it doesn't exist this function will not work.
	item := db.LoadItem(itemId)

	//If the item doesn't exist we return false
	if item == nil {
		//return out of this function with 0 and false
		return 0, errors.New("Item id does not exist!")
	}

	//If everything goes as planned we return the price and true for success.
	return item.Price, nil
}

/*
	Function to set the price of target item, this function takes in an
	id and a price. This will set a new price to a function with the given id and return
	the new price together with a boolean for success.
*/
func SetPrice(itemId int, itemPrice float64) (float64, error) {
	//Load the specified item from the db, if it doesn't exist this function will not work.
	item := db.LoadItem(itemId)

	//If the item is nil (Doesn't exist)
	if item == nil {
		//return out of this function with a 0 and new error string.
		return 0, errors.New("Item id does not exist!")
	}

	//Set the loaded item's price to the price specified in the parameters.
	item.Price = itemPrice

	//Return the new price and the boolean true for success.
	return item.Price, nil
}

/*
	Function that is used to add a new item to the ItemDB. If an item isn't created all the other functions
	will not work since you cannit pricecheck an item that isn't created. Thats why this is used to add an item
	to our storage.
	This function is really similar to the function inside db.go that is getting called, this function will though
	work like a validator or a formatter checking so that the Item can be added or not while the function in db.go will
	work just as expected working with models just placing it in the db.
*/
func AddItem(name string, price float64, popular bool) (int, string, float64, bool, error) {
	//Creating a new item using the db.CreateItem function, the function returns an id as well as the item created
	id, item := db.CreateItem(name, price, popular)

	if item == nil {
		return 0, "", 0, false, errors.New("Item already exists!")
	}

	//Return the id of the newly created item as well as the data for the item.
	return id, item.Name, item.Price, item.Popular, nil
}

/*
	Function returning all information from target item using the id passed into the function.
	If the item with the given id does not exist we return an error.
*/
func ShowItem(id int) (int, string, float64, bool, error) {
	item := db.LoadItem(id)

	if item == nil {
		return 0, "", 0, false, errors.New("Item id does not exist!")
	}

	return id, item.Name, item.Price, item.Popular, nil
}

/*
	Function that prints out all the items that exists in our db.go's db map.
	This function will print out the entire content for us.
*/
func ShowAllItems() error {
	//Setting the items to the map returned from the GetItems function.
	items := db.GetItems()

	if items == nil {
		return errors.New("Error retrieving Item database.")
	}

	//Looping through the map using range
	for index, _ := range items {
		//Printing out the content of the current index in the items map. Since our map uses integers as keys
		//we can refer to the items using the index in our for loop and then print out the content.
		fmt.Println(index, items[index].Name, items[index].Price, items[index].Popular)
	}

	return nil
}
