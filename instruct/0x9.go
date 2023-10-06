package instruct

import "fmt"

func (cpu *Chip8) opcode9xy0(x uint16, y uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x \n", x, y)

	/*Skip next instruction if Vx != Vy*/
	if cpu.register[x] != cpu.register[y] {
		cpu.Pc += 2
	}
}
