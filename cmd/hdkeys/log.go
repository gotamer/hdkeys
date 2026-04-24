// log,go
package main

import (
	"io"
	"log"
	"os"
)

var (
	Debug = *log.Default()
	Info  = *log.Default()
	Warn  = *log.Default()
	Error = *log.Default()
)

func init() {

	Debug.SetPrefix("DEBUG ")
	Debug.SetFlags(log.Lshortfile)
	Debug.SetOutput(io.Discard)

	Info.SetPrefix("INFO ")
	Info.SetFlags(log.Lshortfile)
	Info.SetOutput(io.Discard)

	Warn.SetPrefix("WARN ")
	Warn.SetFlags(log.Lshortfile)

	Error.SetPrefix("ERROR ")
	Error.SetFlags(log.Lshortfile)

}

func debug(b bool) {
	if b {
		Info.SetOutput(os.Stderr)
	} else {
		Info.SetOutput(io.Discard)
	}
	Info.Println("Debug level on")
}

func verbose(b bool) {
	if b {
		Debug.SetOutput(os.Stderr)
	} else {
		Debug.SetOutput(io.Discard)
	}
	Debug.Println("Verbose debug level on")
}
