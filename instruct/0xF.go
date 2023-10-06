package instruct

import "fmt"

func (cpu *Chip8) opcodeFx65(x uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", x)

	/*Fills V0 to VX with values from memory starting at address I*/
	for i := uint16(0); i <= x; i++ {
		cpu.register[i] = cpu.Memory[cpu.I+i]
	}
	cpu.I += x + 1
}

func (cpu *Chip8) opcodeFx55(x uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", x)

	/*Stores V0 to VX in memory starting at address I*/
	for i := uint16(0); i <= x; i++ {
		cpu.Memory[cpu.I+i] = cpu.register[i]
	}
	cpu.I += x + 1
}

func (cpu *Chip8) opcodeFx33(x uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", x)

	/*Stores BCD representation of Vx in memory locations I, I+1, I+2*/
	number := cpu.register[x]
	fmt.Printf(" Original nb : %v \n", number)
	V0 := number / 100
	fmt.Printf(" V0 : %v \n", V0)
	V1 := (number % 100) / 10
	fmt.Printf(" V1 : %v \n", V1)
	V2 := (number % 100) % 10
	fmt.Printf(" V2 : %v \n", V2)

	cpu.Memory[cpu.I] = V0
	cpu.Memory[cpu.I+1] = V1
	cpu.Memory[cpu.I+2] = V2
}

func (cpu *Chip8) opcodeFx1E(x uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", x)

	/*Set I = I + Vx*/
	cpu.I = cpu.I + uint16(cpu.register[x])
}
