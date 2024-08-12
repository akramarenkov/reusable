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
// Must return a zero value when the input value is zero.
//
// Must return a positive value when the input value is positive.
//
// Must be monotonically non-decreasing.
type Grower func(requested int) int

// Returns the buffer capacity exactly equal to its requested length.
func Exactly(requested int) int {
	return requested
}

// Returns the buffer capacity 25 percent greater than its requested length.
func Quarter(requested int) int {
	// First negative value will cause panic and will not lead to a chain of
	// incorrect calculations
	if requested <= 0 {
		return requested
	}

	addition := requested >> binaryPowerOfFour

	if addition == 0 {
		addition = 1
	}

	// Cannot be overflowed to a non-negative value
	newed := requested + addition

	// overflowed, maximum value is returned to satisfy the condition of monotone
	// non-decreasingness
	if newed < 0 {
		return math.MaxInt
	}

	return newed
}

// Returns the buffer capacity from 700 to 25 percent greater than its requested length.
func Waning(requested int) int {
	const (
		// Value is chosen speculatively
		tinyThreshold = 4
		// This value is selected so that the result with an input value of
		// tinyThreshold is the same in both the tiny and small versions
		tinyConjugation = 8
		// Value is chosen speculatively
		smallThreshold = 256
		// This value is selected so that the result with an input value of
		// smallThreshold is the same in both the small and main versions i.e. that
		// the condition (requested+smallConjugation)/4 == requested*2 is fulfilled with
		// requested equal to smallThreshold
		smallConjugation = 768
	)

	// First negative value will cause panic and will not lead to a chain of
	// incorrect calculations
	if requested <= 0 {
		return requested
	}

	switch {
	case requested <= tinyThreshold:
		return tinyConjugation
	case requested <= smallThreshold:
		return requested << binaryPowerOfTwo
	}

	// Cannot be overflowed to a non-negative value
	interim := (requested + smallConjugation)

	// overflowed, maximum value is returned to satisfy the condition of monotone
	// non-decreasingness
	if interim < 0 {
		return math.MaxInt
	}

	addition := interim >> binaryPowerOfFour

	// Cannot be overflowed to a non-negative value
	newed := requested + addition

	// overflowed, maximum value is returned to satisfy the condition of monotone
	// non-decreasingness
	if newed < 0 {
		return math.MaxInt
	}

	return newed
}
