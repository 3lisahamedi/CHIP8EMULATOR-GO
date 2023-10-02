package instruct

import (
	"fmt"
)

type opcode struct {
	x   byte
	y   byte
	n   byte
	nn  byte
	nnn byte
}

var recognisedOpcode opcode

func Opcode(input uint16) { /* function that grabs the hexadecimal and associates it to corresponding cmd*/
	fmt.Println("Success. Now going to split opcode into 4 parts.")
	splitOpcode(input)
	/* Need a split function. */
	/*op1 := byte(input>>12) & 0x000F*/

	/*	switch op1 {
	 *//*0xx for the hexadecimal. Goes through first parts. */

	/*case 0x0:
		fmt.Println("Currently in case 0x0.")
		break
	case 0x1:
		fmt.Println("Currently in case 0x1.")
		break
	case 0x2:
		fmt.Println("Currently in case 0x2.")
		break
	case 0x3:
		fmt.Println("Currently in case 0x3.")
		break
	case 0x4:
		fmt.Println("Currently in case 0x4.")
		break
	case 0x5:
		fmt.Println("Currently in case 0x5.")
		break
	case 0x6:
		fmt.Println("Currently in case 0x6.")
		break
	case 0x7:
		fmt.Println("Currently in case 0x7.")
		break
	case 0x8:
		fmt.Println("Currently in case 0x8.")
		break
	case 0x9:
		fmt.Println("Currently in case 0x9.")
		break
	case 0xA:
		fmt.Println("Currently in case 0xA.")
		break
	case 0xB:
		fmt.Println("Currently in case 0xB.")
		break
	case 0xC:
		fmt.Println("Currently in case 0xC.")
		break
	case 0xD:
		fmt.Println("Currently in case 0xD.")
		break
	case 0xE:
		fmt.Println("Currently in case 0xE.")
		break
	case 0xF:
		fmt.Println("Currently in case 0xF.")
		break
	}

	fmt.Println("Found no opcodes")*/
}

func splitOpcode(wholeOpcode uint16) opcode {
	/*Trying to isolate each digit*/
	/*firstDigit := wholeOpcode >> 12 & 0xF000
	secondDigit := wholeOpcode >> 8 & 0x0F00
	thirdDigit := wholeOpcode >> 4 & 0x00F0
	fourthDigit := wholeOpcode & 0x000F*/
	/*Better version. Splits properly + into decimal notation we can switch case on later.*/
	firstDigit := (wholeOpcode & 0xF000) >> 12
	secondDigit := (wholeOpcode & 0x0F00) >> 8
	thirdDigit := (wholeOpcode & 0x00F0) >> 4
	fourthDigit := wholeOpcode & 0x000F

	/*Any prints need to be in %02x because they are hexadecimals ( and that's how you print them apparently*/
	/*fmt.Printf("so we have %02x \n", wholeOpcode)
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

					break
				}
				break
			}
			break
		}
	case 0x01:
		fmt.Println("Fell on the case 1NNN. Congrats.")
		break
	case 0x06:
		fmt.Println("Fell on the case 6XNN. Congrats.")
		break
	case 0x0A:
		fmt.Println("Fell on the case ANNN. Congrats.")
		break
	case 0x0D:
		recognisedOpcode.x = byte(secondDigit)
		recognisedOpcode.y = byte(thirdDigit)
		recognisedOpcode.n = byte(fourthDigit)
		Cpu.opcodeDxyn(recognisedOpcode.x, recognisedOpcode.y, recognisedOpcode.n)
		fmt.Println("Fell on the case DXYN. Congrats.")
	default:
		fmt.Println("Didn't get it.")
		break
	}
	return opcode{}
}

func (cpu *Chip8) opcodeDxyn(x byte, y byte, n byte) {
	/* Drawing sprite at x, y with a width of 8 bits ( 1 byte ) and a height of n byte */
	/*0 et 1*/
	cpu.screenModified = false

	/*We make a copy, so we can check against makeshift screen ( the register ) */
	Vx := cpu.register[x]
	Vy := cpu.register[y]

	/*Never changes so no need for a loop on each*/
	screenY := Vy % 32
	var screenX byte

	/*Our collision flag*/
	cpu.register[0xF] = 0
	/*initialise in advance*/
	bytecount := byte(0)
	/*Get entire line that is supposed to be drawn in throught calculation + memory*/
	entireLine := cpu.memory[cpu.index+uint16(bytecount)]

	/*Height first according to n ( n is total height ), going line by line*/
	for ; bytecount < entireLine; bytecount++ {
		/*Width second according to default width ( 8 bits, so 8 ) */
		for width := byte(0); width < 8; width++ {
			/*Get actual X coordinate adapted to CHIP-8 screen. X is what moves for the animations, not y */
			screenX = Vx % 64
		}
	}

	/*Actual turn on or off of the pixel*/
	/*If already turned on*/
	if cpu.Screen[screenX][screenY] == 1 {
		cpu.register[0xF] = 1
	} else {
		/*If not already. Turns it on*/
		cpu.Screen[screenX][screenY] ^= 1
	}
	cpu.screenModified = true
}
