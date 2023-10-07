package instruct

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

func (cpu *Chip8) opcodeFx65(x uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", x)

	/*Fills V0 to VX with values from memory starting at address I*/
	for i := uint16(0); i <= x; i++ {
		cpu.register[i] = cpu.Memory[cpu.I+i]
	}
	cpu.I += x + 1
}

func (cpu *Chip8) opcodeFx55(x uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", x)

	/*Stores V0 to VX in memory starting at address I*/
	for i := uint16(0); i <= x; i++ {
		cpu.Memory[cpu.I+i] = cpu.register[i]
	}
	cpu.I += x + 1
}

func (cpu *Chip8) opcodeFx33(x uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", x)

	/*Stores BCD representation of Vx in memory locations I, I+1, I+2*/
	number := cpu.register[x]

	V0 := number / 100
	V1 := (number % 100) / 10
	V2 := (number % 100) % 10

	cpu.Memory[cpu.I] = V0
	cpu.Memory[cpu.I+1] = V1
	cpu.Memory[cpu.I+2] = V2
}

func (cpu *Chip8) opcodeFx1E(x uint16) {
	fmt.Printf("COMMENCEMENT : %02x \n", x)

	/*Set I = I + Vx*/
	cpu.I = cpu.I + uint16(cpu.register[x])
}

func (cpu *Chip8) OpcodeFx0A(x uint16) {
	/*Fx0A waits for a key press and returns the pressed key in vX.*/
	cpu.Keys = [16]bool{
		ebiten.IsKeyPressed(ebiten.Key1),
		ebiten.IsKeyPressed(ebiten.Key2),
		ebiten.IsKeyPressed(ebiten.Key3),
		ebiten.IsKeyPressed(ebiten.Key4),
		ebiten.IsKeyPressed(ebiten.KeyQ),
		ebiten.IsKeyPressed(ebiten.KeyW),
		ebiten.IsKeyPressed(ebiten.KeyE),
		ebiten.IsKeyPressed(ebiten.KeyR),
		ebiten.IsKeyPressed(ebiten.KeyA),
		ebiten.IsKeyPressed(ebiten.KeyS),
		ebiten.IsKeyPressed(ebiten.KeyD),
		ebiten.IsKeyPressed(ebiten.KeyF),
		ebiten.IsKeyPressed(ebiten.KeyZ),
		ebiten.IsKeyPressed(ebiten.KeyX),
		ebiten.IsKeyPressed(ebiten.KeyC),
		ebiten.IsKeyPressed(ebiten.KeyV),
	}

	secondDigit := (x & 0x0F00) >> 8

	for i := range cpu.Keys {
		if cpu.Keys[i] {
			cpu.register[secondDigit] = byte(i + 1) // TODO index
		}
	}
}
