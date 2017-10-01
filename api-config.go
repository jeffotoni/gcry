/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package gcry

const (
	KEY       = "AES256Key-32Characters1234567890"
	KeySize   = 32
	NonceSize = 12
	UID_SIZE  = 128
)

var (
	Nonce []byte

	SHA1_SALT = "3838x7xx2"
)
