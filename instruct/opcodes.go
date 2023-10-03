package instruct

import (
	"fmt"
	"math"
)

func Opcode(input uint16) { /* function that grabs the hexadecimal and associates it to corresponding cmd*/
	fmt.Println("Success. Now going to split opcode into 4 parts.")
	splitOpcode(input)
	/* Need a split function. */
	/*op1 := byte(input>>12) & 0x000F*/

	/*	switch op1 {
	 */ /*0xx for the hexadecimal. Goes through first parts. */

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

func splitOpcode(wholeOpcode uint16) {
	/*Trying to isolate each digit*/
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

	fmt.Println("Cleared screen.")
}

func (cpu *Chip8) opcodeDxyn(x uint16, y uint16, n uint16) {
	/*X, Y and N are 4-bit values -> hexadecimal, so uint16*/
	fmt.Printf(" le X : %02x \n", x)
	fmt.Printf(" le Y : %02x \n", y)

	/*x, y coordinates. Width 8bits by default, height of n bits*/
	width := byte(1)
	cpu.Screen[byte(x)+width][byte(y)+byte(n)] = 1
}

func (cpu *Chip8) opcodeAnnn(n1 uint16, n2 uint16, n3 uint16) {
	/*Three parts of a single value, NNN, that needs to be read as one*/
	/*Conversion from hexadecimal to decimal + making sure it fits in memory ( % )*/
	nnn := (int(n1)*int(math.Pow(16, 4)) + int(n2)*int(math.Pow(16, 2)) + int(n3)*int(math.Pow(16, 0))) % len(cpu.memory)

	/*Set I = nnn*/
	cpu.index = nnn

	if cpu.index == nnn {
		fmt.Println("index called and matches")
	}
}

func (cpu *Chip8) opcode6xnn(x uint16, n1 uint16, n2 uint16) {
	Vx := cpu.register[x]
	nn := int(n1)*int(math.Pow(16, 2)) + int(n2)*int(math.Pow(16, 0))%len(cpu.memory)

	/*Set Vx = nn*/
	Vx = byte(nn)

	if Vx == byte(nn) {
		fmt.Println("register value called and matches")
	}
}

func (cpu *Chip8) opcode1nnn(n1 uint16, n2 uint16, n3 uint16) {
	cpu.Pc = n1*uint16(math.Pow(16, 4)) + n2*uint16(math.Pow(16, 2)) + n3*uint16(math.Pow(16, 0))
	fmt.Printf("Program counter : %02x \n", cpu.Pc)
}
