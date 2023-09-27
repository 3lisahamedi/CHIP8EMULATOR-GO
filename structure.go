package main

// Chip8 représente la structure du CPU CHIP-8
type Chip8 struct {
	memory     [4096]byte   // Mémoire de 4096 octets
	registers  [16]byte     // 16 registres de 8 bits
	index      uint16       // Registre d'adresse de 16 bits
	pc         uint16       // Registre de programme de 16 bits
	stack      [16]uint16   // Pile de 16 registres de 16 bits
	sp         uint8        // Pointeur de pile
	delayTimer byte         // Registre de délai de 8 bits
	soundTimer byte         // Registre de son de 8 bits
	keypad     [16]byte     // Clavier de 16 touches
	screen     [64][32]byte // Écran de 64x32 pixels
	clock      int          // Horloge de 60Hz
}
