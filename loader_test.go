package loader

import (
	"bytes"
	"os"
	"testing"
	"time"
)

// test for creating a new spinner
func TestNew(t *testing.T) {
	// create a new spinner
	s := New(os.Stdout, "rs", time.Millisecond, "", "")

	// check the spinner
	if s.Name != "rs" || s.Delay != time.Millisecond {
		t.Errorf("New() failed. Expected %s, %d, got %s, %d", "rs", time.Millisecond, s.Name, s.Delay)
	}
}


// test for running the spinner
func TestInitialize(t *testing.T) {
	// create a new spinner
	s := New(os.Stdout, "rs", time.Millisecond, "", "")

	// create a new buffer
	var b bytes.Buffer

	// set the spinner output
	s.Output = &b

	// run the spinner
	s.Initialize()

	// check the spinner output
	if b.String() != "\\"|| b.String() != "|"|| b.String() != "/"|| b.String() != "-" || b.String() != "" {
		t.Errorf("Initialize() failed. Expected %s or %s or %s or %s, got %s", "\\", "|", "/", "-", b.String())
	}
}

// test the spinner stop
func  TestEnd(t *testing.T) {
	// create a new spinner
	s := New(os.Stdout, "rs", time.Millisecond, "", "Completed")

	// create a new buffer
	var b bytes.Buffer

	// set the spinner output
	s.Output = &b

	// run the spinner
	s.Initialize()

	// stop the spinner
	s.End("Completed")

	// check the spinner output
	if b.String() != "Completed" {
		t.Errorf("End() failed. Expected %s , got %s", "Completed", b.String())
	}
}