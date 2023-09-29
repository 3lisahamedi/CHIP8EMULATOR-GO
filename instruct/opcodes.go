package instruct

import (
	"fmt"
)

/*func Opcode00E0(screen *ebiten.Image) string {
	/* Clear the display
	screen.Clear()
	return "Cleared"
}
func Opcode6xnn() {
	/* Load normal register with immediate value

	fmt.Println("Load normal register")
}
func OpcodeAnnn() {
	/*  Load index register with immediate value
	fmt.Println("Load index register")
}
func OpcodeDxyn(screen *ebiten.Image, img *ebiten.Image) {
	/* Draw sprite to screen (only aligned)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(1, 1)
	screen.DrawImage(img, op)
	fmt.Println("Order to draw sprite received.")
}
func Opcode1nnn() {
	/* Jump (at the end)

	fmt.Println("Jump")
}*/

func Opcode(input uint16) { /* function that grabs the hexadecimal and associates it to corresponding cmd*/
	fmt.Println("Success. Now going to read opcode.")
	switch input {
	case 00e0:
		fmt.Println("Currently in case 00E0, clearing screen.")
		break
	}
	fmt.Println("Found no opcodes")
}
