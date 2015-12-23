//###########################################
//# Understanding error handling in Go
//# http://blog.golang.org/error-handling-and-go
//###########################################
package main

import (
	"fmt"
)

/*
	When reading more about how the Errors work here i think i finally got it.

	The error interface requires only a Error method; specific error implementations might have additional methods.

	What i believe this means is that the error interface which is returned in TestError function just require and
	type which has a Error() method. That is why when we  call ShortStringError(string) it automatically knows that
	we are using it as an error thats why the Error() method for ShortStringError is called.

	Stringers should be working in a similar way, This is what the fmt package documentation says about Stringers:

	Stringer is implemented by any value that has a String method, which defines the “native” format for that value.
	The String method is used to print values passed as an operand to any format that accepts a
	string or to an unformatted printer such as Print.
*/
type ShortStringError string

//The error method which returns our error message.
func (s ShortStringError) Error() string {
	return fmt.Sprintf("The string is too short: %s", string(s))
}

//Function that returns an error if the input string is less than 5 characters long
func TestError(short string) (int, error) {
	if len(short) < 5 {
		//Return a ShortStringError
		return 0, ShortStringError(short)
	}

	return 1, nil
}

func main() {
	fmt.Println("In main")

	v, err := TestError("Shor")

	if err != nil {
		//Print out the error message
		fmt.Println(err)
	}

	fmt.Println(v)
}
