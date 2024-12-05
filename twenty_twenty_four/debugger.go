package main

import "testing"

func Test_twoa(t *testing.T) {
	err := FourA() // change to day, run test with debugging
	if err != nil {
		t.Fatal(err)
	}
}
