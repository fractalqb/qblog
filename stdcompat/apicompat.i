package main

import (
	"os"
#ifdef GO_STDLOG
	"log"
#else
	log "git.fractalqb.de/fractalqb/qblog"
#endif
)

#ifdef GO_STDLOG
func stdlogCompat() {
#else
func qblogCompat() {
#endif
	log.Print("foo")
	log.Printf("bar %d", 7)
	log.Println("baz")
	log.Panic("foo")
	log.Panicf("bar %d", 8)
	log.Panicln("baz")
	log.Fatal("foo")
	log.Fatalf("bar %d", 8)
	log.Fatalln("baz")

#ifdef GO_STDLOG
 var logger = log.New(os.Stdout, "stdlog", log.LstdFlags)
#else
 var logger = log.New(os.Stdout, "qblog")
#endif
	logger.Print("foo")
	logger.Printf("bar %d", 7)
	logger.Println("baz")
	logger.Panic("foo")
	logger.Panicf("bar %d", 8)
	logger.Panicln("baz")
	logger.Fatal("foo")
	logger.Fatalf("bar %d", 8)
	logger.Fatalln("baz")
}
