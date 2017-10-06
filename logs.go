package log

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

func prin(flag string, s ...interface{}) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	log.Println(flag, file+":"+strconv.Itoa(line), fmt.Sprint(s...))
}

func Info(s ...interface{}) {
	prin("Info", s...)
}

func Warn(s ...interface{}) {
	prin("Warn", s...)
}

func Err(s ...interface{}) {
	prin("Error", s...)
}

func Fat(s ...interface{}) {
	prin("Fatal", s...)
}
