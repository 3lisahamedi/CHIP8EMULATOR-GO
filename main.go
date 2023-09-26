package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

var img *ebiten.Image
var memory int = 4096

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("chip-8-logo.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	/*ebitenutil.DebugPrint(screen, "Hello, World!")*/
	/*ebitenutil.DebugPrint(screen, "Hello World !")
	ebitenutil.DebugPrintAt(screen, instruct.Opcode00E0(), 0, 11)*/

	/*
		* clear the screen
		instruct.Opcode00E0(screen)
	*/
	/*
		* draw sprite to screen (only aligned)
		instruct.OpcodeDxyn(screen, img)
	*/
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(760, 385)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	/*	instruct.Opcode00E0()
		instruct.Opcode6xnn()
		instruct.OpcodeAnnn()
		instruct.OpcodeDxyn()
		instruct.Opcode1nnn()*/
	fmt.Println("Hello, World!")
}
