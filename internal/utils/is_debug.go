package utils

import (
	"fmt"
	"os"
)

func LogIfDebug(format string, a ...interface{}) {
	val, isSet := os.LookupEnv("VERBOSE")
	if isSet || val != "" {
		fmt.Printf(format, a)
	}
}
