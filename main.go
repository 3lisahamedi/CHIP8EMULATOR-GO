package main

import (
	"CHIP8emulator/instruct"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"time"
)

var img *ebiten.Image
var array []byte

type Game struct{}

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("chip-8-logo.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	/*ebitenutil.DebugPrint(screen, "Hello, World!")*/
	/*ebitenutil.DebugPrint(screen, "Hello World !")
	ebitenutil.DebugPrintAt(screen, instruct.Opcode00E0(), 0, 11)*/

	/*
	* clear the screen

	 */
	/*
	* draw sprite to screen (only aligned)

	 */
	/*

		instruct.Opcode6xnn()*/
	/*	instruct.OpcodeDxyn(screen, img)
	 */
	time.Sleep(2 * time.Second)

	instruct.OpcodeDxyn(screen, img)
	fmt.Println("Drew image")

	time.Sleep(2 * time.Second)

	for _, val := range array {
		for i := range instruct.Cpu.Memory {
			instruct.Cpu.Memory[i] = val
		}
	}
	for i := range instruct.Cpu.Memory {
		instruct.Opcode(instruct.Cpu.Memory[i])
		fmt.Println("Going through memory table")
		s := fmt.Sprintf("%d is here", instruct.Cpu.Memory[i])
		fmt.Println(s)
	}
	fmt.Println("Trying to clear")
	fmt.Println("Done with clearing")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(760, 385)
	ebiten.SetWindowTitle("Hello, World!")

	array, _ = instruct.ReadFileAndReturnByteArray("1-chip8-logo.ch8")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	/*	instruct.Opcode00E0()
		instruct.Opcode6xnn()
		instruct.OpcodeAnnn()
		instruct.OpcodeDxyn()
		instruct.Opcode1nnn()*/
	fmt.Println("Goodbye, World!")
}
