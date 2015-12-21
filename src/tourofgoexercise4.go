/*###########################################
  # Using a function to return a function which returns fibonacci numbers
  ###########################################*/
package main

import "fmt"

/*
	https://en.wikipedia.org/wiki/Fibonacci_number

	Fibonacci numbers are: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
	Fn = F1 + F2

	What we want to happend is values start with 1 + 1 and then take the previous used value in the new addition.
	Example sequence: 1+1, 2+1, 3+2, 5+3, 8+5, 13+8
*/

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum := 0            //Initializing the sum variable that will hold the return value
	oldsum := 1         //Initializing the old sum varaible which will hold the previous sum used in the fibonacci algorithm
	return func() int { //Returning a function
		temp := sum        //Initializing a temporary variable so we can set our oldsum to the value after addition has been done.
		sum = sum + oldsum //Make the addition (Fn = F1 + F2)
		oldsum = temp      //Setting the old sum to the old sum
		return sum         //Return the new fibonacci number
	}
}

/*
	Our Fibonacci function has 2 variables sum and oldsum, theese are initialized when we initialize our f variable
	in main and they are remembered since our f variable is remembered so they will get stored when the function call
	has been made. That is why we successfully can call the function several times and get different values out of it.
*/

func main() {
	f := fibonacci()           //We are initializing the variable f to hold a function, so the content of f is the call to fibonacci()
	for i := 0; i <= 11; i++ { //Creating a loop that will loop x times
		fmt.Println(f()) //Calling our f x times.
	}
}
