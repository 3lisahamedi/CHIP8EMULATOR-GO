package instruct

import "fmt"

func (cpu *Chip8) opcodeDxyn(x uint16, y uint16, n uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x, %02x \n", x, y, n)
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
