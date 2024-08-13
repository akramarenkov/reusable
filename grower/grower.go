// Several functions are implemented here that calculate the buffer capacity based
// on the requested buffer length.
package grower

import "math"

const (
	binaryPowerOfFour = 2
	binaryPowerOfTwo  = 1
)

// Calculates the buffer capacity based on the requested buffer length.
//
// Must return a zero capacity when the requested length is zero.
//
// Calculated capacity cannot be less than the requested length.
//
// Dependence of the calculated capacity on the requested length must be monotonically
// non-decreasing.
type Grower func(length int) int

// Returns the capacity exactly equal to its requested length.
func Exactly(length int) int {
	return length
}

// Returns the capacity 25 percent greater than its requested length.
func Quarter(length int) int {
	// First negative value will cause panic and will not lead to a chain of
	// incorrect calculations
	if length <= 0 {
		return length
	}

	addition := length >> binaryPowerOfFour

	if addition == 0 {
		addition = 1
	}

	// Cannot be overflowed to a non-negative value
	capacity := length + addition

	// overflowed, maximum value is returned to satisfy the condition of monotone
	// non-decreasingness
	if capacity < 0 {
		return math.MaxInt
	}

	return capacity
}

// Returns the capacity from 700 to 25 percent greater than its requested length.
func Waning(length int) int {
	const (
		// Value is chosen speculatively
		tinyThreshold = 4
		// This value is selected so that the capacity with an requested length equal to
		// tinyThreshold is the same in both the tiny and small versions
		tinyConjugation = 8
		// Value is chosen speculatively
		smallThreshold = 256
		// This value is selected so that the capacity with an requested length equal to
		// smallThreshold is the same in both the small and main versions i.e. that
		// the condition length+(length+smallConjugation)/4 == 2*length is
		// fulfilled with requested length equal to smallThreshold
		smallConjugation = 768
	)

	// First negative value will cause panic and will not lead to a chain of
	// incorrect calculations
	if length <= 0 {
		return length
	}

	switch {
	case length <= tinyThreshold:
		return tinyConjugation
	case length <= smallThreshold:
		return length << binaryPowerOfTwo
	}

	// Cannot be overflowed to a non-negative value
	interim := (length + smallConjugation)

	// overflowed, maximum value is returned to satisfy the condition of monotone
	// non-decreasingness
	if interim < 0 {
		return math.MaxInt
	}

	addition := interim >> binaryPowerOfFour

	// Cannot be overflowed to a non-negative value
	capacity := length + addition

	// overflowed, maximum value is returned to satisfy the condition of monotone
	// non-decreasingness
	if capacity < 0 {
		return math.MaxInt
	}

	return capacity
}
