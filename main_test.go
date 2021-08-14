package main

import "testing"

func TestDivide(t *testing.T) {
	_, error := Divide(10.0, 1.0)

	if error != nil {
		t.Error("got an error  when we should not have")
	}
}
