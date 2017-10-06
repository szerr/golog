package log

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

var (
	Level = 0
)

const (
	InfoLevel = iota
	WarnLevel
	ErrorLevel
	FatalLevel
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
	if Level <= InfoLevel {
		prin("Info", s...)
	}
}

func Warn(s ...interface{}) {
	if Level <= WarnLevel {
		prin("Warn", s...)
	}
}

func Err(s ...interface{}) {
	if Level <= ErrorLevel {
		prin("Error", s...)
	}
}

func Fat(s ...interface{}) {
	if Level <= FatalLevel {
		prin("Fatal", s...)
	}
}
