package main

import (
	"fmt"
	"math/bits"
)

func main() {
	test := [4]byte{170, 187, 204, 221}
	normal := normalHash(test)
	normal = normal
	asuint := touint32(test)
	hash := festHash(asuint)
	fmt.Println(hash)
}

func normalHash(bytes [4]byte) []byte {
	r := bytes[0] * 3
	g := bytes[1] * 5
	b := bytes[2] * 7
	a := bytes[3] * 11
	hash := (r + g + b + a) % 64
	return []byte{r, g, b, a, hash}
}

func touint32(u [4]uint8) uint32 {
	return uint32(u[3])<<24 | uint32(u[2])<<16 | uint32(u[1])<<8 | uint32(u[0])
}

func festHash(u uint32) uint32 {
	const AGHash uint32 = 0x05000B00
	const BRHash uint32 = 0x03000700
	const Mask uint32 = 0x00FF00FF

	AG := u >> 8 & Mask
	BR := u & Mask
	AGHashed := AG * AGHash
	BRHashed := BR * BRHash
	added := AGHashed + BRHashed
	hash := added >> 24 & 63
	return hash
}

func festHash64(u uint32) uint32 {
	const mask uint64 = 0xFF00FF0000FF00FF
	const multiplier = 0x030007000005000B
	n := (uint64(u)<<32 | uint64(u)) & mask
	multiplied := n * multiplier
	hash := (bits.ReverseBytes64(multiplied)) & 63
	return uint32(hash)
}
