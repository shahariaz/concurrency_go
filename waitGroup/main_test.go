package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrintHello(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w
	var wg sync.WaitGroup
	wg.Add(1)
	go printHello(&wg)
	wg.Wait()
	w.Close()
	out, _ := io.ReadAll(r)
	output := string(out)
	os.Stdout = stdOut
	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected output to contain 'Hello World', but got: %s", output)
	}
}
