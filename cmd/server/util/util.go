package util

import (
	"log"
)

func Check(err error) {
	if err != nil {
		log.Println("[Util] Error..", err)
		// panic(err)
	}
}
