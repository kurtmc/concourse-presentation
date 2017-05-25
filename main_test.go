package main

import "testing"

func TestMain(t *testing.T) {
	expected := "Concourse!!!"
	actual := getString()

	if expected != actual {
		t.Errorf("%s != %s\n", expected, actual)
	}
}
