package shamir

import (
	"fmt"
	"math/rand/v2"
)

// GenerateShares takes in a secret and splits it into n shares
// Threshold is the minimum number of shares required to reconstruct the secret
func GenerateShares(secret []byte, numberOfShares, threshold int) ([][]byte, error) {
	if numberOfShares < threshold || numberOfShares > 255 {
		return nil, fmt.Errorf("shares must be between %d and 255", threshold)
	}

	if threshold < 2 || threshold > 255 {
		return nil, fmt.Errorf("threshold must be between 2 and 255")
	}

	if len(secret) == 0 {
		return nil, fmt.Errorf("secret must not be empty")
	}

	// Generate random list of x coordinates
	xCoordinates := rand.Perm(255)

	// Initialize the share output with a random x-coordinate as the final byte
	shares := make([][]byte, numberOfShares)
	for idx := range shares {
		shares[idx] = make([]byte, len(secret)+1)
		shares[idx][len(secret)] = uint8(xCoordinates[idx]) + 1
	}

	// Iterate over each byte of the secret to generate a random polynomial for each byte
	for idx, intercept := range secret {
		polynomial, err := generatePolynomial(intercept, uint8(threshold))
		if err != nil {
			return nil, fmt.Errorf("failed to generate polynomial: %w", err)
		}

		// Using the x-coordinate for each share, compute the polynomial value
		for i := 0; i < numberOfShares; i++ {
			shares[i][idx] = evaluate(polynomial, shares[i][len(secret)]) // The x-coordinate for each share
		}
	}

	return shares, nil
}

// Reconstruct This method takes n number of shares and reconstructs the original secret
func Reconstruct(shares [][]byte) (secret []byte, err error) {
	if len(shares) < 2 {
		return nil, fmt.Errorf("miminum of two shares required to reconstruct a secret")
	}

	// Use the first share as the length to check against
	shareLength := len(shares[0])

	if shareLength < 2 {
		return nil, fmt.Errorf("shares must be a minimum of two bytes")
	}

	for i := 1; i < len(shares); i++ {
		if len(shares[i]) != shareLength {
			return nil, fmt.Errorf("shares must all be the same lenth")
		}
	}

	secret = make([]byte, shareLength-1)
	xValues := make([]uint8, len(shares))
	yValues := make([]uint8, len(shares))

	// Retrieve the x-coordinate from the end of each share
	for i, share := range shares {
		xValues[i] = share[shareLength-1]
	}

	// Reconstruct each byte of the potential secret
	for idx := range secret {
		// Retrieve the y-coordinate from each share
		for i, share := range shares {
			yValues[i] = share[idx]
		}

		secret[idx], err = interpolate(xValues, yValues)
		if err != nil {
			return nil, fmt.Errorf("failed to interpolate shares: %w", err)
		}
	}

	return secret, nil
}
