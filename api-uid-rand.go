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
	"math/rand"
	"time"
)

func UidRand() string {

	// generate 64 bits timestamp
	unix64bits := uint64(time.Now().UTC().UnixNano())

	buff := make([]byte, UID_SIZE)

	numRead, err := rand.Read(buff)

	if numRead != len(buff) || err != nil {
		fmt.Println(err)
	}

	unixUid := fmt.Sprintf("%x", unix64bits)

	return unixUid
}

func BlowfishUid() string {

	// generate 64 bits timestamp
	unix64bits := uint64(time.Now().UTC().UnixNano())

	buff := make([]byte, UID_SIZE)

	numRead, err := rand.Read(buff)

	if numRead != len(buff) || err != nil {
		fmt.Println(err)
	}

	unixUid := fmt.Sprintf("%x", unix64bits)

	unixUid = Sha1(Blowfish(unixUid))

	return unixUid
}
