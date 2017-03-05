package main

import "fmt"

const (
	polinom byte = 0xb5
)

// --- MAIN ---

func main() {
	var b byte = 0xd6

	fmt.Printf("slow: %08b\n", crcSlow(b))
	fmt.Printf("slow: %08b\n", crcNaive(b))

	fmt.Printf("0xd6: %08b\n", 0xd6)
	fmt.Printf("0xd6: %08b\n", 0xb5)

}

func crcSlow(message byte) byte {
	//var crc byte = 0xFF

	for i := 0; i < 7; i++ {
		if message&0x80 == 0x80 {
			message ^= polinom
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

/*
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
