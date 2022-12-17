package golog

import (
	"errors"
	"io"
	"log"
	"os"
)

const (
	LstdFlags = log.LstdFlags | log.Lshortfile
)

var stdLogger *logger

type logger struct {
	logger *log.Logger
}

func init() {
	stdLogger = &logger{
		logger: log.New(os.Stdout, "", LstdFlags),
	}
}

func SetOutput(w io.Writer) {
	stdLogger.logger.SetOutput(w)
}

func SetFlags(flag int) {
	stdLogger.logger.SetFlags(flag)
}

func SetPrefix(prefix string) {
	stdLogger.logger.SetPrefix(prefix)
}

func Info(v ...any) {
	stdLogger.logger.Println(append([]any{"Info"}, v...)...)
}

func Warn(v ...any) {
	stdLogger.logger.Println(append([]any{"Warn"}, v...)...)
}

func Error(v ...any) {
	stdLogger.logger.Println(append([]any{"Error"}, v...)...)
}

func Fatal(v ...any) {
	stdLogger.logger.Println(append([]any{"Fatal"}, v...)...)
}

func New(w io.Writer, prefix string, flag int) *logger {
	return &logger{
		logger: log.New(w, prefix, flag),
	}
}

func (l *logger) Info(v ...any) {
	l.logger.Println(append([]any{"Info"}, v...)...)
}

func (l *logger) Warn(v ...any) {
	l.logger.Println(append([]any{"Warn"}, v...)...)
}

func (l *logger) Error(v ...any) {
	l.logger.Println(append([]any{"Error"}, v...)...)
}

func (l *logger) Fatal(v ...any) {
	l.logger.Println(append([]any{"Fatal"}, v...)...)
}

func main() {
	Info(errors.New("测试"), "信息")
}
