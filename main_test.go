package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	expectedResult := "{\"message\":\"hello\"}"
	actualResult := Hello()

	if actualResult != expectedResult {
		t.Errorf("Did not get the correct result")
	}
}

func TestOkay(t *testing.T) {
	expectedResult := "ok"
	actualResult := Okay()

	if actualResult != expectedResult {
		t.Errorf("Did not get the correct result")
	}
}
