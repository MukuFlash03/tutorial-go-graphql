package utils

import (
	"log"
)

func CheckError(err error, errorType string) {
    if err != nil {
		switch errorType {
		case "panic":
			log.Panic(err)
		case "fatal":
			log.Fatal(err)
		}
    }
}