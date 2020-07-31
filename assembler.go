package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

//func parseInstruction(string instruction) int16 {
//	switch instruction[:2] {
//	case "add":
//	}
//}
