# gcry

Lib that abstracts the encryption algorithms used by golang.

The goal is to further facilitate the native ualization of existing libs in Go.

We use numbers like CBC, CFB, GCM, OFB, MD5, BLOWFISH etc.

## Structure of the program

```go

- gcry
	- example
		- main-blofish-uid-rand.go
		- main-blofish.go
		- main-cbc.go
		- main-cfb.go
		- main-gcm.go
		- main-hash-file.go
		- main-sh1.go
		- main-uid-rand.go

	- api-blofish-uid-rand.go
	- api-blofish.go
	- api-cbc.go
	- api-cfb.go
	- api-gcm.go
	- api-hash-file.go
	- api-sh1.go
	- api-uid-rand.go

```

# Using EncryptGCM

```go

package main

import (
	"fmt"
	gcry "github.com/jeffotoni/gcry"
)

fsunc main() {

	textCry := gcry.EncryptGCM("Let's encrypt our text here.")
	fmt.Println(textCry)

	textDescry := gcry.DecryptGCM(textCry)
	fmt.Println(textDescry)
}

```

# Using EncryptCFB

```go

package main

import (
	"fmt"
	gcry "github.com/jeffotoni/gcry"
)

//
//
//
func main() {

	cipherText := gcry.EncryptCFB("new functional ex: of crifra CFB ")

	fmt.Println(cipherText)

	textplan := gcry.DecryptCFB(cipherText)
	fmt.Println(textplan)
}

```

# Using EncryptCBC

```go

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

# Using Uid Rand

```go

package main

import (
	"fmt"
	gcry "github.com/jeffotoni/gcry"
)

func main() {

	fmt.Println(gcry.UidRand())
}

```

# Using Sha1

```go

package main

import (
	"fmt"
	gcry "github.com/jeffotoni/gcry"
)

func main() {

	password := "1234567890#$"

	hash := gcry.Sha1(password)

	fmt.Println(hash)
}

```

# Using Blowfish

```go

package main

import (
	"fmt"
	gcry "github.com/jeffotoni/gcry"
)

func main() {

	password := "1234567890#$"

	hashBlo := gcry.Blowfish(password)

	if gcry.CheckBlowfish(password, hashBlo) {

		fmt.Println("ok, password correct!")

	} else {

		fmt.Println("ok, password incorrect!")
	}

}

```

# Using Blowfish Uid

```go

package main

import (
	"fmt"
	gcry "github.com/jeffotoni/gcry"
)

func main() {

	fmt.Println(gcry.BlowfishUid())
}

```

# Using HashFile

```go

package main

import (
	"fmt"
	gcry "github.com/jeffotoni/gcry"
)

func main() {

	hash, _ := gcry.HashFile("main-cbc.go")
	fmt.Println(hash)
}

```