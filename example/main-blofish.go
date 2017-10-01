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

	hashBlo := Blowfish(password)

	if CheckBlowfish(password, hashBlo) {

		fmt.Println("ok, password correct!")

	} else {

		fmt.Println("ok, password incorrect!")
	}

}
