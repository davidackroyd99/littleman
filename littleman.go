package main

import (
	"fmt"
	"strconv"
)

func main() {
	Execute([]int16{103, 104, 902, 2, 3})
}

// LittleMan is the state of the CPU
type LittleMan struct {
	memory [100]int16
	ac     int16       // Accumulator
	pc     uint8       // Program Counter
	ir     uint16      // Instruction Register
	ar     uint8       // Address Register
	in     chan int16  // Input Area
	out    chan string // Output Area
}

// Execute a program on the little man
func Execute(data []int16) {
	man := LittleMan{}
	man.in = make(chan int16)
	man.out = make(chan string)

	for i, v := range data {
		man.memory[i] = v
	}

	go man.execute()

	for {
		output := <-man.out

		if output == "QUIT" {
			fmt.Println("The show's over and the little man has gone home.")
			break
		} else {
			fmt.Println(fmt.Sprintf("The little man says %v", output))
		}
	}
}

// Execute the little man
func (man *LittleMan) execute() {
	instruction := man.fetchInstruction()
	man.pc++

	switch man.ir {
	case 0:
		man.stop()
	case 1:
		man.add()
	case 2:
		man.sub()
	case 3:
		man.store()
	case 5:
		man.load()
	case 6:
		man.branch()
	case 7:
		man.branchZero()
	case 8:
		man.branchPositive()
	case 9:
		ok := man.io()
		if !ok {
			fmt.Println(fmt.Sprintf("Invalid i/o mode %d", man.ar))
		}
	default:
		fmt.Println(fmt.Sprintf("Invalid instruction %d", instruction))
	}

	if man.pc > 99 {
		man.out <- "QUIT"
	} else {
		man.execute()
	}
}

// Fetch the next instruction, fill the instruction and address registers and then return the instruction
func (man *LittleMan) fetchInstruction() int16 {
	instruction := man.memory[man.pc]
	man.ir = uint16(instruction / 100)
	man.ar = uint8(instruction % 100)

	return instruction
}

// Stop pauses execution (HLT)
func (man *LittleMan) stop() {
	return
}

// Add the value of memory pointed to by the address register to the accumulator
func (man *LittleMan) add() {
	man.ac += man.memory[man.ar]
}

// Sub the value of memory pointed to by the address register from the accumulator
func (man *LittleMan) sub() {
	man.ac -= man.memory[man.ar]
}

// Store writes the value in the accumulator to the address given
func (man *LittleMan) store() {
	man.memory[man.ar] = man.ac
}

// Load writes the value to the accumulator from the address given
func (man *LittleMan) load() {
	man.ac = man.memory[man.ar]
}

// Branch sets the program counter to the value in the address register
func (man *LittleMan) branch() {
	man.pc = man.ar
}

// BranchZero branches if the accumulator is zero
func (man *LittleMan) branchZero() {
	if man.ac == 0 {
		man.branch()
	}
}

// BranchPositive branches if the accumulator is zero or positive
func (man *LittleMan) branchPositive() {
	if man.ac >= 0 {
		man.branch()
	}
}

// IO handles input (if address register is 1) output (if address register is 2) or output as char (if address register is 22)
func (man *LittleMan) io() (ok bool) {
	switch man.ar {
	case 1:
		man.out <- "The little man wants input."
		man.ac = <-man.in
	case 2:
		man.out <- strconv.Itoa(int(man.ac))
	case 22:
		man.out <- string(man.ac)
	default:
		return false
	}

	return true
}
