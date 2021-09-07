package st_logger

import (
	"bytes"
	"log"
	"os"
)

var logger *log.Logger

const (
	INFO  = iota
	ERROR = iota
	FATAL = iota
	DEBUG = iota
)

func InitializeLogger() {
	logger = log.New(os.Stdout,
		"",
		log.Ldate|log.Ltime)
}

func WriteLogMessage(level int, message ...string) {
	var buffer bytes.Buffer
	if level == INFO {
		buffer.WriteString("info : ")
	} else if level == ERROR {
		buffer.WriteString("error : ")
	} else if level == FATAL {
		buffer.WriteString("fatal : ")
	} else if level == DEBUG {
		buffer.WriteString("debug : ")
	}

	for _, m := range message {
		buffer.WriteString(m)
		buffer.WriteString(" ")
	}

	logger.Println(buffer.String())

	if level == FATAL {
		logger.Println("fatal error occure commiting suicide")
		os.Exit(1)
	}
}
