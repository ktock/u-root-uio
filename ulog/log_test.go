package ulog

import (
	"log"
	"strings"
	"testing"
)

func TestDefault(t *testing.T) {
	var b strings.Builder
	log.SetPrefix("[foobar] ")
	log.SetOutput(&b)
	log.SetFlags(0)

	Log.Printf("Some output")

	want := "[foobar] Some output\n"
	if got := b.String(); got != want {
		t.Errorf("log is %q, want %q", got, want)
	}
}
