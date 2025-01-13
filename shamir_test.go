package shamir

import (
	"bytes"
	"math/rand/v2"
	"testing"
)

func TestGenerateShares(t *testing.T) {
	secret := []byte("Yo Cuz!")
	numberOfShares := []int{4, 6, 8}
	thresholds := []int{2, 3, 6}

	for idx, cnt := range numberOfShares {
		shares, err := GenerateShares(secret, cnt, thresholds[idx])
		if err != nil {
			t.Error(err)
		}

		// Check for the correct number of shares
		if len(shares) != cnt {
			t.Errorf("Expected %d shares, got %d", cnt, len(shares))
		}

		// Check the length of each share for correct length
		for _, share := range shares {
			if len(share) != len(secret)+1 {
				t.Errorf("Expected %d shares, got %d", len(secret)+1, len(shares))
			}
		}
	}

	// Negative Tests
	_, err := GenerateShares([]byte{}, 6, 3)
	if err == nil {
		t.Error("Expected error when using an empty secret")
	}

	_, err = GenerateShares(secret, 4, 6)
	if err == nil {
		t.Error("Expected error when calculating fewer shares than the threshold")
	}

	_, err = GenerateShares(secret, 381, 6)
	if err == nil {
		t.Error("Expected error when calculating more than 255 shares")
	}

	_, err = GenerateShares(secret, 6, 1)
	if err == nil {
		t.Error("Expected error when using a threshold less than 2")
	}

	_, err = GenerateShares(secret, 6, 381)
	if err == nil {
		t.Error("Expected error when using a threshold greater than 255")
	}
}

func TestReconstructShares(t *testing.T) {

	secrets := [][]byte{
		[]byte("Yo Cuz!"),
		[]byte("1632c1ee-8a67-40c3-92a2-8404f100e15f"),
		[]byte("mmdjphbfkdjkxpowjumccoyivysdfkxsgzbvdlxfllzvszggfozlwsyryggastbpvxdmetsazvtapyferrodlerplldmdyccivrx" +
			"wtxmvlivnuaniqyrfykdwzebrflqixgdmpzgdmuxmdwopvvxunjxdbwxizhkpuudamugbglwyxdfdlpyjxhuraolmrpafvivinthnaz" +
			"mwarajcvxlwqptwrrfpoxcuynrukymsmbcovjtdongfyhlzzdxuusgkgaourfaysvmlgmusvcmpmyclbaccnrgtgmdcmoomfhajmdsn" +
			"wepalqemmddywlviolhobzsndwfwpijwuldgeedwvyoxamtjbgbkxeuvrvdkhaolynrvpbyxmdvzvbguqlmanovgzlsvokgzcpabmsa" +
			"ikzfofmgmvnfekacstqepwrnnrzlkmumopcmuygukvyamuvgouzktfvvgqpzjbchhyeokafdvjjwzbjoxffoqwjstmjwzcnsvfxxehe"),
	}
	numberOfShares := []int{4, 6, 8}
	thresholds := []int{2, 3, 6}

	for idx, secret := range secrets {
		shares, err := GenerateShares(secret, numberOfShares[idx], thresholds[idx])
		if err != nil {
			t.Error(err)
		}

		// Randomize the shares
		rand.Shuffle(len(shares), func(i, j int) {
			shares[i], shares[j] = shares[j], shares[i]
		})

		reconstructedSecret, err := Reconstruct(shares[:numberOfShares[idx]])
		if err != nil {
			t.Error(err)
		}

		if bytes.Compare(reconstructedSecret, secret) != 0 {
			t.Errorf("Expected secret %s, got %s", secret, reconstructedSecret)
		}
	}

	// Negative Tests
	badShares := [][]byte{{3}}
	_, err := Reconstruct(badShares)
	if err == nil {
		t.Error("Expected error when passing fewer than 2 shares")
	}

	badShares = [][]byte{{3}, {4}, {5}}
	_, err = Reconstruct(badShares)
	if err == nil {
		t.Error("Expected error when share length is less than 2")
	}

	badShares = [][]byte{{1, 2, 3, 4}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5, 6}}
	_, err = Reconstruct(badShares)
	if err == nil {
		t.Error("Expected error when share length is not all the same")
	}

	badShares = [][]byte{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	_, err = Reconstruct(badShares)
	if err == nil {
		t.Error("Expected error when shares collide (force denominator to be 0 for interpolation)")
	}
}
