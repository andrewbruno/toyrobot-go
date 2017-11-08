package main

import (
	"os"
	"toyrobot/ui"
)

func main() {

	ui := &ui.UI{
		WriterOK:  os.Stdout,
		WriterErr: os.Stderr,
		ReaderIn:  os.Stdin,
	}

	ui.Start()

}
