package shadiff

import (
	"encoding/binary"

	"github.com/homholueng/gopl.io/ch2/popcount"
)

func Sha256DiffBits(s1 [32]byte, s2 [32]byte) uint {

	return uint(popcount.PopCount(binary.LittleEndian.Uint64(s1[0:8])^binary.LittleEndian.Uint64(s2[0:8])) +
		popcount.PopCount(binary.LittleEndian.Uint64(s1[8:16])^binary.LittleEndian.Uint64(s2[8:16])) +
		popcount.PopCount(binary.LittleEndian.Uint64(s1[16:24])^binary.LittleEndian.Uint64(s2[16:24])) +
		popcount.PopCount(binary.LittleEndian.Uint64(s1[24:32])^binary.LittleEndian.Uint64(s2[24:32])))

}
