package grower

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	quarterMaxUnoverflowed = math.MaxInt/5*4 + 2
	waningMaxUnoverflowed  = (math.MaxInt - 768/4) / 5 * 4
)

func TestExactly(t *testing.T) {
	for length := range 1 << 16 {
		require.Equal(t, length, Exactly(length))
		require.Equal(t, -length, Exactly(-length))
	}

	require.Equal(t, math.MaxInt, Exactly(math.MaxInt))
	require.Equal(t, math.MinInt, Exactly(math.MinInt))
}

func TestQuarter(t *testing.T) {
	require.Equal(t, 0, Quarter(0))
	require.Equal(t, 2, Quarter(1))
	require.Equal(t, 3, Quarter(2))
	require.Equal(t, 4, Quarter(3))
	require.Equal(t, 5, Quarter(4))
	require.Equal(t, 6, Quarter(5))
	require.Equal(t, 7, Quarter(6))
	require.Equal(t, 8, Quarter(7))
	require.Equal(t, 10, Quarter(8))
	require.Equal(t, 11, Quarter(9))
	require.Equal(t, 12, Quarter(10))
	require.Equal(t, 13, Quarter(11))
	require.Equal(t, 15, Quarter(12))
	require.Equal(t, 16, Quarter(13))
	require.Equal(t, 17, Quarter(14))
	require.Equal(t, 18, Quarter(15))
	require.Equal(t, 20, Quarter(16))

	require.Equal(t, 1<<10+1<<8, Quarter(1<<10))
	require.Equal(t, 1<<20+1<<18, Quarter(1<<20))
	require.Equal(t, 1<<30+1<<28, Quarter(1<<30))
	require.Equal(t, math.MaxInt-1, Quarter(quarterMaxUnoverflowed-1))
	require.Equal(t, math.MaxInt, Quarter(quarterMaxUnoverflowed))

	// capacity is overflowed
	require.Equal(t, math.MaxInt, Quarter(quarterMaxUnoverflowed+1))
	require.Equal(t, math.MaxInt, Quarter(math.MaxInt))

	for length := range 1 << 16 {
		require.Equal(t, -length, Quarter(-length))
	}

	require.Equal(t, math.MinInt, Quarter(math.MinInt))
}

func TestWaning(t *testing.T) {
	require.Equal(t, 0, Waning(0))
	require.Equal(t, 8, Waning(1))
	require.Equal(t, 8, Waning(2))
	require.Equal(t, 8, Waning(3))

	// edge of tiny and small
	require.Equal(t, 8, Waning(4))

	require.Equal(t, 10, Waning(5))
	require.Equal(t, 12, Waning(6))
	require.Equal(t, 14, Waning(7))
	require.Equal(t, 16, Waning(8))
	require.Equal(t, 18, Waning(9))
	require.Equal(t, 20, Waning(10))

	// edge of small and main
	require.Equal(t, 512, Waning(256))

	require.Equal(t, 513, Waning(257))
	require.Equal(t, 514, Waning(258))
	require.Equal(t, 515, Waning(259))
	require.Equal(t, 517, Waning(260))
	require.Equal(t, 518, Waning(261))

	require.Equal(t, 1472, Waning(1<<10))
	require.Equal(t, 1310912, Waning(1<<20))
	require.Equal(t, 1342177472, Waning(1<<30))
	require.Equal(t, math.MaxInt-2, Waning(waningMaxUnoverflowed-1))
	require.Equal(t, math.MaxInt, Waning(waningMaxUnoverflowed))

	// capacity is overflowed
	require.Equal(t, math.MaxInt, Waning(waningMaxUnoverflowed+1))

	// interim is overflowed
	require.Equal(t, math.MaxInt, Waning(math.MaxInt))

	for length := range 1 << 16 {
		require.Equal(t, -length, Waning(-length))
	}

	require.Equal(t, math.MinInt, Waning(math.MinInt))
}

func BenchmarkExactly(b *testing.B) {
	// capacity and require is used to prevent compiler optimizations
	capacity := 0

	for length := range b.N {
		capacity = Exactly(length)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, capacity)
}

func BenchmarkQuarter(b *testing.B) {
	// capacity and require is used to prevent compiler optimizations
	capacity := 0

	for length := range b.N {
		capacity = Quarter(length)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, capacity)
}

func BenchmarkWaning(b *testing.B) {
	// capacity and require is used to prevent compiler optimizations
	capacity := 0

	for length := range b.N {
		capacity = Waning(length)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, capacity)
}
