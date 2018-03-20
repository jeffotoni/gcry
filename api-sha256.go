/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package gcry

import (
	"crypto/sha256"
	"fmt"
)

//
//
//
func Sha256(password string) string {

	pass := password + HASH_SALT

	h := sha256.New()
	h.Write([]byte(pass))

	return fmt.Sprintf("%x", h.Sum(nil))
}
