package main

import (
	"fmt"
	"os"
)

type Person struct { //A struct, Like an object or in this language like a class.
	Name string
	Age  int
}

func main() {
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

	os.Exit(1)
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

func (p *Person) boundFunc() {
	fmt.Println("Bound Function")
}
