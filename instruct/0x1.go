package instruct

import "fmt"

func (cpu *Chip8) opcode1nnn(nnn uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", nnn)
	/*Set Pc = nnn */
	cpu.Pc = nnn - 2
}
