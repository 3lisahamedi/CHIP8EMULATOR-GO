package instruct

import (
	"fmt"
	"log"
	"os"
)

// Chip8 represents CPU CHIP-8 structure
type Chip8 struct {
	Memory   [4096]byte // 4096 octets memory
	register [16]byte   // 16 registers of 8 bits ( so one of 16 bytes )
	/*index        int          // Adress register of 16 bits*/
	I              uint16       // ONE register of 16bits ( so seeing it as a hexadecimal )
	Pc             uint16       // Program counter of 16 bits (commence a 0x200)
	stack          [16]uint16   // Stack of 16 registers of 16 bits
	sp             uint8        // Pointer of the stack
	delayTimer     byte         // Delay register of 8 bits
	soundTimer     byte         // Sound register of 8 bits
	keypad         [16]byte     // 16 keys keyboard
	Screen         [32][64]byte // 64x32 pixels screen
	screenModified bool
	clock          int // 60Hz clock

	Opcode uint16
}

var Cpu Chip8

func TransferROMToMemory(filePath string) {
	/*offset = setting to a custom point, 0x200 because default start for CHIP-8 <- but not that exactly but go equivalent ( 0xc8 )*/
	/*offset := 0x200*/

	dataFromROM, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	/*From the offset point onwards, we are copying the data one by one*/
	/*	copy(Cpu.Memory[offset:], dataFromROM)*/
	for i := 0; i < len(dataFromROM); i++ {
		Cpu.Memory[512+i] = dataFromROM[i]
	}

	/*Debug loop to see in terminal all data*/
	for i, val := range Cpu.Memory {
		if i >= 512 && i <= 712 {
			sprintf := fmt.Sprintf("%v is current value, %v is current index", val, i)
			fmt.Println(sprintf)
		}
	}
	Cpu.Pc = 0x200 /*Initialising the counter here so starts at 512 decimal placement later */
}

func (Cpu *Chip8) Transferx200ToActualOpcodes() uint16 { /*Memory has data stored. Now we grab each data, TWO by TWO because data is cut up byte by byte */
	/*Program counter is used what is used for incrementation*/
	var currentOpcode uint16

	if Cpu.Pc < 4096 {
		firstPart := Cpu.Memory[Cpu.Pc]
		/*	fmt.Printf("Program counter : %v \n", Cpu.Pc)
		 */Cpu.Pc += 1
		secondPart := Cpu.Memory[Cpu.Pc]
		/*	fmt.Printf("Program counter : %v \n", Cpu.Pc)
		 */Cpu.Pc += 1

		/*each two bytes make one opcode ( hexadecimal )*/
		currentOpcode = Cpu.WholeOpcode(firstPart, secondPart)
	}
	return currentOpcode
}
func (Cpu *Chip8) WholeOpcode(firstPart byte, secondPart byte) uint16 {
	/* Side note : Converts 8bit to 16bits, 2 bytes to 1 hexadecimal, uint8 + uint8 to uint16*/
	var wholeopcode uint16
	/* We stick a bunch of 0s behind the first part, and the second part gets added with a | ( OR )*/
	wholeopcode = (uint16(firstPart) << 8) | uint16(secondPart)
	return wholeopcode
}
