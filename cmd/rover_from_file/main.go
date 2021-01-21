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

	flag.StringVar(&inputFile, "input", "input.txt", `The input is the path to a file containing instructions.
Can be set to @std to read from the terminal input.`)
	flag.BoolVar(&verbose, "verbose", false, "Output verbose logging.")
	flag.Parse()

	var (
		reader io.Reader
		err    error
		isStd  = inputFile == "@std"
	)

	// Set the reader.
	if isStd {
		reader = os.Stdin
	} else {
		reader, err = os.Open(inputFile)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to open file: %s", err)
			return
		}
	}

	var executor = marsrover.NewExecutor(reader)

	// Capture a ctrl+c signal to exit program
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc

		complete(executor)
		os.Exit(0)
	}()
	if isStd {
		fmt.Println("Press Ctrl+C to complete and exit.")
	}

	// Application loop.
	for err != marsrover.ErrorEndOfInstructions {
		// Format console.
		if isStd {
			fmt.Print("> ")
		}

		err = executor.Tick()
		if err != nil && err != marsrover.ErrorEndOfInstructions {
			_, _ = fmt.Fprintf(os.Stderr, "Tick: %s\n", err)
		}
	}

	complete(executor)
}

// complete will output the locations of all rovers to stdout.
func complete(executor marsrover.Executor) {
	fmt.Print("\n----------\n")
	for _, v := range executor.Rovers {
		fmt.Println(v.CurrentLocation)
	}
}
