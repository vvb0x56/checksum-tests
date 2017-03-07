package main
/*
// http://www.sunshine2k.de/articles/coding/crc/understanding_crc.html
// http://programm.ws/page.php?id=663
// http://ru.bmstu.wiki/%D0%97%D0%B5%D1%80%D0%BA%D0%B0%D0%BB%D1%8C%D0%BD%D1%8B%D0%B9_%D1%82%D0%B0%D0%B1%D0%BB%D0%B8%D1%87%D0%BD%D1%8B%D0%B9_%D0%B0%D0%BB%D0%B3%D0%BE%D1%80%D0%B8%D1%82%D0%BC_CRC32_(asm_x86)
// http://www.zlib.net/crc_v3.txt

// Online check tool: https://www.ghsi.de/CRC/index.php?Polynom=100010011&Message=12345678

CRC supported list : http://protocoltool.sourceforge.net/CRC%20list.html

*/

import "fmt"

const (
	polinom8 byte = 0x13
)

var crc8table [256]byte

// --- MAIN ---

func main() {
	b := []byte{0x12, 0x34, 0x56, 0x78}

	crcInit()

	//printByteAsHex(crc8table)

	//printByteAsHex(b)
	fmt.Printf("crc: 0x%x\n", crc8(b))
	fmt.Printf("crc fast %08b\n", crc8fast(b))
	//fmt.Printf("0x%x\n", crc8table[0x1f])
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

