package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

//Reader Interface reads bytes
type Reader interface {
	Read(p []byte) (n int, err error)
}

//Writer Interface writes bytes
type Writer interface {
	Write(p []byte) (n int, err error)
}

//GoNumReader stuct implement Reader interface
type GoNumReader struct{}

func (fooReader *GoNumReader) Read(b []byte) (int, error) {
	// Read some data from somewhere, anywhere.
	fmt.Print("gonum> ")
	return os.Stdout.Read(b)
}

//GoNumWriter implements Writer interface
type GoNumWriter struct{}

func (fooWriter *GoNumWriter) Write(b []byte) (int, error) {
	// Write data somewhere
	fmt.Print("> ")
	return os.Stdout.Write(b)
}

func main() {
	//Instantiate reader and writer
	var (
		reader GoNumReader
		writer GoNumWriter
	)

	fmt.Printf("\u2591\u2588\u2580\u2580\u2591\u2588\u2580\u2588\u2591\u2588\u2580\u2588\u2591\u2588\u2591\u2588\u2591\u2588\u2584\u2588\u2591\u2588\u2580\u2580\u2591\u2588\u2580\u2584\u2591\u2588\u2580\u2588\u2591\u2580\u2588\u2580\u2591\u2588\u2580\u2588\u2591\u2588\u2580\u2584 \n")
	fmt.Printf("\u2591\u2588\u2591\u2588\u2591\u2588\u2591\u2588\u2591\u2588\u2591\u2588\u2591\u2588\u2591\u2588\u2591\u2588\u2591\u2588\u2591\u2588\u2580\u2580\u2591\u2588\u2580\u2584\u2591\u2588\u2580\u2588\u2591\u2591\u2588\u2591\u2591\u2588\u2591\u2588\u2591\u2588\u2580\u2584 \n")
	fmt.Printf("\u2591\u2580\u2580\u2580\u2591\u2580\u2580\u2580\u2591\u2580\u2591\u2580\u2591\u2580\u2580\u2580\u2591\u2580\u2591\u2580\u2591\u2580\u2580\u2580\u2591\u2580\u2591\u2580\u2591\u2580\u2591\u2580\u2591\u2591\u2580\u2591\u2591\u2580\u2580\u2580\u2591\u2580\u2591\u2580 \n")
	fmt.Printf(" -----------------------------------\n")
	fmt.Printf("/  Looking for clues?...            \\ \n")
	fmt.Printf("\\  Enumerate! Enumerate! Enumerate! /\n")
	fmt.Printf(" -----------------------------------  \n")
	fmt.Printf("        \\   ^__^\n")
	fmt.Printf("         \\  (oo)\\_______\n")
	fmt.Printf("            (__)\\       )\\/\\ \n")
	fmt.Printf("                 ||----w |       \n")
	fmt.Printf("                 ||     ||       \n")

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}

}
