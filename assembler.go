package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type DecimalInstruction struct {
	mnemonic string
	address  int16
}

// TODO implement variables
// Assemble the code in filePath and return it
func Assemble(filePath string) []int16 {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var machineCode []int16
	for scanner.Scan() {
		if len(scanner.Text()) >= 3 {
			machineCode = append(machineCode, parseInstruction(createInstruction(scanner.Text())))
		}
	}

	return machineCode
}

func cleanLine(line string) string {
	return strings.ToLower(strings.TrimSpace(strings.Split(line, "#")[0]))
}

func createInstruction(line string) DecimalInstruction {
	cleaned := cleanLine(line)

	var instruction DecimalInstruction
	instruction.mnemonic = cleaned[:3]

	instruction.address = parseAddress(cleaned)

	return instruction
}

// TODO test me
func parseAddress(line string) int16 {
	for i := 3; i > 0; i-- {
		address, err := strconv.Atoi(line[len(line)-i:])

		if err == nil {
			return int16(address)
		}
	}

	return 0
}

// TODO add some sort of error handling
func parseInstruction(instruction DecimalInstruction) int16 {
	var opcode int16
	appendAddress := true

	switch instruction.mnemonic {
	case "hlt":
		opcode = 0
		appendAddress = false
	case "add":
		opcode = 100
	case "sub":
		opcode = 200
	case "sta":
		opcode = 300
	case "lda":
		opcode = 500
	case "bra":
		opcode = 600
	case "brz":
		opcode = 700
	case "brp":
		opcode = 800
	case "inp":
		opcode = 901
		appendAddress = false
	case "out":
		opcode = 902
		appendAddress = false
	case "otc":
		opcode = 922
		appendAddress = false
	case "dat":
		return instruction.address
	}

	if appendAddress {
		return opcode + int16(instruction.address)
	} else {
		return opcode
	}
}
