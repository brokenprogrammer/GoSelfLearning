/*###########################################
  # Solve a cipher by rotating letters by 13 spaces
  ###########################################*/
package main

import (
	"io"
	"os"
	"strings"
)

/*
	https://en.wikipedia.org/wiki/ROT13

	Rot13Reader is used to crack codes written by a Rotate by 13 cipher. It rotates letters 13 spaces.

	This was at first very tricky since i didn't understand what GoTour meant by Rot13. I figured it out after
	i googled rot13.

	Example character mapping:
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm

	This all is working with the byte codes or ASCII codes for the letters.
*/

type rot13Reader struct {
	r io.Reader //Implements the io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)       //Initialize n as the rot13Reader structures r io.Reader and call its read function on b
	for i := 0; i < len(b); i++ { //Loop through the given byte slice and crack the code by rotating letters by 13 spaces
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			//If the current letter is between A - N we can add 13 to the value and still be inside the character mappings.
			b[i] += 13
		} else if (b[i] >= 'N' && b[i] < 'Z') || (b[i] >= 'n' && b[i] < 'z') {
			//If the current letter is between N - Z we can decrement 13 to the value and still be inside the character mappings.
			b[i] -= 13
		}
	}

	return n, err
}

func main() {
	//s = new String with the value of a reader with a string inside.
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	//r = a rot13Reader structure and the io.Reader inside the structure is set to the reader we got from the S
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
