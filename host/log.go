package main

import (
	"fmt"
	"os"
)

type Logger struct {
	file *os.File
}

var (
	log *Logger
)

func init() {
	file, err := os.OpenFile("foxpi.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	log = &Logger{file}
}

func (l *Logger) Log(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)

	log.file.Write([]byte(message + "\n"))
	log.file.Sync()
}
