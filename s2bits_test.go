package s2bits

import (
	"strconv"
	"testing"

	"github.com/golang/geo/s2"
)

func TestCellIDLevel(t *testing.T) {
	id := s2.CellIDFromLatLng(s2.LatLngFromDegrees(40, -70))
	for l := 30; l >= 0; l-- {
		t.Run(strconv.Itoa(l), func(t *testing.T) {
			id = id.Parent(l)
			level := CellIDLevel(id)
			s2Level := id.Level()
			if level != s2Level {
				t.Errorf("expected %d, got %d", s2Level, level)
			}
		})
	}
}

func BenchmarkS2CellIDLevel(b *testing.B) {
	id := s2.CellIDFromLatLng(s2.LatLngFromDegrees(40, -70))
	for l := 30; l >= 0; l-- {
		b.Run(strconv.Itoa(l), func(b *testing.B) {
			id = id.Parent(l)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				id.Level()
			}
		})
	}
}

func BenchmarkCellIDLevel(b *testing.B) {
	id := s2.CellIDFromLatLng(s2.LatLngFromDegrees(40, -70))
	for l := 30; l >= 0; l-- {
		b.Run(strconv.Itoa(l), func(b *testing.B) {
			id = id.Parent(l)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				CellIDLevel(id)
			}
		})
	}
}
