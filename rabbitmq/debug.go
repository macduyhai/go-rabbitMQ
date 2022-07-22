package rabbitmq

import "log"

var Debug bool

func debugf(format string, args ...interface{}) {
	if !Debug {
		return
	}

	log.Printf(format, args...)
}
