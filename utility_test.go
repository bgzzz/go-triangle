// +build unit

package main

import (
	"reflect"
	"testing"
)

func TestParseArgError(t *testing.T) {
	inputArgs := []string{"12,", "10 ", "242.34sdfsdf"}

	_, err := parseArgs(inputArgs)

	if err != nil {
		t.Log(err.Error())
	} else {
		t.Error("It should be an error")
	}

}

func TestParseArg(t *testing.T) {
	inputArgs := []string{"12,", "10 ", "242.34	,", "242.22	;"}
	expectedArray := []float64{12, 10, 242.34, 242.22}

	parsedArray, err := parseArgs(inputArgs)

	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(parsedArray, expectedArray) {
		t.Errorf("Parsed array != expected (%+v %+v)",
			parsedArray, expectedArray)
	}

}
