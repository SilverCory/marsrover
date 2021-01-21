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
		inputFile  string
		outputFile string
	)

	flag.StringVar(&inputFile, "input", "input.txt", `The input is the path to a file containing instructions.
Can be set to @std to read from the terminal input.`)
	flag.StringVar(&outputFile, "output", "@std", `The output is the path to write a file containing the new locations of rovers.
Can be set to @std to write to the terminal output.`)
	flag.Parse()

	var (
		reader   io.Reader
		writer   io.WriteCloser
		err      error
		isStdIn  = inputFile == "@std"
		isStdOut = outputFile == "@std"
	)

	if reader, err = openReader(isStdIn, inputFile); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to open input file: %s", err)
		os.Exit(1)
		return
	}

	if writer, err = openWriter(isStdOut, outputFile); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to open output file: %s", err)
		os.Exit(1)
		return
	}

	// Ensure writer gets closed.
	defer func() {
		if isStdOut {
			return // Don't close stdout.
		}

		if err := writer.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to close output file: %s", err)
			os.Exit(1)
		}
	}()

	// Create executor instance.
	var executor = marsrover.NewExecutor(reader)

	// Capture a ctrl+c signal to exit program
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc

		complete(writer, executor)
		os.Exit(0)
	}()

	// Output a message letting user know to Ctrl+C.
	if isStdIn {
		fmt.Println("Press Ctrl+C to complete and exit.")
	}

	// Application loop.
	for err != marsrover.ErrorEndOfInstructions {
		// Format console.
		if isStdIn {
			fmt.Print("> ")
		}

		err = executor.Tick()
		if err != nil && err != marsrover.ErrorEndOfInstructions {
			_, _ = fmt.Fprintf(os.Stderr, "Tick: %s\n", err)
		}
	}

	// Output results.
	complete(writer, executor)
}

// openReader will open the appropriate reader.
func openReader(isStdIn bool, inputFile string) (reader io.Reader, err error) {
	// Set the reader.
	if isStdIn {
		reader = os.Stdin
	} else {
		reader, err = os.Open(inputFile)
	}

	return reader, err
}

// openWriter will open the appropriate writer.
func openWriter(isStdOut bool, outputFile string) (writer io.WriteCloser, err error) {
	// Set the writer
	if isStdOut {
		writer = os.Stdout
	} else {
		writer, err = os.Create(outputFile)
	}

	return writer, err
}

// complete will output the locations of all rovers to stdout.
func complete(w io.Writer, executor marsrover.Executor) {
	for _, v := range executor.Rovers {
		_, _ = fmt.Fprintln(w, v.CurrentLocation)
	}
}
