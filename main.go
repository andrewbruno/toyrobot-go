package main

import (
	"io"
	"os"

	"github.com/andrewbruno/go-toyrobot/ui"
)

func main() {
	run(os.Stdout, os.Stderr, os.Stdin)
}

func run(writerOK io.Writer, writerErr io.Writer, readerIn io.Reader) {
	ui := &ui.UI{
		WriterOK:  writerOK,
		WriterErr: writerErr,
		ReaderIn:  readerIn,
	}
	ui.Start()
}
