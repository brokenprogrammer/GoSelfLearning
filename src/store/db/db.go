package db

//Our Item struct is holding the items of our imaginary store.
type Item struct {
	Price float64
}

//Declaring a map that will hold all the items in our store,
//mapping the items with id as key and value a pointer to the Item structure.
var ItemDB = make(map[int]*Item)

/*
	Function that will be used to load items from our ItemDB map.
	If the value doesn't exist in the map it will be initialized with the price 9.5
	The function takes an ID in the parameters that will be used with the ItemDB map as a key to get an Item
*/
func LoadItem(id int) *Item {
	/*
		The old return statement

		return &Item{
			Price: 9.5,
		}*/

	//Checking if there is an Item in the ItemDB with the specified ID, if not the create one.
	if _, exist := ItemDB[id]; exist == false {
		//Set the ItemDB[Key] to hold the memory adress of a new Item.
		ItemDB[id] = &Item{
			Price: 9.5,
		}
	}

	//Return the item from the ItemDB map.
	return ItemDB[id]
}
