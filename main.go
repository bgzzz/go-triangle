package main

import (
	"flag"
	"fmt"
)

//error messages
//I have not decided to use logging packages for simplicity purposes
const (
	NotThreeSidesErrorMsg = `3 sides are not provided as input`
	ParsingErrorMsg       = `There is a parsing error of the value`
)

const (
	SuccessOutput = `Current triangle is %s`
)

func main() {

	// prepare to parse arguments
	flag.Usage = usage
	verboseGoErr := flag.Bool("v", false, "-v parameter to verbose low level error messages")
	flag.Parse()
	args := flag.Args()

	// if there are no three values in the input
	if len(args) != 3 {
		exitWithError(NotThreeSidesErrorMsg)
	}

	sides, err := parseArgs(args)
	if err != nil {
		if *verboseGoErr {
			printErr(err.Error())
		}
		exitWithError(ParsingErrorMsg)
	}

	triangleType, err := CategorizeTriangle(sides)

	if err != nil {
		exitWithError(err.Error())
	}

	fmt.Printf(SuccessOutput, triangleType)

}
