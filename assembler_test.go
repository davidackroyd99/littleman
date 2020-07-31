package main

import "testing"

func TestCleanLine(t *testing.T) {
	t1 := cleanLine("ADD 1")

	if t1 != "add 1" {
		t.Errorf("cleanLine(\"ADD 1\") got %v, expected add 1", t1)
	}

	t2 := cleanLine("ADD 1 #this is a comment")

	if t2 != "add 1" {
		t.Errorf("cleanLine(\"ADD 1\") got %v, expected add 1", t2)
	}
}
