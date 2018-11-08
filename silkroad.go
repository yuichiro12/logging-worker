package silkroad

import (
	"io"
	"log"
	"strings"
)

type Logger struct {
	ColumnSeparator string
	RowSeparator    string
}

func NewLogger(csep, rsep string) Logger {
	return Logger{
		ColumnSeparator: csep,
		RowSeparator:    rsep,
	}
}

func (l Logger) LogRow(w io.Writer, rc <-chan []string, ec chan<- error) {
	for r := range rc {
		if _, err := w.Write([]byte(strings.Join(r, l.ColumnSeparator) + l.RowSeparator)); err != nil {
			ec <- err
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
