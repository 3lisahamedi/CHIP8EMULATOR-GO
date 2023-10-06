package instruct

import "fmt"

func (cpu *Chip8) opcode8xy0(x uint16, y uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x \n", x, y)

	/*Set Vx = Vy*/
	cpu.register[x] = cpu.register[y]
}

func (cpu *Chip8) opcode8xy1(x uint16, y uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x \n", x, y)

	/*Set Vx = Vx OR Vy*/
	cpu.register[x] |= cpu.register[y]
}

func (cpu *Chip8) opcode8xy2(x uint16, y uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x \n", x, y)

	/*Set Vx = Vx AND Vy*/
	cpu.register[x] &= cpu.register[y]
}

func (cpu *Chip8) opcode8xy3(x uint16, y uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x \n", x, y)

	/*Set Vx = Vx XOR Vy*/
	cpu.register[x] ^= cpu.register[y]
}

func (cpu *Chip8) opcode8xy4(x uint16, y uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x \n", x, y)

	/*Set Vx = Vx + Vy, set VF = carry*/
	newVX := cpu.register[x] + cpu.register[y]
	if newVX > 255 {
		cpu.register[0xF] = 1
	} else {
		cpu.register[0xF] = 0
	}
	/*Only takes the lowest 8 bits*/
	cpu.register[x] = newVX & 0xFF
}

func (cpu *Chip8) opcode8xy5(x uint16, y uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x \n", x, y)

	/*Set Vx = Vx - Vy, set VF = NOT borrow*/
	if cpu.register[x] > cpu.register[y] {
		cpu.register[0xF] = 1
		/*equal case is not taken into account ?*/
	} else if cpu.register[x] < cpu.register[y] {
		cpu.register[0xF] = 0
	}
	/*Order of subtraction is important so doing it clearer*/
	cpu.register[x] = cpu.register[x] - cpu.register[y]
}

func (cpu *Chip8) opcode8xy6(x uint16, y uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x \n", x, y)

	/*Set Vx = Vx SHR 1*/
	if cpu.register[x]&1 == 1 {
		cpu.register[0xF] = 1
	} else {
		cpu.register[0xF] = 0
	}
	cpu.register[x] /= 2
}

func (cpu *Chip8) opcode8xy7(x uint16, y uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x \n", x, y)

	/*Set Vx = Vy - Vx, set VF = NOT borrow*/
	if cpu.register[y] > cpu.register[x] {
		cpu.register[0xF] = 1
		/*equal case is not taken into account ?*/
	} else if cpu.register[y] < cpu.register[x] {
		cpu.register[0xF] = 0
	}
	/*Order of subtraction is important so doing it clearer*/
	cpu.register[x] = cpu.register[y] - cpu.register[x]
}

func (cpu *Chip8) opcode8xyE(x uint16, y uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x \n", x, y)

	/*Set Vx = Vx SHL 1*/
	if ((cpu.register[x] & 0x80) >> 7) == 1 {
		cpu.register[0xF] = 1
	} else {
		cpu.register[0xF] = 0
	}
	cpu.register[x] *= 2
}
