// Packge qblog implements a thin cnovenience wrapper around go standard
// logging. It goes for few primary ideas. First it shall be rather simple to
// migrate from standard logging simply by changing the import statement.
// Second is the notion of Log Level per logger. And third that loggers can –
// but need not – be hierarchically organized to ease configuration.
package qblog

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
	out  io.Writer
	lvl  int
	subs map[*Logger]bool
}

const (
	Error   = 65535
	Warn    = 32767
	Info    = 0
	Debug   = -32767
	Trace   = -65535
	Default = Info
)

func NewLevel(out io.Writer, prefix string, level int) *Logger {
	res := &Logger{
		Logger: log.New(out, prefix, log.LstdFlags),
		out:    out,
		lvl:    level}
	return res
}

func New(out io.Writer, prefix string) *Logger {
	return NewLevel(out, prefix, Default)
}

func Std(prefix string) *Logger {
	return NewLevel(os.Stderr, prefix, Default)
}

func (l *Logger) Log(level int, v ...interface{}) {
	if level >= l.lvl {
		l.Print(v...)
	}
}

func (l *Logger) Logf(level int, format string, v ...interface{}) {
	if level >= l.lvl {
		l.Printf(format, v...)
	}
}

func (l *Logger) Logln(level int, format string, v ...interface{}) {
	if level >= l.lvl {
		l.Println(v...)
	}
}

func (l *Logger) Level() int { return l.lvl }

func (l *Logger) SetLevel(level int) {
	l.lvl = level
	if l.subs != nil {
		for sub, _ := range l.subs {
			sub.SetLevel(level)
		}
	}
}

func (l *Logger) SetOutput(w io.Writer) {
	l.out = w
	l.Logger.SetOutput(w)
}

func (parent *Logger) AddSub(sub *Logger) {
	if parent.subs == nil {
		parent.subs = make(map[*Logger]bool)
	}
	parent.subs[sub] = true
	sub.SetLevel(parent.Level())
	sub.SetOutput(parent.out)
}

func (parent *Logger) NewSub(prefix string) *Logger {
	res := NewLevel(parent.out, prefix, parent.Level())
	if parent.subs == nil {
		parent.subs = make(map[*Logger]bool)
	}
	parent.subs[res] = true
	return res
}
