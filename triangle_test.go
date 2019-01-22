package main

import (
	"testing"
)

func triangleTestValidation(sides []float64, expectedType string, t *testing.T) {
	triangleType, err := CategorizeTriangle(sides)
	if err != nil {
		t.Error(err.Error())
	}

	if triangleType != expectedType {
		t.Fatalf("Triangle Type != Expected Type (%s != %s)", triangleType, expectedType)
	}
}

func TestNotTriangle(t *testing.T) {
	sides := []float64{10, 15.5, 25.5}

	if ValidateTriangle(sides) {
		t.Fatalf("This is not a triangle: (sides: %+v)", sides)
	}
}

func TestTriangle(t *testing.T) {
	sides := []float64{10, 15.5, 24.5}

	if !ValidateTriangle(sides) {
		t.Fatalf("This is a triangle (sides: %+v)", sides)
	}
}

func TestEquilateral(t *testing.T) {
	sides := []float64{125.3, 125.3, 125.3}

	expectedType := `equilateral`

	triangleTestValidation(sides, expectedType, t)
}

func TestIsoscales(t *testing.T) {
	sides := []float64{125.3, 125.3, 233.334345}

	expectedType := `isosceles`

	triangleTestValidation(sides, expectedType, t)
}

func TestScalene(t *testing.T) {
	sides := []float64{115.7, 125.3, 233.334345}

	expectedType := `scalene`

	triangleTestValidation(sides, expectedType, t)
}

func TestErrorCategorization(t *testing.T) {
	sides := []float64{115.7, 125.3, 233.334345, 233.334345}

	_, err := CategorizeTriangle(sides)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Errorf("There should be an ERROR because it is not triangle (%+v)", sides)
	}
}
