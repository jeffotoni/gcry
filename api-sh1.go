/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package gcry

import (
	"crypto/sha1"
	"fmt"
)

//
//
//
func Sha1(key string) string {

	data := []byte(key + SHA1_SALT)
	return (fmt.Sprintf("%x", sha1.Sum(data)))
}
