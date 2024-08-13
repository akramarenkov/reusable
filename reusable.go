// Reusable temporary buffer of variable length.
package reusable

import (
	"github.com/akramarenkov/reusable/grower"
)

// Reusable temporary buffer of variable length.
type Buffer[Type any] struct {
	grower grower.Grower
	limit  int

	slice []Type
}

// Creates buffer with specified limit of length and growing function.
//
// If the limit is not zero, it will be impossible to get a buffer longer than this.
//
// If the limit is zero, the buffer size is unlimited.
//
// If the limit is not zero, then there is no need to specify the growing function,
// since grower.Exactly will be used forcibly.
//
// If limit is zero and the growing function is not specified, the grower.Waning
// function will be used.
func New[Type any](limit int, growing ...grower.Grower) *Buffer[Type] {
	bfr := &Buffer[Type]{
		limit: limit,
	}

	if bfr.limit != 0 {
		bfr.grower = grower.Exactly
		bfr.remake(limit)

		return bfr
	}

	for _, grower := range growing {
		if grower != nil {
			bfr.grower = grower
			return bfr
		}
	}

	bfr.grower = grower.Waning

	return bfr
}

// Increases/decrease buffer length to specified value and returns it.
//
// Data in the buffer may and most likely will not persist between calls.
func (bfr *Buffer[Type]) Get(length int) []Type {
	if length <= cap(bfr.slice) {
		return bfr.slice[:length]
	}

	if bfr.limit != 0 {
		return bfr.slice
	}

	// There is no need to save data between calls
	bfr.remake(length)

	return bfr.slice
}

func (bfr *Buffer[Type]) remake(length int) {
	bfr.slice = make([]Type, length, bfr.grower(length))
}
