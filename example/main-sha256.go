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

func main() {

	password := "1234567890#$"

	hash := gcry.Sha256(password)

	fmt.Println(hash)
}
