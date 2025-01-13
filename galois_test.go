package shamir

import (
	"testing"
)

func TestAddGF(t *testing.T) {

	if out := addGF(41, 41); out != 0 {
		t.Fatalf("Should be 0. Got: %v", out)
	}

	if out := addGF(41, 7); out != 46 {
		t.Fatalf("Expected 46. Got: %v", out)
	}

	if out := addGF(9, 18); out != 27 {
		t.Fatalf("Expected 27. Got: %v", out)
	}
}

func TestMultiplyGF(t *testing.T) {
	if out := multiplyGF(41, 41); out != 45 {
		t.Fatalf("Expected 45. Got: %v", out)
	}

	if out := multiplyGF(7, 7); out != 21 {
		t.Fatalf("Expected 21. Got: %v", out)
	}

	if out := multiplyGF(0, 7); out != 0 {
		t.Fatalf("Expected 0. Got: %v", out)
	}

	if out := multiplyGF(7, 0); out != 0 {
		t.Fatalf("Expected 0. Got: %v", out)
	}
}

func TestDivideGF(t *testing.T) {
	aValues := []uint8{41, 41, 7, 0, 0}
	bValues := []uint8{41, 7, 41, 7, 0}

	expectedValues := []uint8{1, 102, 54, 0, 0}

	for idx, a := range aValues {
		out, err := divideGF(a, bValues[idx])
		if err != nil {
			t.Fatal(err)
		}
		if out != expectedValues[idx] {
			t.Fatalf("Expected %d. Got: %d", expectedValues[idx], out)
		}
	}

	// Testing for failure when the denominator is 0
	_, err := divideGF(7, 0)
	if err == nil {
		t.Fatal("Expected an error")
	}

}
