package qblog

type PkgRoot struct {
	root *Logger
}

func Root(root *Logger) PkgRoot {
	return PkgRoot{root}
}

func (pkg PkgRoot) SetParent(l *Logger) {
	l.AddSub(pkg.root)
}
