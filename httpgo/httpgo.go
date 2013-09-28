package httpgo

import (
	"log"
)

const Version = "0.0.1"

func LogFatal(message string) {
	log.Fatalf("[httpgo]: %v", message)
}
