package utils

import (
	"fmt"
	"log"
	"os"
)

var (
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
)

func InitCustomLogger(name string) {
	InfoLog = log.New(os.Stdout, fmt.Sprintf("[%v-INFO] ", name), log.Ldate|log.Ltime)
	WarningLog = log.New(os.Stdout, fmt.Sprintf("[%v-WARN] ", name), log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stderr, fmt.Sprintf("[%v-ERROR] ", name), log.Ldate|log.Ltime)
}
