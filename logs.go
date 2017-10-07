package log

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"os"
	"io"
)

var (
	Level = 0
	logger *log.Logger
)

const (
	InfoLevel = iota
	WarnLevel
	ErrorLevel
	FatalLevel
)

func init(){
	SetWriter(os.Stdout)
}

func SetWriter(out io.Writer){
	logger = log.New(out, "", log.Ldate | log.Ltime)
}

func Println(flag string, s ...interface{}) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	logger.Print(flag, " ", file+":"+strconv.Itoa(line), " ", fmt.Sprintln(s...))
}

func Info(s ...interface{}) {
	if Level <= InfoLevel {
		Println("Info", s...)
	}
}

func Warn(s ...interface{}) {
	if Level <= WarnLevel {
		Println("Warn", s...)
	}
}

func Err(s ...interface{}) {
	if Level <= ErrorLevel {
		Println("Error", s...)
	}
}

func Fat(s ...interface{}) {
	if Level <= FatalLevel {
		Println("Fatal", s...)
	}
}
