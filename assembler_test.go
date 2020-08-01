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

func TestCreateInstruction(t *testing.T) {
	if ins := createInstruction("ADD 21"); ins.mnemonic != "add" || ins.address != 21 {
		t.Errorf("createInstruction(\"ADD 21\") got %v %d, expected 'add' 21", ins.mnemonic, ins.address)
	}
	if ins2 := createInstruction("ADD 1"); ins2.mnemonic != "add" || ins2.address != 1 {
		t.Errorf("createInstruction(\"ADD 1\") got %v %d, expected 'add' 1", ins2.mnemonic, ins2.address)
	}
}
