/*###########################################
  # Adding a stringer to a type
  ###########################################*/
package main

import "fmt"

//The type we will use
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {

	//Returning the content in the form of a string, i was trying to find a way to loop through the 4 values but found
	//nothing.
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3]) //Doesn't look like the optimal way
}

func main() {
	//Initializing a new map using our type with a stringer.
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a) //This will print out a string using our stringer function when a is printed.
	}
}
