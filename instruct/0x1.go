package instruct

import (
	"fmt"
)

func (cpu *Chip8) opcode1nnn(nnn uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", nnn)

	/*	time.Sleep(time.Second * 1) // For debug */
	/*Set Pc = nnn */
	cpu.Pc = nnn - 2
}
