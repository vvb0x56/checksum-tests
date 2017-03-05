package main

// http://www.sunshine2k.de/articles/coding/crc/understanding_crc.html

import "fmt"

const (
	polinom byte = 0x1d
)

var crc8table [256]byte

// --- MAIN ---

func main() {
	b := []byte{0x01, 0x02}

	crc8init()

	//printByteAsHex(crc8table)

	//printByteAsHex(b)
	fmt.Printf("crc: 0x%x\n", crc8(b))
	fmt.Printf("crc fast 0x%x\n", crc8fast(b))
	fmt.Printf("0x%x\n", crc8table[0x1f])
}

func crc8(b []byte) byte {

	var remainder byte = 0x00

	for y := 0; y < len(b); y++ {
		remainder ^= b[y]

		for i := 8; i != 0; i-- {
			if remainder&0x80 != 0 {
				remainder = (remainder << 1) ^ polinom
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

func crc8init() {
	//var y byte
	var currentByte byte
	for y := 0; y < len(crc8table); y++ {

		currentByte = byte(y)

		for i := 0; i < 8; i++ {
			if currentByte&0x80 != 0 {
				currentByte = (currentByte << 1) ^ polinom
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

/*

func crcSlow(message byte) byte {
	//var crc byte = 0xFF

	for i := 0; i < 8; i++ {
		if message&0x80 != 0 {
			message = message ^ polinom
		}
		message = (message << 1)

	}

	return message
}

func crcNaive(message byte) byte {
	remainder := message

	for i := 7; i != 0; i-- {
		if remainder&0x80 == 0x80 {
			remainder ^= polinom
		}

		remainder = (remainder << 1)
	}
	return remainder
}


func crcSlow(message []byte) byte {
	var remainder byte = 0xFF

	for by := 0; by < len(message); by++ {

		remainder ^= message[by]

		fmt.Printf("remainder: %08b\n", remainder)

		for i := 8; i > 0; i-- {
			if remainder&0x80 == 0x80 {
				remainder = (remainder << 1) ^ polinom
			} else {
				remainder = (remainder << 1)
			}
		}
	}
	return remainder

}
*/
