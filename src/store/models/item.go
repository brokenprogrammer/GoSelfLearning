package models

/*
	Models is a way to organize our structs that will be used over the application.
	For example if we had a account model we could even add functions here that would
	check the account password etc.
*/

type Item struct {
	Name    string
	Price   float64
	Popular bool
}

//Functions with lower case names will not be accessible from other files
/*
	func newItem() *Item {
		return &Item{
			"Bob",
			12,
			true,
		}
	}

	//Functions with upper case names will be accessible from other files
	func NewItem() *Item {
		return &Item{
			"Bob",
			12,
			true,
		}
	}
*/
