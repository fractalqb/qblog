package qblog

import (
	"io"
)

// PackageLog is a wrapper for a packages root loger that allows users of a
// package to apply some configuration to the package's logger. Generally the
// root logger of a packe shall not be exported to avoid users to fiddle around
// with details that are private. Intended use:
//
//  package mypackage
//
//  var log = qblog.New(os.Stderr, "mypackage")
//  var LogConfig = qblog.Package(log)
//  …
//
// Now an application can use mypackage.LogConfig to configure package logging.
type PackageLog struct {
	pkgRoot *Logger
}

func Package(root *Logger) PackageLog {
	return PackageLog{root}
}

func (pkg PackageLog) SetParent(l *Logger) {
	l.AddSub(pkg.pkgRoot)
}

func (cfg PackageLog) Level() Level { return cfg.pkgRoot.lvl }

func (cfg PackageLog) SetLevel(level Level) { cfg.pkgRoot.SetLevel(level) }

func (cfg PackageLog) SetOutput(w io.Writer) { cfg.pkgRoot.SetOutput(w) }
