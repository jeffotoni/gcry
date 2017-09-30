/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
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
