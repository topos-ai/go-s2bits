package s2bits

import (
	"math/bits"

	"github.com/golang/geo/s2"
)

func CellIDLevel(id s2.CellID) int {
	return 30 - bits.TrailingZeros64(uint64(id))>>1
}
