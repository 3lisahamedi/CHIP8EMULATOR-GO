package instruct

import (
	"fmt"
	"log"
	"os"
)

// Chip8 represents CPU CHIP-8 structure
type Chip8 struct {
	memory     [4096]byte    // 4096 octets memory
	registers  [16]byte      // 16 registers of 8 bits
	index      uint16        // Adress register of 16 bits
	pc         uint16        // Program counter of 16 bits (commence a 0x200)
	stack      [16]uint16    // Stack of 16 registers of 16 bits
	sp         uint8         // Pointer of the stack
	delayTimer byte          // Delay register of 8 bits
	soundTimer byte          // Sound register of 8 bits
	keypad     [16]byte      // 16 keys keyboard
	screen     [64 * 32]byte // 64x32 pixels screen
	clock      int           // 60Hz clock
}

var Cpu Chip8

// offset = setting to a custom point, 0x200 because default start for CHIP-8 <- but not that exactly but go equivalent ( 0xc8 )

func TransferROMToMemory(filePath string) {
	offset := 0xc8
	dataFromROM, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	copy(Cpu.memory[offset:], dataFromROM)
	fmt.Println("Transfer complete.")
	fmt.Println("Checking data :")

	for i, val := range Cpu.memory {
		sprintf := fmt.Sprintf("%v is current value, %v is current index", val, i)
		fmt.Println(sprintf)
	}
}

/*func Transferx200ToRegister() {
	for i, val := range Cpu.memory[] {
		while i > 200 <= 459  {
			for y := 1; i%2 == 0; y++ {
				Cpu.registers[y] = append(Cpu.registers[y], string(val))
			}
		}
	}
}*/
