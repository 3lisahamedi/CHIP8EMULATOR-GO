package instruct

import (
	"fmt"
	"math/rand"
)

func (cpu *Chip8) opcodeCxnn(x uint16, n1 uint16, n2 uint16) {
	/*Set Vx = random byte AND kk. */
	fmt.Printf("COMMENCEMENT : %02x, %02x, %02x \n", x, n1, n2)
	nn := n1*16 + n2
	randInt := uint16(rand.Intn(255))
	randInt &= nn

	cpu.register[x] = byte(randInt)
}
