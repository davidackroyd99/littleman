package main

import (
	"testing"
)

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
	if ins3 := createInstruction("HLT"); ins3.mnemonic != "hlt" || ins3.address != 0 {
		t.Errorf("createInstruction(\"HLT\") got %v %d, expected 'add' 1", ins3.mnemonic, ins3.address)
	}
	if ins4 := createInstruction("DAT 123"); ins4.mnemonic != "dat" || ins4.address != 123 {
		t.Errorf("createInstruction(\"DAT 123\") got %v %d, expected 'dat' 123", ins4.mnemonic, ins4.address)
	}
}

func TestParseInstruction(t *testing.T) {
	if ins := parseInstruction(DecimalInstruction{"hlt", 21}); ins != 0 {
		t.Errorf("parseInstruction(DecimalInstruction{\"hlt\", 21} got %d, expected 0", ins)
	}
	if ins := parseInstruction(DecimalInstruction{"brz", 21}); ins != 721 {
		t.Errorf("parseInstruction(DecimalInstruction{\"brz\", 21} got %d, expected 721", ins)
	}
}
