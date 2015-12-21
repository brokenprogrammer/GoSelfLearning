package main

import (
	"fmt"
	"math/rand"
	"sort"
	"store" //src/store package within this directory
)

type Person struct { //A struct, Like an object or in this language like a class.
	Name string
	Age  int
}

type Dog struct { //A struct, Like an object or in this language like a class.
	Name   string
	Age    int
	Race   string
	Happy  bool
	Father *Dog
}

type SuperPerson struct { //A struct which is a composition, it can use functions of the Person
	*Person
	Pressure int
}

func main() {

	//#############################################################
	//# Day 1, Variables, Printing, Functions, Structs
	//#############################################################

	var power int
	power = 9000

	var super int = 8000

	duper := 7000

	sduper := getValue()

	//name, sound := "John", "Woof!" //Error since declared and not used.

	fmt.Println("Power: ", power)
	fmt.Println("Super: ", super)
	fmt.Println("duper: ", duper)
	fmt.Println("Sduper: ", sduper)

	//fmt.Println("It's over", os.Args[1])

	newPerson := Person{ //new Person struct
		"Josh",
		16,
	}
	//newPerson2 := Person{ //New Person Struct
	//	Name: "Larry",
	//	Age:  22,
	//}

	//newPerson3 := Person{} // New person struct, Different ways of initiating a new struct.
	//newPerson3.Name = "Ann"
	//newPerson3.Age = 44

	pointerPerson := &Person{"Broken", 19} //A new pointer variable

	notChangeStruct(newPerson)      //Doesnt change the actual values of "newPerson"
	doesChangeStruct(pointerPerson) //Does change the actual value of "pointerPerson"

	fmt.Println(newPerson.Age)
	fmt.Println(pointerPerson.Age)

	pointerPerson.boundFunc() //Function bound to the structure
	newPerson.boundFunc()     //This works with the non pointer varables as well.

	//#############################################################
	//# Day 2 Structs, Pointers, Constructors, Composition, Overloading
	//#############################################################

	myDog := Dog{ //Initialize new dog struct
		Name: "Benny",
		Age:  2,
		Race: "Bulldog",
	}

	myDog.setDogHappy(true)     //Setting the dogs value to true using a bound function
	notBoundDogFunction(&myDog) //You can send in a pointer value of your struct, You do not have to initialize the struct as a pointer.

	if myDog.Happy == true { //Checking if the dog is Happy.
		fmt.Println("myDog is Happy!")
		fmt.Println("myDog is: ", myDog.Age) //prints the new age we gave the dog through notBoundDogFunction()
	}

	myOtherDog := new(Dog) //This is the same as &Dog
	myOtherDog.Name = "OtherDog"

	myThirdDog := NewDog("Angel", 5, "PitBull", false) //Using a factory function as constructor.

	//Now i have added a Father to the dog struct of the type *Dog
	newestDog := &Dog{
		Name:  "Lol",
		Age:   5,
		Race:  "One",
		Happy: true,
		Father: &Dog{ //The dogs father an adress to LolSr
			Name:   "LolSr",
			Age:    10,
			Race:   "First",
			Happy:  false,
			Father: nil,
		},
	}

	fmt.Println(myThirdDog.Name) //Just printing out values from the variables so go can compile and run
	fmt.Println(newestDog.Father)

	strangeGuy := &SuperPerson{ // A new superperson structure that has a pointer to a person within it
		Person:   &Person{"IronMan", 10}, // Adress to the person IronMan
		Pressure: 50,
	}

	strangeGuy.Introduce() //The SuperPerson structure can now use functions that were meant for the Person struct,
	//Here we changed the SuperPerson structure to also have an Introduce function and it has now overwritten the Persons
	//Introduce but our strangeGuy variable can still access its old Introduce function like:
	strangeGuy.Person.Introduce() //Calling the Person structures introduce function
	strangeGuy.boundFunc()        //Also a function bound to the Person struct

	//The name field is bound to the Person struct so consider both of theese valid.
	fmt.Println(strangeGuy.Name)        //Prints the name of the strange guy
	fmt.Println(strangeGuy.Person.Name) //Prints the name of the strange guy

	//#############################################################
	//# Day 4 Arrays, Slices, Copy, Make(TYPE, len, cap), Maps
	//#############################################################

	//Creating a new array
	//var myArray [10]int

	//Creating a new slice
	//mySlice := []int{1, 5, 10}

	//Using make instead of new since new is allocating memory.
	//Make(type, length, capacity)
	mySecondSlice := make([]int, 1, 10)
	//The length is made so that we cannot access all of it through varName[x] we have to use append to add values to it.
	mySecondSlice[0] = 120                     //Works since the length is 1
	mySecondSlice = append(mySecondSlice, 510) //After that we have to use append
	fmt.Println("My Second Slice: ", mySecondSlice)

	//Luckily for us we can reslice our slice if we would like
	mySecondSlice = mySecondSlice[2:5]
	mySecondSlice[1] = 240

	fmt.Println("My Second Re-Sliced: ", mySecondSlice)

	//Append is though so special that we do not really have to do it like this
	//When appending to a full slice append will create a new larger slice and copy the old values into it.
	//Exactly how dynamic arrays work in Python, PHP and Ruby.

	//Copy is our way of copying a part of a slice to another, this is something that is easier in go than other languages
	scores := make([]int, 100) //length of 100

	for i := 0; i < len(scores); i++ {
		scores[i] = int(rand.Int31n(100)) //Giving all of scores random values
	}

	sort.Ints(scores) //Sorting the values in scores

	//Worst scores will hold the 5 worst scores from the scores slice
	worstScores := make([]int, 5) //has the length of 5

	//Copy to worstScores, the scores from start to 5.
	copy(worstScores[:], scores[:5]) //If i try to copy more than 5 values they just wont get copied to worstScores
	//Copying less values wont fill the entire worstScores slice and you can specify where in worstScores you want
	//The values to be copied to.
	fmt.Println("Worst scores: ", worstScores)

	//Creating a new map, maps grow dynamically but we can just like with the slices specify a length if we want.
	//If you know how many positions you will need it can be a good idea to specify for performance.
	myMap := make(map[string]int)
	myMap["Paris"] = 5001

	/*myMap["Warsaw"] will return 2 values, the int value bound to the specified key and a boolean (Exists)
	This here is a great example of how Go's multiple return values work, If i would like to not care of a value
	Being returned i would use the _, like the book previously mentioned.*/
	myMapCity, myMapDistance := myMap["Paris"]
	myMapSize := len(myMap)

	fmt.Println("My Maps:")
	fmt.Println("Paris: ", myMapCity, myMapDistance) //Prints out the value bound to Paris and if it exists or not
	fmt.Println("MapSize: ", myMapSize)

	//This is how you would delete a value from the map
	delete(myMap, "Paris")

	fmt.Println("New MapSize: ", len(myMap))

	//Another way to initialize maps
	newMap := map[string]int{
		"England": 1024,
		"Germany": 502,
		"Russia":  5000,
	}

	fmt.Println("Itterating over a new map: ")

	//We can itterate over the map using a for loop with range but it will return it in a random order:
	for key, value := range newMap {
		fmt.Println(key, value)
	}

	/*
		Theese functions does not work anymore since the code was changed. Now you first have to add the product
		for it to exist, trying to check the price of item with id 1 before it existed will ressult in a false message.

		Now you have to do store.AddItem(name, price, popular) and then the item will be added with an incremented id.

			fmt.Println("---USING THE IMPORTED STORE PACKAGE---")
			fmt.Println(store.PriceCheck(1))
			fmt.Println(store.PriceCheck(2))
			fmt.Println(store.PriceCheck(3))
			fmt.Println("---CHANGING PRICES---")
			store.SetPrice(1, 12.5)
			store.SetPrice(2, 150)
			store.SetPrice(3, 12345)
			fmt.Println(store.PriceCheck(1))
			fmt.Println(store.PriceCheck(2))
			fmt.Println(store.PriceCheck(3))
	*/

	//#############################################################
	//# Day 5 Worked with packages, Error handling and Defer
	//#############################################################
	fmt.Println("---TRYING NEW FUNCTIONS---")
	fmt.Println(store.AddItem("Shoe", 500, true))
	fmt.Println(store.AddItem("Cat", 10.4, false))
	fmt.Println(store.AddItem("Hat", 1025, true))
	fmt.Println(store.PriceCheck(3))
	store.SetPrice(3, 1.5)
	fmt.Println(store.PriceCheck(3))

	fmt.Println("---TESTING ERRORS---")

	fmt.Println(store.SetPrice(15, 120)) //Should report an erro message since item id 15 doesnt exist

	_, err := store.SetPrice(6, 55) //Setting the price and if there is an error we print it out.
	if err != nil {
		fmt.Println(err) //If there is an error we print it out.
	}

	_, err = store.PriceCheck(15) //Checking the price of an item that is not created.
	if err != nil {
		fmt.Println(err) //If there is an error we print it out.
	}

	fmt.Println("---TESTING VIEWING FUNCTIONS---")
	fmt.Println(store.ShowItem(1))
	fmt.Println(store.ShowItem(2))
	fmt.Println(store.ShowItem(3))

	fmt.Println("---SHOWING ALL ENTRIES---")
	store.ShowAllItems()
	fmt.Println(store.AddItem("Box", 10.5, false))
	fmt.Println(store.AddItem("Dog", 5, true))
	fmt.Println(store.AddItem("Hat", 50, false)) //Returns an error since there already is an Hat item.

	fmt.Println("---LEARNING ABOUT DEFER")

	//Defer keyword will run a function or statement when the function it is used in returns.
	defer deferFunction()

	//Initialized if statements, initialize a value and use it inside the statement:
	if x := 10; 5 < x {
		fmt.Println("x is larger than 5")
	}

	/*
		That might be a silly example but more realistically you can do something like:
		if err := process(); err != nil{
			return err
		}
		Theese values are not available outside the if statement, but they are available in the
		else if and else statements.
	*/

	//Go book p 40...
}

func getValue() int { //Simple function that returns an integer
	return 123
}

//Prints message
func log(message string) { //Prints the message entered
	fmt.Println(message)
}

func add(a int, b int) int { //returns the sum of 2 values
	return a + b
}

//Multiple return values, return values you dont care about can be called using  _:=
func testing(name string) (int, bool) {
	return 1, false //Returns multiple values
}

//Doesnt change the incoming person because it creates a new local person p
func notChangeStruct(p Person) {
	p.Age += 50
}

//Takes in a pointer to a Person
func doesChangeStruct(p *Person) {
	p.Age += 50
}

func (p *Person) boundFunc() { // A function that is bound to the Person Struct
	fmt.Println("Bound Function")
}

func (d *Dog) setDogHappy(isHappy bool) { //A function that is bound to the Dog Struct
	d.Happy = isHappy
}

func notBoundDogFunction(d *Dog) {
	d.Age += 1
}

func NewDog(name string, age int, race string, happy bool) *Dog { //Constructor (Factory function) The type of the function is like the Dog struct
	//Function returns *Dog just as the type of the function.
	return &Dog{ //This kind of function doesnt have to return a pointer, it can just return a normal version as well.
		Name:  name,
		Age:   age,
		Race:  race,
		Happy: happy,
	}
}

//Function bound to the Person structure, SuperPerson can also use this function
func (p *Person) Introduce() {
	fmt.Println("Hello, ", p.Name)
}

//Our overwritten Introduce function which overwrites the Person structures introduce function, Note that the old is still accessable.
func (s *SuperPerson) Introduce() {
	fmt.Println("Hello Super, ", s.Name)
}

//When using the defer keyword we can make a function or statement run when the function returns.
//This function is defered by the main function. So this will get printed when the main function ends.
func deferFunction() {
	fmt.Println("Inside the 'deferFunction'!, The main function is ending.")
}
