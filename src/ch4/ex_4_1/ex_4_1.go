package ex_4_1

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	var b1, b2 [sha256.Size]byte
	b1 = sha256.Sum256([]byte("x"))
	b2 = sha256.Sum256([]byte("X"))
	fmt.Printf("There is %d different bits in %x and %x", checkSHA256(&b1, &b2), b1, b2)
}

func checkSHA256(b1, b2 *[sha256.Size]byte) int {
	count := 0
	for i := 0; i < sha256.Size; i++ {
		b := b1[i] ^ b2[i]
		count += countBits(b)
	}
	return count
}

func testCountBits() {
	var b byte
	for b = 0; b <= 0xf; b++ {
		fmt.Printf("The number of bit in %b is %d\n", b, countBits(b))
	}
}

// count the number of bit '1' in byte and return it
func countBits(b byte) int {
	count := 0
	for b != 0 {
		b = b & (b - 1)
		count++
	}
	return count
}
