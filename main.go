package main

import (
	"CHIP8emulator/instruct"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Game struct {
}

func init() {
	var err error
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {
	instruct.Opcode(instruct.Cpu.Transferx200ToActualOpcodes())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var screenBool [64][32]bool
	for Vx, width := range instruct.Cpu.Screen { /*ranges over the rows*/
		for Vy := range width { /*ranges over the columns*/
			if !screenBool[Vx][Vy] {
				screen.Set(Vx, Vy, color.Color(color.RGBA{R: 153, G: 102, B: 1})) /*Sets main color*/
			} else {
				screen.Set(Vx, Vy, color.Color(color.RGBA{R: 255, G: 204, B: 1})) /*Sets img color*/
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64, 32
}

func main() {
	ebiten.SetWindowSize(760, 385)
	ebiten.SetWindowTitle("Hello, World!")

	instruct.TransferROMToMemory("1-chip8-logo.ch8")
	instruct.Opcode(instruct.Cpu.Transferx200ToActualOpcodes())

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Goodbye, World!")
}
