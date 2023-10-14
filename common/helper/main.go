package helper

import (
	"log"
	"os"
)

func CheckError(e error, msg string, stop ...bool) {
	if e != nil {
		log.Println("Error:", e, "\n", msg)
		if stop[0] {
			os.Exit(1)
		}
	}

}
