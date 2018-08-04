package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/janbaer/mp3db/cmd"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tj/go/term"
)

func main() {
	handleInterruptSignal()

	defer cleanup()

	term.HideCursor()

	cmd.Execute()
}

func cleanup() {
	term.ShowCursor()
	os.Exit(0)
}

func handleInterruptSignal() {
	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		fmt.Println()
		fmt.Println("Received an interrupt, stop processing... ", sig)
		cleanup()
	}()
}
