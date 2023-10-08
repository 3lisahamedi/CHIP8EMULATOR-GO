package instruct

import (
	"fmt"
	"os"
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
				case 0x0E:
					fmt.Println("Fell on the case 00EE. Congrats.")
					Cpu.opcode00EE()
					break
				case 00:
					fmt.Println("Fell on the case 0E00. Congrats.")
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
	case 0x02:
		fmt.Println("Fell on the case 2NNN. Congrats.")
		Cpu.opcode2nnn(wholeOpcode & 0x0FFF)
		break
	case 0x03:
		fmt.Println("Fell on the case 3XNN/ 3XKK. Congrats.")
		Cpu.opcode3xnn(secondDigit, thirdDigit, fourthDigit)
		break
	case 0x04:
		fmt.Println("Fell on the case 4XNN/ 4XKK. Congrats.")
		Cpu.opcode4xnn(secondDigit, thirdDigit, fourthDigit)
		break
	case 0x05:
		fmt.Println("Fell on the case 5XY0. Congrats.")
		Cpu.opcode5xy0(secondDigit, thirdDigit)
		break
	case 0x06:
		fmt.Println("Fell on the case 6XNN/ 6XKK. Congrats.")
		Cpu.opcode6xnn(secondDigit, thirdDigit, fourthDigit)
		break
	case 0x07:
		fmt.Println("Fell on the case 7XNN/ 7XKK. Congrats.")
		Cpu.opcode7xnn(secondDigit, thirdDigit, fourthDigit)
		break
	case 0x08:
		switch fourthDigit {
		case 0x00:
			fmt.Println("Fell on the case 8XY0. Congrats.")
			Cpu.opcode8xy0(secondDigit, thirdDigit)
			break
		case 0x01:
			fmt.Println("Fell on the case 8XY1. Congrats.")
			Cpu.opcode8xy1(secondDigit, thirdDigit)
			break
		case 0x02:
			fmt.Println("Fell on the case 8XY2. Congrats.")
			Cpu.opcode8xy2(secondDigit, thirdDigit)
			break
		case 0x03:
			fmt.Println("Fell on the case 8XY3. Congrats.")
			Cpu.opcode8xy3(secondDigit, thirdDigit)
			break
		case 0x04:
			fmt.Println("Fell on the case 8XY4. Congrats.")
			Cpu.opcode8xy4(secondDigit, thirdDigit)
			break
		case 0x05:
			fmt.Println("Fell on the case 8XY5. Congrats.")
			Cpu.opcode8xy5(secondDigit, thirdDigit)
			break
		case 0x06:
			fmt.Println("Fell on the case 8XY6. Congrats.")
			Cpu.opcode8xy6(secondDigit, thirdDigit)
			break
		case 0x07:
			fmt.Println("Fell on the case 8XY7. Congrats.")
			Cpu.opcode8xy7(secondDigit, thirdDigit)
			break
		case 0x0E:
			fmt.Println("Fell on the case 8XYE. Congrats.")
			Cpu.opcode8xyE(secondDigit, thirdDigit)
			break
		}
	case 0x09:
		fmt.Println("Fell on the case 9XY0. Congrats.")
		Cpu.opcode9xy0(secondDigit, thirdDigit)
		break
	case 0x0A:
		fmt.Println("Fell on the case ANNN. Congrats.")
		Cpu.opcodeAnnn(wholeOpcode & 0x0FFF)
		break
	case 0x0B:
		fmt.Println("Fell on the case BNNN. Congrats.")
		os.Exit(333)
		Cpu.opcodeBnnn(secondDigit)
		break
	case 0x0C:
		fmt.Println("Fell on the case CXNN/CXKK. Congrats.")
		Cpu.opcodeCxnn(secondDigit, thirdDigit, fourthDigit)
		break
	case 0x0D:
		fmt.Println("Fell on the case DXYN. Congrats.")
		Cpu.opcodeDxyn(secondDigit, thirdDigit, fourthDigit)
		break
	case 0x0E:
		switch thirdDigit {
		case 0x09:
			fmt.Println("Fell on the case EX9E. Congrats.")
			Cpu.opcodeEx9e(secondDigit)
			break
		case 0x0A:
			fmt.Println("Fell on the case EXA1. Congrats.")
			Cpu.opcodeExa1(secondDigit)
			break
		}
	case 0x0F:
		switch thirdDigit {
		case 0x02:
			fmt.Println("Fell on the case FX29. Congrats.")
			Cpu.opcodeFx29(secondDigit)
			break
		case 0x03:
			fmt.Println("Fell on the case FX33. Congrats.")
			Cpu.opcodeFx33(secondDigit)
			break
		case 0x05:
			fmt.Println("Fell on the case FX55. Congrats.")
			Cpu.opcodeFx55(secondDigit)
			break
		case 0x06:
			fmt.Println("Fell on the case FX65. Congrats.")
			Cpu.opcodeFx65(secondDigit)
			break
		}
		switch fourthDigit {
		case 0x05:
			fmt.Println("Fell on the case FX15. Congrats.")
			Cpu.opcodeFx15(secondDigit)
			break
		case 0x07:
			fmt.Println("Fell on the case FX07. Congrats.")
			Cpu.opcodeFx07(secondDigit)
			break
		case 0x0A:
			fmt.Println("Fell on the case FX0A. Congrats.")
			Cpu.OpcodeFx0A(secondDigit)
			break
		case 0x08:
			fmt.Println("Fell on the case FX18. Congrats.")
			Cpu.opcodeFx18(secondDigit)
			break
		case 0x0E:
			fmt.Println("Fell on the case FX1E. Congrats.")
			Cpu.opcodeFx1E(secondDigit)
			break
		}
	default:
		fmt.Println("Didn't get it.")
		break
	}
}
