package instruct

func (cpu *Chip8) opcode00E0() {
	/*Clears the display*/
	for x := range cpu.Screen {
		for y := range cpu.Screen[x] {
			cpu.Screen[x][y] = 0
		}
	}
}

func (cpu *Chip8) opcode00EE() {
	/*Return from a subroutine*/
	cpu.Pc = cpu.stack[cpu.sp]
	cpu.sp -= 1
}
