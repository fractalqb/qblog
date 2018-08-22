// Package stdcompat is not intendet to be used or run. Its only use it to test
// if code written for Go's standard log package would also compile with qblog
// by simply replacing
//
//  import "log"
//
// with the import of qblog
//
//  import log "git.fractalqb.de/fractalqb/qblog"
//
// Note that API compatibility is only intended for the functions generating
// log messages (see api_xxx.go) as these are the heavily used logging
// functions. Configuration and other things won't be compatible.
package main

import (
	"log"
)

//go:generate cpp -P -DGO_STDLOG apicompat.i -o api_stdlog.go
//go:generate cpp -P apicompat.i -o api_qblog.go

func main() {
	log.Fatal("[qb]log source level API compatibility with go \"log\": OK")
}
