package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type DecimalInstruction struct {
	mnemonic string
	address  uint8
}

func Assemble(filePath string) []int16 {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return []int16{1, 1, 1}
}

func cleanLine(line string) string {
	return strings.ToLower(strings.TrimSpace(strings.Split(line, "#")[0]))
}

func createInstruction(line string) DecimalInstruction {
	cleaned := cleanLine(line)

	var instruction DecimalInstruction
	instruction.mnemonic = cleaned[:3]

	addr := cleaned[len(cleaned)-2:]
	address, err := strconv.Atoi(strings.TrimSpace(addr))

	if err == nil {
		instruction.address = uint8(address)
	}

	return instruction
}

func lexInstructions(lines []string) []DecimalInstruction {
	var instructions []DecimalInstruction

	for _, line := range lines {
		instructions = append(instructions, createInstruction(line))
	}

	return instructions
}

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
	}

	if appendAddress {
		return opcode + int16(instruction.address)
	} else {
		return opcode
	}
}
