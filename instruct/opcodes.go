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

func (cpu *Chip8) opcode00E0() {
	for x := range cpu.Screen {
		for y := range cpu.Screen[x] {
			cpu.Screen[x][y] = 0
		}
	}
}

func (cpu *Chip8) opcodeDxyn(x uint16, y uint16, n uint16) {
	println("COMMENCEMENT : ", x, y, n)
	var screenY = cpu.register[y]
	var screenX = cpu.register[x]
	var tempI = cpu.I

	/*Our collision flag*/
	cpu.register[0xF] = 0

	for increasingHeight := uint16(0); increasingHeight < n; increasingHeight++ { /*Going from cpu.I to n bytes in memory */

		spriteData := uint16(cpu.Memory[tempI])
		for bit := uint16(0); bit <= 8; bit++ {
			test := uint16(8) - bit
			isBitOne := (spriteData & (1 << test)) != 0
			if isBitOne {
				cpu.Screen[screenX][screenY] ^= 1
			}
			screenX += 1
		}
		screenX = cpu.register[x]
		screenY += 1
		tempI += 1
	}
}

func (cpu *Chip8) opcodeAnnn(nnn uint16) {
	/*Three parts of a single value, NNN, that needs to be read as one*/
	fmt.Printf("COMMENCEMENT : %02x \n", nnn)
	/*Set I = nnn*/
	cpu.I = nnn
}

func (cpu *Chip8) opcode6xnn(x uint16, n1 uint16, n2 uint16) {
	println("COMMENCEMENT : ", x, n1, n2)
	nn := n1*16 + n2 /*Hexadecimal to decimal calculation*/

	/*Set Vx = nn*/
	cpu.register[x] = byte(nn)

	if cpu.register[x] == byte(nn) {
		println("At register", x, " : ", cpu.register[x]) /*It's decimal and byte values are the same*/
	}
}

func (cpu *Chip8) opcode1nnn(nnn uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", nnn)
	/*Set Pc = nnn */
	cpu.Pc = nnn - 2
}
