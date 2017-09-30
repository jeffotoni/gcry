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

	textCry := gcry.EncryptGCM("Let's encrypt our text here.")
	fmt.Println(textCry)

	textDescry := gcry.DecryptGCM(textCry)
	fmt.Println(textDescry)
}
