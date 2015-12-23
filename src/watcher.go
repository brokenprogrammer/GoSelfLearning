//#######################################
//# An IRC Watcher that i created.
//# Some help taken from theese links:
//# http://archive.oreilly.com/pub/h/1963
//# https://golang.org/pkg/net/#Conn
//#######################################
package main

import (
	"bufio" //https://golang.org/pkg/bufio/
	"fmt"
	"log"
	"net"
	"strings" //https://golang.org/pkg/strings
)

type Watcher struct {
	server  string
	port    string
	nick    string
	channel string
	conn    net.Conn
}

//A factory function for our Watcher structure
func NewBot() *Watcher {
	return &Watcher{
		server:  "irc.freenode.net", //Server our IRC uses
		port:    "6667",             //Port for IRC
		nick:    "BrokenBot",        //Name of our bot
		channel: "#gobotter",        //Channel bot is connecting to
		conn:    nil,                //The conn will get initialized from the Connect function
	}
}

//This function starts a new dial connection using the information our Watcher structure has in it.
func (w *Watcher) Connect() (net.Conn, error) {
	//Start a new dial connection to the server.
	conn, err := net.Dial("tcp", w.server+":"+w.port)
	if err != nil {
		//If an error occurr we return the error message
		log.Fatal("Unable to connect to irc: ", err)
	}
	//Set the Watcher structures conn to the newly created connection
	w.conn = conn
	//Print out the success connection message
	log.Printf("Connected to %s (%s)", w.server, w.conn.RemoteAddr())
	//Return the connection and no error.
	return w.conn, nil
}

func main() {
	//Create a new bot using factory method
	bot := NewBot()

	//Open up new connection
	conn, _ := bot.Connect()
	//Close connection when main closes
	defer conn.Close()

	//Write to the IRC server
	conn.Write([]byte("NICK " + bot.nick + "\r\n"))                         //IRC server requests a nickname for the user
	conn.Write([]byte("USER " + bot.nick + " 8 *  : " + bot.nick + "\r\n")) //IRC server always requests a realname in this format
	conn.Write([]byte("JOIN " + bot.channel + "\r\n"))                      //Using the irc JOIN command to join the channel our bot uses.
	conn.Write([]byte("PRIVMSG " + bot.channel + " :Hello World!\r\n"))

	//Using a Go Routine to handle a Controll Panel for the bot simmultaniously as the bot is running
	go ControllPanel(conn, bot)

	//The bufio reader will read data we get from our connection and return it as a string.
	connBuff := bufio.NewReader(conn)

	for {
		str, err := connBuff.ReadString('\n')
		if len(str) > 0 { //If there is a message from the server
			fmt.Println(str) //Print it out

			//Staying connected to the IRC server
			splitted := strings.Split(str, " ") //Split the string into a slice
			if splitted[0] == "PING" {          //If the IRC Server is pinging us
				fmt.Println(splitted)
				conn.Write([]byte("PONG " + splitted[1] + "\r\n"))                            //Respond back with a PONG
				conn.Write([]byte("PRIVMSG " + bot.channel + " :Hello I'm Still here! \r\n")) //Tell the chat you're still here
			}
		}
		if err != nil {
			break
		}
	}
}

//Experimenting with a ControllPanel for the bot.
func ControllPanel(conn net.Conn, bot *Watcher) {
	for {
		fmt.Println("Waiting for input: ")

		var input string

		fmt.Scan(&input)

		switch input {
		case "Hello":
			conn.Write([]byte("PRIVMSG " + bot.channel + " :Hello I'm At The ControllPanel! \r\n"))
		case "Quit":
			conn.Write([]byte("QUIT " + "\r\n"))
		}
	}
}
