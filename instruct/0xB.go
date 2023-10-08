package instruct

import (
	"fmt"
)

func (cpu *Chip8) opcodeBnnn(nnn uint16) {
	/*Jump to location nnn + V0*/
	fmt.Printf("COMMENCEMENT : %02x \n", nnn)

	cpu.Pc = nnn + uint16(cpu.register[0])
	fmt.Printf("cpu.PC : %02x \n", cpu.Pc)
}
