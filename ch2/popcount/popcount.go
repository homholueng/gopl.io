package popcount

// pc[i] is the population count of i.
var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoop returns the population count (number of set bits) of x, loop edition
func PopCountLoop(x uint64) int {
	bitCount := byte(0)

	for i := uint(0); i < 8; i++ {
		bitCount += pc[byte(x>>i*8)]
	}
	return int(bitCount)
}

func PopCountMove(x uint64) int {
	bitCount := 0
	for i := uint(0); i < 64; i++ {
		bitCount += int((x >> i) & 1)
	}
	return bitCount
}

func PopCountClearLowBit(x uint64) int {
	bitCount := 0
	for ; x != 0; bitCount++ {
		x &= (x - 1)
	}
	return bitCount
}
