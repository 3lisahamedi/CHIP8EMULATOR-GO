package instruct

import (
	"fmt"
	"math"
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
		/*fmt.Printf("secondDigit : %02x \n", secondDigit)
		fmt.Printf("thirdDigit : %02x \n", thirdDigit)
		fmt.Printf("fourthDigit : %02x \n", fourthDigit)*/
		/*Cpu.opcode1nnn(secondDigit, thirdDigit, fourthDigit)*/
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
}

func (cpu *Chip8) opcodeDxyn(x uint16, y uint16, n uint16) {
	println("COMMENCEMENT : ", x, y, n)

	println("Vx : ", cpu.register[x]) /*Fueled in by previous opcodes ran previous to this one*/
	println("Vy : ", cpu.register[y]) /*Fueled in by previous opcodes ran previous to this one*/

	var screenY = cpu.register[y]
	var screenX = cpu.register[x]
	/*	screenX = (byte(x) + byte(i)) % 64*/
	var tempI = cpu.I

	/*Our collision flag*/
	cpu.register[0xF] = 0

	for increasingHeight := uint16(0); increasingHeight < n; increasingHeight++ { /*Going from cpu.I to n bytes in memory */
		println("increasingHeight : ", increasingHeight)
		println("cpu.I :", tempI)
		println("Memory at cpu.I placement ( start data ) : ", cpu.Memory[tempI])
		fmt.Printf("Binary  : %b ", cpu.Memory[tempI])
		spriteData := uint16(cpu.Memory[tempI])
		println("spriteData : ", spriteData)
		for bit := uint16(8); bit > 0; bit-- {
			isBitOne := (spriteData & (1 << bit)) != 0
			println("bit : ", bit)
			println("isBitOne : ", isBitOne)
			if isBitOne {
				println("screenX : ", screenX)
				println("correct x ? : ", cpu.register[x])
				println("screenY : ", screenY)
				println("correct y ? :", cpu.register[y])
				cpu.Screen[screenX][screenY] ^= 1
				println("cpu.register[0xF] : ", cpu.register[0xF])
			}
			screenX += 1
		}
		screenX = cpu.register[x]
		screenY += 1
		tempI += 1
	}
	print("Current screen: \n")
	fmt.Print(cpu.Screen)
	fmt.Println()
}

func (cpu *Chip8) opcodeAnnn(n1 uint16, n2 uint16, n3 uint16) {
	/*Three parts of a single value, NNN, that needs to be read as one*/
	fmt.Printf("COMMENCEMENT : %02x, %02x, %02x \n", n1, n2, n3)
	nnn := int(float64(n1)*math.Pow(16, 2)) + int(n2*16) + int(n3)
	/*Hexadecimal to decimal*/

	/*Set I = nnn*/
	cpu.I = uint16(nnn)

	if cpu.I == uint16(nnn) {
		println("At register I :", nnn) /*It's decimal and byte values are the same*/
	}
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

func (cpu *Chip8) opcode1nnn(n1 uint16, n2 uint16, n3 uint16) {
	/*Set Pc = nnn */
	cpu.Pc = (n1 * (16 * 16)) + n2*16 + n3

	println("Program counter : ", cpu.Pc) /*It's decimal and byte values are the same*/
}
