/*
*	FILE			: toHex.go
*	PROJECT			: RGB-to-Hex
*	PROGRAMMER		: Daniel Pieczewski
*	FIRST VERSION	: 2019-12-01
*	DESCRIPTION		:
*		The function(s) in this file take red, green, and blue values from the command line
*		and return the hex code representation of the corresponding colour
 */
package main

import (
	"fmt"
	"os"
	"strconv"
)

//Constants
//MinColourValue = Minimum a colour can be in RGB format
//MaxColourValue = Maximum a colour can be in RGB format
//BadArgs = Error code for the wrong number of arguments
//BadValue = Error code for an invalid R, G, or B value
const constMinColourValue int = 0
const constMaxColourValue int = 255
const constBadArgs int = -1
const constBadValue int = -2

func main() {
	//Find argc
	var argc int = len(os.Args)
	//If there are too many command line args:
	if argc != 4 {
		fmt.Printf("Error! Incorrect number of arguments!\n")
		fmt.Printf("Usage: toHex <r> <g> <b>\n")
		os.Exit(constBadArgs)
	}

	//Get the cmd arguments and convert to ints
	//Garbage the error values by assigning them to '_'
	var red, _ = strconv.Atoi(os.Args[1])
	var green, _ = strconv.Atoi(os.Args[2])
	var blue, _ = strconv.Atoi(os.Args[3])

	//Check the ranges of all the inputs
	//Return an error is any of them are invalid
	if !checkRange(red, constMinColourValue, constMaxColourValue) {
		fmt.Printf("Red value is invalid!\n")
		os.Exit(constBadValue)
	}
	if !checkRange(green, constMinColourValue, constMaxColourValue) {
		fmt.Printf("Green value is invalid!\n")
		os.Exit(constBadValue)
	}
	if !checkRange(blue, constMinColourValue, constMaxColourValue) {
		fmt.Printf("Blue value is invalid!\n")
		os.Exit(constBadValue)
	}

	//Convert to hex and combine into one string
	var hexString string = fmt.Sprintf("%02X%02X%02X", red, green, blue)

	//Print the result to the cmd
	fmt.Printf("\t#%s\n\n", hexString)

	//Exit the program
	os.Exit(0)
}

/*
*	FUNCTION		: checkRange
*	DESCRIPTION		: Determines whether or not a number is within a given range (inclusive)
*
*	PARAMETERS		:
*		int test	: Value to check
*		int min		: Minimum of range (inclusive)
*		int	max		: Maximum of range (inclusive)
*
*	RETURNS			:
*		bool		: true if in range, false otherwise
 */
func checkRange(test int, min int, max int) bool { return (min <= test && test <= max) }
