package main

import "fmt"

/*
!!! main article: http://www.sunshine2k.de/articles/coding/crc/understanding_crc.html
http://programm.ws/page.php?id=663
http://ru.bmstu.wiki/%D0%97%D0%B5%D1%80%D0%BA%D0%B0%D0%BB%D1%8C%D0%BD%D1%8B%D0%B9_%D1%82%D0%B0%D0%B1%D0%BB%D0%B8%D1%87%D0%BD%D1%8B%D0%B9_%D0%B0%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC_CRC32_(asm_x86)
http://www.zlib.net/crc_v3.txt

Online check tool: https://www.ghsi.de/CRC/index.php?Polynom=100010011&Message=12345678
CRC supported list : http://protocoltool.sourceforge.net/CRC%20list.html
*/

const (
	polinom8  uint8  = 0x13
	polinom16 uint16 = 0x1021
	polinom32 uint32 = 0x12211221

	bitsInByte int = 8

	// tablesize should not be more than 256
	tablesize int = 256
)

var crc8table [tablesize]uint8
var crc16table [tablesize]uint16
var crc32table [tablesize]uint32

// --- MAIN ---

func main() {
	// Here we generating precomputed values for each byte in 0..255, check the polinom also
	crc8tableGenerator()
	crc16tableGenerator()
	crc32tableGenerator()

	// Just check that tables was generated and not emty
	fmt.Println("=== table of uint8 remainders ===")
	showGeneratedTable8(&crc8table)

	fmt.Printf("\n=== table of uint16 remainders ===\n")
	showGeneratedTable16(&crc16table)

	fmt.Printf("\n=== table of uint32 remainders ===\n")
	showGeneratedTable32(&crc32table)

	// Using crc calculation for slice of byte
	message := []uint8{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}

	fmt.Printf("\n=== crc results for provided message ===\n")
	fmt.Printf("crc8: 0x%x\n", crc8(message))
	fmt.Printf("crc16: 0x%x\n", crc16(message))
	fmt.Printf("crc32: 0x%x\n\n", crc32(message))

}

// --- END OF MAIN ---

func crc8(message []uint8) uint8 {
	var crc uint8
	for i := 0; i < len(message); i++ {
		crc ^= message[i]
		crc = crc8table[crc]
	}
	return crc
}

func crc16(message []uint8) uint16 {
	var crc uint16
	for i := 0; i < len(message); i++ {
		crc ^= uint16(message[i]) << 8
		crc = crc16table[crc>>8] ^ ((crc & 0x00FF) << 8)
	}
	return crc
}

func crc32(message []uint8) uint32 {
	var crc uint32
	for i := 0; i < len(message); i++ {
		crc ^= uint32(message[i]) << 24
		crc = crc32table[crc>>24] ^ ((crc & 0x00FFFFFF) << 8)
	}
	return crc
}

func crc8tableGenerator() {
	for y := 0; y < tablesize; y++ {
		remainder := byte(y)
		for i := 0; i < bitsInByte; i++ {
			if remainder&0x80 != 0 {
				remainder = (remainder << 1) ^ polinom8
			} else {
				remainder = (remainder << 1)
			}
		}
		crc8table[y] = remainder
	}
}

func crc16tableGenerator() {
	for y := 0; y < tablesize; y++ {
		remainder := uint16(y) << 8
		for i := 0; i < bitsInByte; i++ {
			if remainder&0x8000 != 0 {
				remainder = (remainder << 1) ^ polinom16
			} else {
				remainder = remainder << 1
			}
		}
		crc16table[y] = remainder
	}
}

func crc32tableGenerator() {
	for y := 0; y < tablesize; y++ {
		remainder := uint32(y) << 24
		for i := 0; i < bitsInByte; i++ {
			if remainder&0x80000000 != 0 {
				remainder = (remainder << 1) ^ polinom32
			} else {
				remainder = remainder << 1
			}
		}
		crc32table[y] = remainder
	}
}

func showGeneratedTable8(t *[tablesize]uint8) {
	for y := 0; y < tablesize; y++ {
		if y != tablesize-1 {
			fmt.Printf("0x%02X, ", t[y])
		} else {
			fmt.Printf("0x%02X\n", t[y])
		}
	}
}

func showGeneratedTable16(t *[tablesize]uint16) {
	for y := 0; y < tablesize; y++ {
		if y != tablesize-1 {
			fmt.Printf("0x%04X, ", t[y])
		} else {
			fmt.Printf("0x%04X\n", t[y])
		}
	}
}

func showGeneratedTable32(t *[tablesize]uint32) {
	for y := 0; y < tablesize; y++ {
		if y != tablesize-1 {
			fmt.Printf("0x%X, ", t[y])
		} else {
			fmt.Printf("0x%X\n", t[y])
		}
	}
}
