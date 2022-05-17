package main

import (
	"errors"
	"fmt"
	"time"
)

const (
	BLACK = (iota + 30)
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

type Logger struct {
	prefix string
	level  string
	color  int
}

func NewLog(level string) (*Logger, error) {
	res := &Logger{
		level: level,
	}
	switch level {
	case "info":
		res.prefix = "[INFO] "
		res.color = GREEN
	case "warning":
		res.prefix = "[WARNING] "
		res.color = YELLOW
	case "error":
		res.prefix = "[ERROR] "
		res.color = RED
	default:
		return nil, errors.New("unknown logger level")
	}
	return res, nil
}

func (l *Logger) SetColor(color int) error {
	if color < 30 || color > 37 {
		return errors.New("unknown color number")
	}
	l.color = color
	return nil
}

func (l Logger) Print(text string) {
	fmt.Printf("\033[1;%dm%s\033[0m\n", l.color, l.prefix+time.Now().Format("02-01-2006 15:04:05")+" "+text)
}
