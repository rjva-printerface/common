package helpers

import (
	"fmt"
)

type Log struct {
	prefix string
}

func NewLog(prefix string) *Log {
	return &Log{prefix}
}

func (l *Log) Print(message string, color func(...interface{}) string) {
	fmt.Println("[" + color(l.prefix) + "] " + message)
}
