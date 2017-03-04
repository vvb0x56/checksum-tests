package main

import "fmt"

type crc byte

const (
	polinomial byte = 0x31
)

// --- MAIN ---

func main() {

	bf := []byte("test")
	/*
		fmt.Println(reflect.TypeOf(polinomial))
		fmt.Printf("%08b\n", polinomial)
		fmt.Printf("%08b\n", bf)
	*/
	fmt.Printf("Poly 0x31: [%08b]\n", 0x31)
	fmt.Printf("\"test\": %08b\n", bf)
	crcResult := crcSlow(bf)
	fmt.Printf("CRC result: %d [%08b]\n", crcResult, crcResult)
}

// ------

/* --- EXAMPLE FROM WIKI ---

11010011101100 000 <--- input right padded by 3 bits
1011               <--- divisor
01100011101100 000 <--- result (note the first four bits are the XOR with the divisor beneath, the rest of the bits are unchanged)
 1011              <--- divisor ...
00111011101100 000
  1011
00010111101100 000
   1011
00000001101100 000 <--- note that the divisor moves over to align with the next 1 in the dividend (since quotient for that step was zero)
       1011             (in other words, it doesn't necessarily move one bit per iteration)
00000000110100 000
        1011
00000000011000 000
         1011
00000000001110 000
          1011
00000000000101 000
           101 1
-----------------
00000000000000 100 <--- remainder (3 bits).  Division algorithm stops here as dividend is equal to zero.

*/

func crcNaive(message byte) byte {
	remainder := message

	for i := 8; i != 0; i-- {
		if remainder&0x80 == 0x80 {
			remainder ^= polinomial
		}

		remainder = (remainder << 1)
	}
	return remainder >> 4
}

/*
"test" = [01110100 01100101 01110011 01110100]



*/

func crcSlow(message []byte) byte {
	var remainder byte = 0xFF
	fmt.Printf("remainder start: [%08b]\n", remainder)

	for by := 0; by < len(message); by++ {
		fmt.Printf("remainder and %d byte of the message: [%08b] -- [%08b]\n", by, remainder, message[by])

		remainder ^= message[by]

		fmt.Printf("remainder xor %d byte: [%08b]\n", by, remainder)

		for i := 8; i > 0; i-- {
			if remainder&0x80 == 0x80 {
				remainder = (remainder << 1) ^ polinomial
			} else {
				remainder = (remainder << 1)
			}
			fmt.Printf("remainder in %d bit: [%08b]\n", i, remainder)
			fmt.Printf("polynomial          [%08b]\n", polinomial)
		}

	}
	return remainder

}
