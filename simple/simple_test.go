package simple

import (
	"testing"
	"unsafe"

	"github.com/akramarenkov/reusable/grower"
	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	buffer := New[byte](1024, nil)
	require.Len(t, buffer.Get(0), 1024)
	require.Equal(t, 1024, cap(buffer.Get(0)))

	old := buffer.Get(0)

	_ = buffer.Get(1024)
	require.Len(t, buffer.Get(0), 1024)
	require.Equal(t, 1024, cap(buffer.Get(0)))
	require.Same(t, unsafe.SliceData(old), unsafe.SliceData(buffer.Get(0)))

	old = buffer.Get(0)

	_ = buffer.Get(2048)
	require.Len(t, buffer.Get(0), 2048)
	require.Equal(t, 2048, cap(buffer.Get(0)))
	require.NotSame(t, unsafe.SliceData(old), unsafe.SliceData(buffer.Get(0)))

	old = buffer.Get(0)

	_ = buffer.Get(512)
	require.Len(t, buffer.Get(0), 512)
	require.Equal(t, 2048, cap(buffer.Get(0)))
	require.Same(t, unsafe.SliceData(old), unsafe.SliceData(buffer.Get(0)))

	old = buffer.Get(0)

	buffer.Reset()
	require.Len(t, buffer.Get(0), 1024)
	require.Equal(t, 1024, cap(buffer.Get(0)))
	require.NotSame(t, unsafe.SliceData(old), unsafe.SliceData(buffer.Get(0)))

	old = buffer.Get(0)

	buffer.Reset()
	require.Len(t, buffer.Get(0), 1024)
	require.Equal(t, 1024, cap(buffer.Get(0)))
	require.Same(t, unsafe.SliceData(old), unsafe.SliceData(buffer.Get(0)))
}

func BenchmarkSimple(b *testing.B) {
	buffer := New[byte](0, grower.Waning)

	// slice and require is used to prevent compiler optimizations
	slice := buffer.Get(0)

	for range b.N {
		slice = buffer.Get(1024)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, slice)
}

func BenchmarkSimpleEverIncreasing(b *testing.B) {
	buffer := New[byte](0, grower.Waning)

	// slice and require is used to prevent compiler optimizations
	slice := buffer.Get(0)

	for length := range b.N {
		slice = buffer.Get(length)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, slice)
}
