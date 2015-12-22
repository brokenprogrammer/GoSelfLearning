/*###########################################
  # Using a for loop to get the Square Root of a value.
  # With added error return value
  ###########################################*/
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64 //A type of float64 which will be used as an error return value

//A method bound to the ErrNegativeSqrt type.
func (e ErrNegativeSqrt) Error() string {
	//Returning an error message
	return fmt.Sprintf("Cannot Sqrt negative number: %g", float64(e))
}

/*
	Function that we use in this exercise to get the Square root of a specified value.
	Sqrt returns the square root of x.
*/
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	//Z is our starting value for testing the x value.
	z := float64(1)

	//D is our delta value, its the difference so we know when our loop ends.
	d := float64(1)

	//The loop will continue while the delta value is higher than 10^-10. Aka there is barely a difference
	for d > 10E-10 {
		//z0 will be used to compare new and old z values.
		z0 := z

		//Using Newtons method used to approximate the square root of a value.
		//This algorithm was provided in the exercise.
		z = z - (z*z-x)/(2*z)

		//Delta is the difference between our old z value and the new we got from
		//Newtons method above.
		d = z - z0

		//Delta can become both a negative value and a positive, here we make the value
		//Positive if it is negative already so the loop can continue.
		if d < 0 {
			//	d = -d
			d = math.Abs(d) //There is two ways of making the delta value positive, the math function or the above
		}
	}

	//Returns Square root of x
	return z, nil
}

func main() {
	//Comparing our own Sqrt function and testing error messages
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(3))
	fmt.Println(Sqrt(-3))
	fmt.Println(Sqrt(5))
	fmt.Println(Sqrt(-5))
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-9))
}
