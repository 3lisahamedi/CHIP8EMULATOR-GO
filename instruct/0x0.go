package instruct

func (cpu *Chip8) opcode00E0() {
	for x := range cpu.Screen {
		for y := range cpu.Screen[x] {
			cpu.Screen[x][y] = 0
		}
	}
}
