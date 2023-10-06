package instruct

import "fmt"

func (cpu *Chip8) opcode2nnn(nnn uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n ", nnn)

	/*Call subroutine at nnn*/
	/*	fmt.Printf("nnn currently : %02x \n", nnn)*/
	/*	fmt.Printf("cpu.sp currently ( decimal ) : %v \n", cpu.sp)
		fmt.Printf("cpu.PC currently ( decimal ) : %v \n", cpu.Pc)*/
	cpu.sp++
	cpu.stack[cpu.sp] = cpu.Pc
	/*	fmt.Printf("what was put in cpu.stack ( decimal ): %v \n", cpu.stack[cpu.sp])*/
	cpu.Pc = nnn

	/*	fmt.Printf("cpu.PC with nnn value now : %02x \n", cpu.Pc)*/
}
