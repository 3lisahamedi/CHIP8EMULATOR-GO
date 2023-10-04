package instruct

import (
	"fmt"
	"math"
)

func SplitOpcode(wholeOpcode uint16) { /*Function that deals and manipulates opcodes*/
	/*	fmt.Println("Success. Now going to split opcode into 4 parts.")
	 */ /*Trying to isolate each digit*/
	/*Better version. Splits properly + into decimal notation we can switch case on later.*/
	firstDigit := (wholeOpcode & 0xF000) >> 12
	secondDigit := (wholeOpcode & 0x0F00) >> 8
	thirdDigit := (wholeOpcode & 0x00F0) >> 4
	fourthDigit := wholeOpcode & 0x000F

	/*Any prints need to be in %02x because they are hexadecimals ( and that's how you print them apparently*/
	/*	fmt.Printf("so we have %02x \n", wholeOpcode)
		fmt.Printf("which would give : %02x, as the firstDigit \n", firstDigit)
		fmt.Printf("which would give : %02x, as the secondDigit \n", secondDigit)
		fmt.Printf("which would give : %02x, as the thirdDigit \n", thirdDigit)
		fmt.Printf("which would give : %02x, as the fourthDigit \n", fourthDigit)*/

	switch firstDigit {
	case 0x00:
		switch secondDigit {
		case 00:
			switch thirdDigit {
			case 0x0E:
				switch fourthDigit {
				case 00:
					fmt.Println("Fell on the case 0x0E00. Congrats.")
					Cpu.opcode00E0()
					break
				}
				break
			}
			break
		}
	case 0x01:
		fmt.Println("Fell on the case 1NNN. Congrats.")
		/*fmt.Printf("secondDigit : %02x \n", secondDigit)
		fmt.Printf("thirdDigit : %02x \n", thirdDigit)
		fmt.Printf("fourthDigit : %02x \n", fourthDigit)*/
		Cpu.opcode1nnn(secondDigit, thirdDigit, fourthDigit)
		break
	case 0x06:
		fmt.Println("Fell on the case 6XNN/ 6XKK. Congrats.")
		Cpu.opcode6xnn(secondDigit, thirdDigit, fourthDigit)
		break
	case 0x0A:
		fmt.Println("Fell on the case ANNN. Congrats.")
		Cpu.opcodeAnnn(secondDigit, thirdDigit, fourthDigit)
		break
	case 0x0D:
		fmt.Println("Fell on the case DXYN. Congrats.")
		Cpu.opcodeDxyn(secondDigit, thirdDigit, fourthDigit)
	default:
		fmt.Println("Didn't get it.")
		break
	}
}

func (cpu *Chip8) opcode00E0() {
	for x := range cpu.Screen {
		for y := range cpu.Screen[x] {
			cpu.Screen[x][y] = 0
		}
	}

	/*	fmt.Println("Cleared screen.")
	 */
}

func (cpu *Chip8) opcodeDxyn(x uint16, y uint16, n uint16) {
	println("COMMENCEMENT : ", x, y, n)
	/*X, Y and N are 4-bit values -> hexadecimal, so uint16*/
	/*	fmt.Printf(" le X : %02x \n", x)
		fmt.Printf(" le Y : %02x \n", y)*/

	/*x, y coordinates. Width 8bits by default, height of n bits*/
	/*	width := byte(1)
		for i := 0; uint16(i) < n; i++ {
			cpu.Screen[byte(x)+width][y+uint16(i)] ^= 1
		}*/
	/* Drawing sprite at x, y with a width of 8 bits ( 1 byte ) and a height of n byte */
	/*0 et 1*/
	cpu.screenModified = false

	/*We make a copy, so we can check against makeshift screen ( the register ) */
	/*Never changes so no need for a loop on each*/
	screenY := y % 32
	var screenX byte

	/*Our collision flag*/
	cpu.register[0xF] = 0

	/*Height first according to n ( n is total height ), going line by line*/
	/*Get actual X coordinate adapted to CHIP-8 screen. X is what moves for the animations, not y */
	for i := 0; uint16(i) <= n; i++ {
		screenX = (byte(x) + byte(i)) % 64
		/*If not already. Turns it on*/
		cpu.Screen[screenY][screenX] ^= 1
		println("counter : ", i)

		/*Actual turn on or off of the pixel*/
		/*If already turned on*/
		cpu.screenModified = true
	}
	print("Current screen: \n")
	fmt.Print(cpu.Screen)
}

func (cpu *Chip8) opcodeAnnn(n1 uint16, n2 uint16, n3 uint16) {
	/*Three parts of a single value, NNN, that needs to be read as one*/
	/*Conversion from hexadecimal to decimal + making sure it fits in memory ( % )*/
	nnn := n1*uint16(math.Pow(16, 4)) + n2*uint16(math.Pow(16, 2)) + n3*uint16(math.Pow(16, 0))

	/*Set I = nnn*/
	cpu.I = nnn

	if cpu.I == nnn {
		/*		fmt.Println("index called and matches")
		 */
	}
}

func (cpu *Chip8) opcode6xnn(x uint16, n1 uint16, n2 uint16) {
	Vx := cpu.register[x]
	nn := int(n1)*int(math.Pow(16, 2)) + int(n2)*int(math.Pow(16, 0))%len(cpu.Memory)

	/*Set Vx = nn*/
	Vx = byte(nn)

	if Vx == byte(nn) {
		/*		fmt.Println("register value called and matches")
		 */
	}
}

func (cpu *Chip8) opcode1nnn(n1 uint16, n2 uint16, n3 uint16) {
	/*Set Pc = nnn */
	cpu.Pc = int(n1*uint16(math.Pow(16, 4)) + n2*uint16(math.Pow(16, 2)) + n3*uint16(math.Pow(16, 0)))
	/*	fmt.Printf("Program counter : %02x \n", cpu.Pc)
	 */
}
