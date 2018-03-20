/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package gcry

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Blowfish(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {

		fmt.Println(err)
	}

	return string(bytes)
}

//
//
//
func CheckBlowfish(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}
