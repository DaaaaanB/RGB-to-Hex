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
	"log"
	"math"
	"os"
	"strconv"
)

//Constants
//BadArgs = Error code for the wrong number of arguments
//BadValue = Error code for an invalid R, G, or B value
const constBadArgs int = -1
const constBadValue int = -2

func main() {
	//If there are an incorrect amount of command line args:
	if len(os.Args) != 4 {
		log.Printf("Error! Incorrect number of arguments!\n")
		log.Printf("Usage: toHex <r> <g> <b>\n\n")
		os.Exit(constBadArgs)
	}

	//Get the cmd arguments and convert to ints
	//Garbage the error values by assigning them to '_'
	red, err := strconv.Atoi(os.Args[1])
	green, err := strconv.Atoi(os.Args[2])
	blue, err := strconv.Atoi(os.Args[3])

	//IF the error variable is filled:
	if err != nil {
		log.Printf("One of your values is not an integer!\n\n")
		os.Exit(constBadValue)
	}

	//Check the ranges of all the inputs
	//Return an error if any of them are invalid
	if !checkRange(red, 0, math.MaxUint8) {
		log.Printf("Red value is invalid!\n\n")
		os.Exit(constBadValue)
	}
	if !checkRange(green, 0, math.MaxUint8) {
		log.Printf("Green value is invalid!\n\n")
		os.Exit(constBadValue)
	}
	if !checkRange(blue, 0, math.MaxUint8) {
		log.Printf("Blue value is invalid!\n\n")
		os.Exit(constBadValue)
	}

	//Convert to hex and combine into one string
	hexString := fmt.Sprintf("%02X%02X%02X", red, green, blue)

	//Print the result to the cmd
	fmt.Printf("\t#%s\n\n", hexString)
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
