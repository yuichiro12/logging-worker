package logworker

import (
	"io"
	"log"
	"strings"
)

type Logger struct {
	w         io.Writer
	errChan   chan error
	Separator string
}

func NewLogger(w io.Writer, errChan chan error, sep string) Logger {
	return Logger{
		w:         w,
		errChan:   errChan,
		Separator: sep,
	}
}

func (l Logger) LogRow(rc <-chan []string) {
	for r := range rc {
		if _, err := l.w.Write([]byte(strings.Join(r, l.Separator) + "\n")); err != nil {
			l.errChan <- err
		}
	}
}

func LogError(w io.Writer, ec <-chan error) {
	for e := range ec {
		if _, err := w.Write([]byte(e.Error() + "\n")); err != nil {
			log.Fatalln(err)
		}
	}
}
