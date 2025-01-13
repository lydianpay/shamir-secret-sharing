# Shamir's Secret Sharing Algorithm
---
<div align="center">

[![Go Report Card](https://goreportcard.com/badge/Tether-Payments/shamir-secret-sharing)](https://goreportcard.com/report/Tether-Payments/shamir-secret-sharing)
[![codecov](https://codecov.io/gh/Tether-Payments/shamir-secret-sharing/graph/badge.svg?token=TBTZIA620I)](https://codecov.io/gh/Tether-Payments/shamir-secret-sharing)
[![Maintainability](https://api.codeclimate.com/v1/badges/314cd38ef7019cac4d7b/maintainability)](https://codeclimate.com/github/Tether-Payments/shamir-secret-sharing/maintainability)
[![CodeQL](https://github.com/Tether-Payments/shamir-secret-sharing/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/Tether-Payments/shamir-secret-sharing/actions/workflows/github-code-scanning/codeql)

</div>
Written in Go ('Golang' for search engines) with zero external dependencies, this package implements
[Shamir's Secret Sharing](https://en.wikipedia.org/wiki/Shamir%27s_secret_sharing).

---

## Installation & Usage

1. Once confirming you have [Go](https://go.dev/doc/install) installed, the command below will add
`shamir` as a dependency to your Go program.
```shell
go get -u github.com/tether-payments/shamir-secret-sharing
```
2. Import the package into your code
```go
package main

import (
    "github.com/Tether-Payments/shamir-secret-sharing"
)
```
3. Create n number of shares
```go
shares, err := shamir.GenerateShares(secret, numberOfShares, threshold)
```

4. Reconstruct the original secret
```go
secret, err := shamir.Reconstruct(shares)
```

## Example

```go
package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	secret := []byte("Tether Payments Rocks!") // Replace with your secret
	numberOfShares := 6                        // Number of shares to be created
	threshold := 3                             // Minimum number of shares required to reconstruct the secret

	shares, err := shamir.GenerateShares(secret, numberOfShares, threshold)
	if err != nil {
		log.Fatal(err)
	}

	// If you want to view the shares in base64 encoding
	for idx, share := range shares {
		shareValue := base64.StdEncoding.EncodeToString(share)
		fmt.Println(fmt.Sprintf("Share %d: %s", idx, shareValue))
	}
	
	reconstructedSecret, err := shamir.Reconstruct(shares[:3])
	if err != nil {
		log.Fatal(err)
    }
	
	fmt.Println("Reconstructed Secret: ", reconstructedSecret)
}
```

## Contributing
TBD