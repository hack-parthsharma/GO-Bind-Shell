package gobindshell

import (
	"log"
)

func termError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}
