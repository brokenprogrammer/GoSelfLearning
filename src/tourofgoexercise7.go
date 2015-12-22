/*###########################################
  # Creating our own reader that always prints out 'A'
  ###########################################*/
package main

import "golang.org/x/tour/reader" //https://github.com/golang/tour/blob/master/reader/validate.go#L13
import "fmt"                      //Used for the printing of an error message.

//Structure for our reader.
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
/*
	This exercise is where i started to have to think a bit but it was since i didn't quite understand
	what to do after i got the information i did in the tour of go. The information for the task was:

	Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

	I now understand what they meant and this is how it should be impletented.

	This function is bound to MyReader structure and it takes in a slice of bytes that will be checked.
	We always want the bytes to read 'A' so i created a for loop to assign the whole slice to just have 'A'.
	Then we return the length and an error message, if no error occurred then we just return nil.
*/
func (m MyReader) Read(b []byte) (int, error) {
	if len(b) == 0 { //If there is no data in our byte slice
		return 0, fmt.Errorf("Buffer is not long enough") //return an error message
	}

	//Looping through the byte slice and assigning values to 'A'
	for i, _ := range b {
		b[i] = 'A'
	}

	//Return the length and error.
	return len(b), nil
}

func main() {
	//This validation function is putting a 1024*2048 sized slice into our Read function above and then loops
	//through it to check if the values are actually 'A'
	reader.Validate(MyReader{})
}
