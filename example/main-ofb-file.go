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

	//
	fileCry := gcry.EncryptOFBFile("main-cbc.go")

	fmt.Println(fileCry)

	fileDecry := gcry.DecryptOFBFile(fileCry)

	fmt.Println(fileDecry)
}
