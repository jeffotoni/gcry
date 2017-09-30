# gcry

Lib that abstracts the encryption algorithms used by golang.

The goal is to further facilitate the native ualization of existing libs in Go.

We use numbers like CBC, CFB, GCM, OFB, MD5, BLOWFISH etc.

## Structure of the program

```go

- gcry
	- example
		- main-cbc.go
		- main-cbc.go
		- main-cbc.go
		- main-cbc.go
		- main-cbc.go
		- main-cbc.go

 - api-cbc.go
 - api-cbc.go
 - api-cbc.go
 - api-cbc.go
 - api-cbc.go
 - api-cbc.go


```

# Using EncryptCBC

```go
/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"fmt"
	gcry "github.com/jeffotoni/gcry"
)

//
//
//
func main() {

	// multiple 16 len text
	cipherText := gcry.EncryptCBC("new functional ex: of crifra CBC")

	fmt.Println(cipherText)

	textplan := gcry.DecryptCBC(cipherText)

	fmt.Println(textplan)
}

```

# Using EncryptCBC

```go

func main() {

	textCry := gcry.EncryptGCM("Let's encrypt our text here.")
	fmt.Println(textCry)

	textDescry := gcry.DecryptGCM(textCry)
	fmt.Println(textDescry)
}

```