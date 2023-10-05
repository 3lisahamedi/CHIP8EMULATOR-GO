package instruct

import "fmt"

func (cpu *Chip8) opcode6xnn(x uint16, n1 uint16, n2 uint16) {
	fmt.Printf("COMMENCEMENT : %02x, %02x, %02x \n", x, n1, n2)
	nn := n1*16 + n2 /*Hexadecimal to decimal calculation*/

	/*Set Vx = nn*/
	cpu.register[x] = byte(nn)
}
