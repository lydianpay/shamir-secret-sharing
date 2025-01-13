package shamir

import (
	"crypto/rand"
)

// generatePolynomial constructs a random polynomial with the given intercept
func generatePolynomial(intercept, threshold uint8) ([]byte, error) {

	// Create an empty byte array to hold the polynomial
	polynomial := make([]byte, threshold)

	// Set the first byte of the polynomial as the intercept
	polynomial[0] = intercept

	// Fill the rest of the polynomial with cryptographically secure random bytes
	if _, err := rand.Read(polynomial[1:]); err != nil {
		return polynomial, err
	}

	return polynomial, nil
}

// evaluate returns the value of the polynomial for the given x
func evaluate(polynomial []byte, x uint8) uint8 {

	out := uint8(0)
	degree := len(polynomial) - 1

	// Using Horner's method, calculate the polynomial value
	for i := degree; i >= 0; i-- {
		out = multiplyGF(out, x)
		out = addGF(out, polynomial[i])
	}
	return out
}

// interpolate uses Lagrange interpolation to derive the intercepts at the lowest degree
func interpolate(xSamples, ySamples []uint8) (result uint8, err error) {

	limit := len(xSamples)
	for i := 0; i < limit; i++ {
		term := uint8(1)
		for m := 0; m < limit; m++ {
			if i != m {
				numerator := xSamples[m]
				denominator := addGF(xSamples[i], xSamples[m])
				quotient, err := divideGF(numerator, denominator)
				if err != nil {
					return 0, err
				}
				term = multiplyGF(term, quotient)
			}
		}
		group := multiplyGF(ySamples[i], term)
		result = addGF(result, group)
	}
	return result, nil
}
