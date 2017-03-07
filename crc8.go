package main

import "fmt"

const (
	polinom8 byte = 0x13
)

var crc8table [256]byte

// --- MAIN ---

func main() {
	b := []byte{0x12, 0x34, 0x56, 0x78}

	crcInit()

	fmt.Printf("crc: 0x%x\n", crc8(b))
	fmt.Printf("crc fast %08b\n", crc8fast(b))
}

func crc8(b []byte) byte {

	var remainder byte = 0x00

	for y := 0; y < len(b); y++ {
		remainder ^= b[y]

		for i := 8; i != 0; i-- {
			if remainder&0x80 != 0 {
				remainder = (remainder << 1) ^ polinom8
			} else {
				remainder = remainder << 1
			}
		}

	}
	return remainder

}

func crc8fast(b []byte) byte {
	var crc byte = 0x00

	for i := 0; i < len(b); i++ {
		tmp := crc ^ b[i]
		crc = crc8table[tmp]
	}

	return crc
}

func crcInit() {
	var currentByte byte
    // 'y' - cause bYte 
    // 'i' - cause bIy  :)
	for y := 0; y < len(crc8table); y++ {

		currentByte = byte(y)

		for i := 0; i < 8; i++ {
			if currentByte&0x80 != 0 {
				currentByte = (currentByte << 1) ^ polinom8
			} else {
				currentByte = currentByte << 1
			}
		}
		crc8table[y] = currentByte
	}
}

func printByteAsHex(b [256]byte) {
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d) 0x%x\n ", i, b[i])
	}
}

