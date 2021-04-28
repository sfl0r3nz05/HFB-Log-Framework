package main

import (
	"io"
	"os"
	"log"
	"bytes"	
)

func captureOutput(f func()) string {
	reader, writer, _:= os.Pipe()
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	f()
	writer.Close()
	return <-out
}