/*###########################################
  # Using a for loop combined with slices to draw an image.
  ###########################################*/
package main

//Source code of the imported package this code uses.
//https://github.com/golang/tour/blob/master/pic/pic.go#L15
import "golang.org/x/tour/pic"

/*
	Function that we use in this exercise to generate a 2 dimensional Slice.
	Pic returns a 2 dimensional Slice of the type uint8.
*/
func Pic(dx, dy int) [][]uint8 {
	//Declare the variable mySlice which is the 2D Slice we will return.
	var mySlice [][]uint8

	//Two dimensional loop looping through each row at the time.
	for y := 0; y <= dy; y++ {
		//Our innerSlice variable will hold the values used in each row of mySlice
		var innerSlice []uint8

		//The inner loop looping through a whole row in the x axis.
		for x := 0; x <= dx; x++ {

			/*	Here we append the row values into our innerSlice, the values are
				Current x * y, the outcome is a number which will represent a
				color where low number is dark and high is bright.

				Other patterns to try: (x+y) / 2, x*y, x^y
			*/
			innerSlice = append(innerSlice, uint8(x)*uint8(y))
		}
		//Append the row into the mySlice Slice.
		mySlice = append(mySlice, innerSlice)
	}

	//return our finished Slice
	return mySlice
}

func main() {
	/*	Function from imported package that draws an image using our Pic function.
		Browse through the Github link i provided by the import and you will see the code
		is very understandable.
	*/
	pic.Show(Pic)
}
