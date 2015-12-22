/*###########################################
  # Create a simple webserver that responds to 2 requests (String & Struct)
  ###########################################*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

//https://golang.org/pkg/net/http/#Handle
//Will act as one of our routes
type String string

//https://golang.org/pkg/net/http/#Handle
//Will act as the second route or in this case a Handler
type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

//https://golang.org/pkg/net/http/#HandlerFunc.ServeHTTP
//Function that serves the HTTP to our page, This function decides what will be displayed on the specific route.
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

//https://golang.org/pkg/net/http/#HandlerFunc.ServeHTTP
//Function that serves the HTTP to our page, This function decides what will be displayed on the specific route.
func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

	//https://golang.org/pkg/net/http/#ListenAndServe
	//Listen and serve is the function that starts the web server.
	err := http.ListenAndServe("localhost:4000", nil)

	if err != nil {
		//Log errors that occurr.
		log.Fatal("ListenAndServe: ", err)
	}
}
