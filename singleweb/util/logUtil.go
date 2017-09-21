package util

import (
	"flag"
	"os"
	"fmt"
	"log"
)

var logFileName = flag.String("log", "webservice.log", "Log file name")
var logFile, logErr = os.OpenFile(*logFileName, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)

func init() {
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile )
}

