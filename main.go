// bytecode interpreter
// 32-bit stack machine with minimized code size (for IoT application)

package main

import "fmt"

// config

const (
	Msz = 0x1000 // main memory size, bytes
	Rsz = 0x100  // return stack size, cells
	Dsz = 0x10   // data stack size, cells
)

// ============================================ virtual machine data structures

var Sp uint = 0 //data stack pointer

var Rp uint = 0 // return stack pointer

var Ip uint = 0 // instruction pointer
var Cp uint = 0 // compiler pointer

// ======================================================= bytecode definitions

const (
	NOP = iota
	JMP
	qJMP
	CALL
	RET
	LIT

	DUP
	DROP
	SWAP
	OVER
	PRESS
	DROPALL

	ADD
	SUB
	MUL
	DIV
)

// ================================================================ system init

func main() {
	fmt.Println("")
	fmt.Println("bytecode interpreter")
	fmt.Println("")
	fmt.Println("main memory\tM[", Msz, "]")
	fmt.Println("return stack\tM[", Rsz, "]")
	fmt.Println("data stack\tM[", Dsz, "]")
	fmt.Println("")
}
