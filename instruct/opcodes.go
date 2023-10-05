package instruct

import (
	"fmt"
)

func SplitOpcode(wholeOpcode uint16) { /*Function that deals and manipulates opcodes*/
	firstDigit := (wholeOpcode & 0xF000) >> 12
	secondDigit := (wholeOpcode & 0x0F00) >> 8
	thirdDigit := (wholeOpcode & 0x00F0) >> 4
	fourthDigit := wholeOpcode & 0x000F

	if firstDigit == uint16(0) && secondDigit == uint16(0) && thirdDigit == uint16(0) && fourthDigit == uint16(0) {
		return
	}
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
		Cpu.opcode1nnn(wholeOpcode & 0x0FFF)
		break
	case 0x06:
		fmt.Println("Fell on the case 6XNN/ 6XKK. Congrats.")
		Cpu.opcode6xnn(secondDigit, thirdDigit, fourthDigit)
		break
	case 0x07:
		fmt.Println("Fell on the case 7XNN/ 7XKK. Congrats.")
		Cpu.opcode7xnn(secondDigit, thirdDigit, fourthDigit)
	case 0x0A:
		fmt.Println("Fell on the case ANNN. Congrats.")
		Cpu.opcodeAnnn(wholeOpcode & 0x0FFF)
		break
	case 0x0D:
		fmt.Println("Fell on the case DXYN. Congrats.")
		Cpu.opcodeDxyn(secondDigit, thirdDigit, fourthDigit)
	default:
		fmt.Println("Didn't get it.")
		break
	}
}
