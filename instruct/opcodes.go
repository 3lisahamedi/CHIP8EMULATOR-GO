package instruct

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

func Opcode00E0(screen *ebiten.Image) string {
	/* Clear the display */
	screen.Clear()
	return "Cleared"
}
func Opcode6xnn() {
	/* Load normal register with immediate value */
	fmt.Println("Load normal register")
}
func OpcodeAnnn() {
	/*  Load index register with immediate value */
	fmt.Println("Load index register")
}
func OpcodeDxyn() {
	/* Draw sprite to screen (only aligned) */
	fmt.Println("Draw sprite")
}
func Opcode1nnn() {
	/* Jump (at the end) */
	fmt.Println("Jump")
}
