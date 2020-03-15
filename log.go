package utilities

import (
	"log"
	"runtime"
)

func HandlePrintf(msg interface{}) {
	log.Printf("[I] %v", msg)
}

func HandleErrorPrintf(err interface{}) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[E] %v %s:%d", err, fn, line)
	}
}

func HandleWarnPrintf(msg interface{}) {
	log.Printf("[W] %v", msg)
}

func HandleErrorFatalf(err interface{}) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Fatalf("[E] %v %s:%d", err, fn, line)
	}
}
