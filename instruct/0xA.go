package instruct

import "fmt"

func (cpu *Chip8) opcodeAnnn(nnn uint16) {
	/*Three parts of a single value, NNN, that needs to be read as one*/
	fmt.Printf("COMMENCEMENT : %02x \n", nnn)
	/*Set I = nnn*/
	cpu.I = nnn
}
