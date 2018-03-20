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

	hash, _ := gcry.HashFile("main-cbc.go")
	fmt.Println(hash)
}
