package instruct

import (
	"fmt"
	"log"
	"os"
)

// Chip8 represents CPU CHIP-8 structure
type Chip8 struct {
	Memory     [4096]byte    // 4096 octets memory
	registers  [16]byte      // 16 registers of 8 bits
	index      uint16        // Adress register of 16 bits
	pc         uint16        // Program counter of 16 bits
	stack      [16]uint16    // Stack of 16 registers of 16 bits
	sp         uint8         // Pointer of the stack
	delayTimer byte          // Delay register of 8 bits
	soundTimer byte          // Sound register of 8 bits
	keypad     [16]byte      // 16 keys keyboard
	screen     [64 * 32]byte // 64x32 pixels screen
	clock      int           // 60Hz clock
}

var Cpu Chip8

func ReadFileAndReturnByteArray(extractedFilePath string) ([]byte, error) {
	file, err := os.Open("1-chip8-logo.ch8")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return buffer, nil
}
