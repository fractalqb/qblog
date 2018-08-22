package qblog

import (
	"os"
)

var stdLogger = New(os.Stderr, "qblog")

func Print(v ...interface{}) { stdLogger.Print(v...) }

func Printf(format string, v ...interface{}) { stdLogger.Printf(format, v...) }

func Println(v ...interface{}) { stdLogger.Println(v...) }

func Panic(v ...interface{}) { stdLogger.Print(v...) }

func Panicf(format string, v ...interface{}) { stdLogger.Printf(format, v...) }

func Panicln(v ...interface{}) { stdLogger.Println(v...) }

func Fatal(v ...interface{}) { stdLogger.Print(v...) }

func Fatalf(format string, v ...interface{}) { stdLogger.Printf(format, v...) }

func Fatalln(v ...interface{}) { stdLogger.Println(v...) }
