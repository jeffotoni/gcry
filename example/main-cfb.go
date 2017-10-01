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

	cipherText := gcry.EncryptCFB("new functional ex: of crifra CBC ")

	fmt.Println(cipherText)

	textplan := gcry.DecryptCFB(cipherText)
	fmt.Println(textplan)
}
