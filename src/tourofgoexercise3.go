/*###########################################
  # Using maps to count words in a string
  ###########################################*/
package main

//We get to use the strings library to return words from a string as well as a library bound to the
//tour of go which is used to test results of our function.
import (
	"golang.org/x/tour/wc" //https://github.com/golang/tour/blob/master/wc/wc.go#L10
	"strings"              //https://golang.org/pkg/strings/#Fields
)

/*
	Function that we will use in this exercise to count the words in a string
	and map the words into a map as keys with the value of how many times they appear
	in the string.

	WordCount returns a map with keys as words and the values as the times they appear in the string.
*/
func WordCount(s string) map[string]int {
	//Declaring the map we will use, we do not know how many words will be passed in so we cannot give it a specified length.
	ourMap := make(map[string]int)

	//A slice to hold all the words, here same as above we cannot give it a specified length since we do not know how many words there is.
	var theWords []string

	//Using the strings library together with the function Fields to return a slice with all the words in it.
	theWords = strings.Fields(s)

	//We can then use range to itterate over the whole slice of words, Range returns 2 values (Index and Value)
	//Here we only need the value since we are working with the keys of the map. You can tell Go that you do not want
	//one of the returned values by assigning it the "_" name like in this for loop.
	for _, val := range theWords {
		//Incrementing the value for the key val. If the key doesn't exist it will create it.
		ourMap[val] += 1
	}

	//Return the map together with our words counted.
	return ourMap
}

func main() {
	//wc.Test is a test suite made for Tour of Go which will test our WordCount function.
	wc.Test(WordCount)
}
