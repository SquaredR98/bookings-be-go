package main

import "testing"

func TestRun(t *testing.T) {
	err := runServer()

	if err != nil {
		t.Error("failed runServer()")
	}
}
