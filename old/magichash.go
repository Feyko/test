package main

import "fmt"

func main() {
	var OP uint8 = 0b_01_11_00_10
	val := extract(OP)
	fmt.Println(val)
}

func extract(op uint8) [4]uint8 {
	var valueMask uint8 = 0b00111111
	var expandMask uint32 = 0x03030300

	val := uint32(op & valueMask)
	var multiplier uint32 = 1 + 1<<8 + 1<<14 + 1<<20
	expanded := val * multiplier
	val = expanded & expandMask
	return uint32Bytes(val)
}

func uint32Bytes(u uint32) [4]uint8 {
	return [4]uint8{uint8((u >> 24) & 255), uint8((u >> 16) & 255), uint8((u >> 8) & 255), uint8(u & 255)}
}
