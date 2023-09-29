package main

import (
	"CHIP8emulator/instruct"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct{}

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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(760, 385)
	ebiten.SetWindowTitle("Hello, World!")

	instruct.TransferROMToMemory("1-chip8-logo.ch8")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Goodbye, World!")
}
