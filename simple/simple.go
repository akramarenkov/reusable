// Simple temporary buffer of variable length. Does not automatically recreate the
// buffer to reduce memory usage.
package simple

import (
	"github.com/akramarenkov/reusable/grower"
)

// Simple temporary buffer of variable length.
//
// Does not automatically recreate the buffer to reduce memory usage.
type Buffer[Type any] struct {
	grower  grower.Grower
	initial int

	slice []Type
}

// Creates simple buffer with specified initial length and growing function.
//
// If the growing function is not specified, the grower.Exactly function will be used.
func New[Type any](initial int, growing grower.Grower) *Buffer[Type] {
	bfr := &Buffer[Type]{
		grower:  growing,
		initial: initial,
	}

	if bfr.grower == nil {
		bfr.grower = grower.Exactly
	}

	bfr.remake(bfr.initial)

	return bfr
}

// Increases/decrease buffer length to specified value and returns it.
//
// If a length of zero is specified, the current buffer will simply be returned without
// any additional processing.
//
// Data in the buffer may and most likely will not persist between calls with a non-zero
// length.
func (bfr *Buffer[Type]) Get(length int) []Type {
	if length == 0 {
		return bfr.slice
	}

	if length <= cap(bfr.slice) {
		bfr.slice = bfr.slice[:length]
		return bfr.slice
	}

	// There is no need to save data between calls
	bfr.remake(length)

	return bfr.slice
}

// Recreates the buffer with the initial length if the capacity of the current buffer
// is greater than the initial length. Used to free up memory from an overgrown buffer.
func (bfr *Buffer[Type]) Reset() {
	if cap(bfr.slice) > bfr.grower(bfr.initial) {
		bfr.remake(bfr.initial)
	}
}

func (bfr *Buffer[Type]) remake(length int) {
	bfr.slice = make([]Type, bfr.grower(length))
}
