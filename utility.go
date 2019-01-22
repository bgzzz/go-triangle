package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// characters that are used for trimming while parsing the input list of
	// triangle sides
	TrimCharachters = ", 	;"
)

const (
	UsageOutput = `go-triangle is the utility to determine a type of the triangle
Usage: go-triangle [options] <a_side> <b_side> <c_side>
Options:
	-v	parameter to verbose low level error messages
Example: go-triangle 10 20 15
`
)

const (
	PrefixErrorMsg = `ERROR`
)

func usage() {
	fmt.Printf(UsageOutput)
}

func exitWithError(msg string) {
	printErr(msg)
	usage()
	os.Exit(2)
}

func printErr(msg string) {
	fmt.Printf("%s: %s\n", PrefixErrorMsg, msg)
}

func parseArgs(args []string) ([]float64, error) {
	var sides []float64
	for i := range args {
		sideStr := strings.Trim(args[i], TrimCharachters)
		sideFloat, err := strconv.ParseFloat(sideStr, 64)
		if err != nil {
			return nil, err
		}
		sides = append(sides, sideFloat)
	}

	return sides, nil
}
