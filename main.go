package main

import (
	"CHIP8emulator/instruct"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Game struct {
	cpu instruct.Chip8
}

func init() {
	var err error
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {
	for i := 0; i < len(g.cpu.Memory)-1; i++ {
		g.cpu.Opcode = instruct.Cpu.Transferx200ToActualOpcodes()
		instruct.SplitOpcode(g.cpu.Opcode)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for Vy, width := range instruct.Cpu.Screen { /*ranges over the rows*/
		for Vx, screenBool := range width { /*ranges over the columns*/
			if screenBool == 1 {
				screen.Set(Vx, Vy, color.RGBA{R: 255, G: 204, B: 1}) /*Sets img color*/
			} else {
				screen.Set(Vx, Vy, color.RGBA{R: 153, G: 102, B: 1}) /*Sets main color*/
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}

func main() {
	ebiten.SetWindowSize(640, 320) /*64 x 32*/
	ebiten.SetWindowTitle("Hello, World!")

	instruct.TransferROMToMemory("1-chip8-logo.ch8")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Goodbye, World!")
}
