package logger

import (
	"log"
	"os"
)

var errLog = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags)
var inforLog = log.New(os.Stdout, "\033[1;32mINFOR: \u001b[0m", log.LstdFlags)
var warningLog = log.New(os.Stdout, "\u001b[33mWARNING: \u001b[0m", log.LstdFlags)

func LogError(err interface{}) {
	errLog.Println("[go-redis]:", err)
}
func LogInfor(msg interface{}) {
	inforLog.Println("[go-redis]:", msg)
}
func LogWarning(msg interface{}) {
	warningLog.Println("[go-redis]:", msg)
}
