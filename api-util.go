/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package gcry

import (
	"os"
)

func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
