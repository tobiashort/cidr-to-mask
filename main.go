package main

import (
	"fmt"
	"os"
	"strconv"
)

func printUsage() {
	fmt.Println("Usage: cidr-to-mask CIDR")
	os.Exit(-1)
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		return
	}
	cidr, err := strconv.Atoi(os.Args[1])
	if err != nil || cidr < 0 || cidr > 32 {
		printUsage()
		return
	}
	mask := uint32(0b11111111_11111111_11111111_11111111)
	mask = (mask << (32 - cidr)) & mask
	mask3 := uint8(mask >> 24)
	mask2 := uint8(mask >> 16)
	mask1 := uint8(mask >> 8)
	mask0 := uint8(mask)
	fmt.Printf("%d.%d.%d.%d\n", mask3, mask2, mask1, mask0)
}
