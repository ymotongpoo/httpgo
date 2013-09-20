package httpgo

import (
	"fmt"
	"log"
)

func LogFatal(message string) {
	log.Fatalf("[httpgo]: %v", message)
}
