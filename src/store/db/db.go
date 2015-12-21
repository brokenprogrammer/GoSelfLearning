package db

import (
	"store/models" //Our Item struct is holding the items of our imaginary store.
)

//Declaring a map that will hold all the items in our store,
//mapping the items with id as key and value a pointer to the Item structure.
var ItemDB = make(map[int]*models.Item)

/*
	Function that will be used when returning the entire ItemDB
	This will return the entire map so we can use it.
*/
func GetItems() map[int]*models.Item {
	return ItemDB
}

/*
	Function that will be used when creating new Items.
	This function will accept an item name, price and a boolean if its popular or not.
	Then the item will be created by checking the length of our ItemDB map and the incrementing the value so we get
	a similar system to an MySQL database.
*/
func CreateItem(name string, price float64, popular bool) (int, *models.Item) {
	//Initializing a new Item structure and giving it the values from the parameters
	item := &models.Item{
		Name:    name,    //Name of the item
		Price:   price,   //Price of the item
		Popular: popular, //Popular or not
	}

	//A for loop looping through the entire map to check if there already is an item with the specified name
	for index, _ := range ItemDB {
		if ItemDB[index].Name == name {
			item = nil //Setting the item to nil
		}
	}

	//Adding the newly created item to the map by setting the ID as the maps length + 1 so it acts as
	//an incremented number
	if item != nil {
		ItemDB[len(ItemDB)+1] = item //If the item is not nil (Already exists) we add it
	}

	//Return the id of the item as well as the item itself.
	return len(ItemDB), item
}

/*
	Function that will be used to load items from our ItemDB map.
	If the value doesn't exist in the map it will be initialized with the price 9.5
	The function takes an ID in the parameters that will be used with the ItemDB map as a key to get an Item
*/
func LoadItem(id int) *models.Item {
	/*
		The old return statement

		return &Item{
			Price: 9.5,
		}*/

	//Checking if there is an Item in the ItemDB with the specified ID, if not the create one.
	/*if _, exist := ItemDB[id]; exist == false {
		//Set the ItemDB[Key] to hold the memory adress of a new Item.
		ItemDB[id] = &models.Item{
			Price: 9.5,
		}
	}*/

	//Return the item from the ItemDB map.
	return ItemDB[id]
}
