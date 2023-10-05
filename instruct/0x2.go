package instruct

import "fmt"

func (cpu *Chip8) opcode2nnn(nnn uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n ", nnn)

	/*Call subroutine at nnn*/
	cpu.sp++
	cpu.stack[cpu.sp] = cpu.Pc
	cpu.Pc = nnn
}
