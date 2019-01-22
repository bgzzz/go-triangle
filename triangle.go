package main

import (
	"errors"
)

const (
	NotATriangleErrorMsg = "This is not a triangle"
)

// I made it in separate functions for readability
// on the other hand it can validate and categorize at the same time if better performance needed

// made both functions from capital letter in case we decided to put it to the separate package

// ValidateTriangle determines whether triangle can or
// can not be made from sides length array
func ValidateTriangle(sides []float64) bool {
	// made this check in case function will be used somewhere else
	// for the usage in this particular app can be commented
	if len(sides) != 3 {
		return false
	}

	// make it in a loop fashion
	// could be done via strict indexing in array
	for i := range sides {
		sum := 0.0
		for j := range sides {
			if i != j {
				sum += sides[j]
			}
		}

		if sum <= sides[i] {
			return false
		}
	}

	return true
}

// CategorizeTriangle analyses triangle based on the array of provided sides
// It returns the error if is is impossible to form a triangle based on the
// provided sides lengths
// if no error occurred it returns the type of the triangle
func CategorizeTriangle(sides []float64) (string, error) {

	if !ValidateTriangle(sides) {
		return "", errors.New(NotATriangleErrorMsg)
	}

	// here we are counting how many sides of the same length we got
	sideCounter := map[float64]int{}
	currentCategory := ScaTriangleType
	for i := range sides {
		sideCounter[sides[i]] += 1

		// check how many sides have the same length
		// and determine the type of the triangle
		if sideCounter[sides[i]] == 2 {
			currentCategory = IsoTriangleType
		} else if sideCounter[sides[i]] == 3 {
			currentCategory = EqTriangleType
		}
	}

	return currentCategory, nil

}
