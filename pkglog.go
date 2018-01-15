package qblog

import (
	"io"
)

type PkgConfig struct {
	pkgRoot *Logger
}

func Package(root *Logger) PkgConfig {
	return PkgConfig{root}
}

func (pkg PkgConfig) SetParent(l *Logger) {
	l.AddSub(pkg.pkgRoot)
}

func (cfg PkgConfig) Level() int { return cfg.pkgRoot.lvl }

func (cfg PkgConfig) SetLevel(level int) { cfg.pkgRoot.SetLevel(level) }

func (cfg PkgConfig) SetOutput(w io.Writer) { cfg.pkgRoot.SetOutput(w) }
