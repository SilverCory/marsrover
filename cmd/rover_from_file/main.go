package main

import (
	"flag"
	"fmt"
	"io"
	"marsrover"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		inputFile string
		verbose   bool
	)

	flag.StringVar(&inputFile, "input", "@std", `The input is the path to a file containing instructions.
Can be set to @std to read from the terminal input.`)
	flag.BoolVar(&verbose, "verbose", false, "Output verbose logging.")
	flag.Parse()

	var reader io.Reader
	var err error
	if inputFile == "@std" {
		reader = os.Stdin
	} else {
		reader, err = os.Open(inputFile)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to open file: %s", err)
			return
		}
	}

	var executor = marsrover.NewExecutor(reader)

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
		fmt.Print("\n\n\n")

		for _, v := range executor.Rovers {
			fmt.Println(v.CurrentLocation)
		}
		os.Exit(0)
	}()

	for err != marsrover.ErrorEndOfInstructions {

		if inputFile == "@std" {
			fmt.Print("> ")
		}

		err = executor.Tick()
		if err != nil && err != marsrover.ErrorEndOfInstructions {
			_, _ = fmt.Fprintf(os.Stderr, "Tick: %s\n", err)
		}
	}
}
