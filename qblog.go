// Packge qblog implements a thin cnovenience wrapper around go standard
// logging. It goes for few primary ideas. First it shall be rather simple to
// migrate from standard logging simply by changing the import statement.
// Second is the notion of Log Level per logger. And third that loggers can –
// but need not – be hierarchically organized to ease configuration.
package qblog

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Level uint8

type Logger struct {
	*log.Logger
	pfx  string
	out  io.Writer
	lvl  Level
	subs map[*Logger]bool
}

const (
	lfatal   Level = 255
	lpanic   Level = 254
	Lerror   Level = 252
	Lwarn    Level = 189
	Linfo    Level = 126
	Ldebug   Level = 63
	Ltrace   Level = 0
	Ldefault Level = Linfo
)

func NewLevel(out io.Writer, prefix string, level Level) *Logger {
	res := &Logger{
		Logger: log.New(out, prefix, log.LstdFlags),
		pfx:    prefix,
		out:    out,
		lvl:    level}
	return res
}

func New(out io.Writer, prefix string) *Logger {
	return NewLevel(out, prefix, Ldefault)
}

func Std(prefix string) *Logger {
	return NewLevel(os.Stderr, prefix, Ldefault)
}

func (l *Logger) Trace(v ...interface{}) { l.Log(Ltrace, v...) }

func (l *Logger) Tracef(fmt string, v ...interface{}) {
	l.Logf(Ltrace, fmt, v...)
}

func (l *Logger) Debug(v ...interface{}) { l.Log(Ldebug, v...) }

func (l *Logger) Debugf(fmt string, v ...interface{}) {
	l.Logf(Ldebug, fmt, v...)
}

func (l *Logger) Info(v ...interface{}) { l.Log(Linfo, v...) }

func (l *Logger) Infof(fmt string, v ...interface{}) {
	l.Logf(Linfo, fmt, v...)
}

func (l *Logger) Warn(v ...interface{}) { l.Log(Lwarn, v...) }

func (l *Logger) Warnf(fmt string, v ...interface{}) {
	l.Logf(Lwarn, fmt, v...)
}

func (l *Logger) Error(v ...interface{}) { l.Log(Lerror, v...) }

func (l *Logger) Errorf(fmt string, v ...interface{}) {
	l.Logf(Lerror, fmt, v...)
}

func (l *Logger) Logs(level Level) bool {
	return level >= l.lvl
}

func (l *Logger) Log(level Level, v ...interface{}) {
	if l.Logs(level) {
		l.adjpfx(level)
		l.Print(v...)
	}
}

func (l *Logger) Logf(level Level, format string, v ...interface{}) {
	if l.Logs(level) {
		l.adjpfx(level)
		l.Printf(format, v...)
	}
}

func (l *Logger) Level() Level { return l.lvl }

func (l *Logger) SetLevel(level Level) {
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
	for sub, _ := range l.subs {
		sub.SetOutput(w)
	}
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

var lvlNames [256]string

func init() {
	lvlNames[Ltrace] = "trace"
	lvlNames[Ldebug] = "debug"
	lvlNames[Linfo] = "info"
	lvlNames[Lwarn] = "warn"
	lvlNames[Lerror] = "error"
	lvlNames[lpanic] = "panic"
	lvlNames[lfatal] = "fatal"
}

// Short hack to get log levels into messages – I don't really like this
func (l *Logger) adjpfx(lvl Level) {
	var pfx string
	if nm := lvlNames[lvl]; len(nm) > 5 {
		pfx = fmt.Sprintf("%s [%s] ", l.pfx, nm[:5])
	} else if len(nm) > 0 {
		pfx = fmt.Sprintf("%s [%-5s] ", l.pfx, nm)
	} else {
		pfx = fmt.Sprintf("%s [%05d] ", l.pfx, lvl)
	}
	l.SetPrefix(pfx)
}
