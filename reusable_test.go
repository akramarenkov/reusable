package reusable

import (
	"testing"
	"unsafe"

	"github.com/akramarenkov/reusable/grower"

	"github.com/stretchr/testify/require"
)

func TestLimited(t *testing.T) {
	buffer := New[byte](1024)

	first := buffer.Get(1024)
	require.Len(t, first, 1024)
	require.Len(t, buffer.Get(0), 1024)
	require.Equal(t, 1024, cap(first))

	second := buffer.Get(512)
	require.Len(t, second, 512)
	require.Len(t, buffer.Get(0), 512)
	require.Equal(t, 1024, cap(second))
	require.Same(t, unsafe.SliceData(first), unsafe.SliceData(second))

	third := buffer.Get(2048)
	require.Len(t, third, 1024)
	require.Len(t, buffer.Get(0), 1024)
	require.Equal(t, 1024, cap(third))
	require.Same(t, unsafe.SliceData(first), unsafe.SliceData(third))
}

func TestUnlimited(t *testing.T) {
	buffer := New[byte](0)

	first := buffer.Get(1024)
	require.Len(t, first, 1024)
	require.Len(t, buffer.Get(0), 1024)
	require.Equal(t, 1472, cap(first))

	second := buffer.Get(1472)
	require.Len(t, second, 1472)
	require.Len(t, buffer.Get(0), 1472)
	require.Equal(t, 1472, cap(second))
	require.Same(t, unsafe.SliceData(first), unsafe.SliceData(second))

	third := buffer.Get(2048)
	require.Len(t, third, 2048)
	require.Len(t, buffer.Get(0), 2048)
	require.Equal(t, 2752, cap(third))
	require.NotSame(t, unsafe.SliceData(first), unsafe.SliceData(third))

	fourth := buffer.Get(1024)
	require.Len(t, fourth, 1024)
	require.Len(t, buffer.Get(0), 1024)
	require.Equal(t, 2752, cap(fourth))
	require.Same(t, unsafe.SliceData(third), unsafe.SliceData(fourth))
}

func TestCustomGrowing(t *testing.T) {
	buffer := New[byte](0, grower.Quarter)

	first := buffer.Get(1024)
	require.Len(t, first, 1024)
	require.Len(t, buffer.Get(0), 1024)
	require.Equal(t, 1280, cap(first))

	second := buffer.Get(1280)
	require.Len(t, second, 1280)
	require.Len(t, buffer.Get(0), 1280)
	require.Equal(t, 1280, cap(second))
	require.Same(t, unsafe.SliceData(first), unsafe.SliceData(second))

	third := buffer.Get(2048)
	require.Len(t, third, 2048)
	require.Len(t, buffer.Get(0), 2048)
	require.Equal(t, 2560, cap(third))
	require.NotSame(t, unsafe.SliceData(first), unsafe.SliceData(third))

	fourth := buffer.Get(1024)
	require.Len(t, fourth, 1024)
	require.Len(t, buffer.Get(0), 1024)
	require.Equal(t, 2560, cap(fourth))
	require.Same(t, unsafe.SliceData(third), unsafe.SliceData(fourth))
}

func BenchmarkBuffer(b *testing.B) {
	buffer := New[byte](0)

	// slice and require is used to prevent compiler optimizations
	slice := buffer.Get(0)

	for range b.N {
		slice = buffer.Get(1024)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, slice)
}

func BenchmarkBufferEverIncreasing(b *testing.B) {
	buffer := New[byte](0)

	// slice and require is used to prevent compiler optimizations
	slice := buffer.Get(0)

	for length := range b.N {
		slice = buffer.Get(length + 1)
	}

	b.StopTimer()

	// meaningless check
	require.NotNil(b, slice)
}
