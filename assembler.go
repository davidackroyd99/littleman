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
	fmt.Printf(addr)
	address, err := strconv.Atoi(addr)
	if err == nil {
		instruction.address = uint8(address)
	}

	return instruction
}

// func parseInstruction(string instruction) int16 {
//	switch instruction[:2] {
//	case "add":
//
//	}
//}
