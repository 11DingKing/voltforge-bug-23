package service

import (
	"errors"
	"testing"
)

func TestVoltForge23(t *testing.T) {
	if got := ClassifySessionLookup(false); got != "retest_required" {
		t.Fatalf("got %s", got)
	}
	if !errors.Is(LoadSessionLookup(false), ErrSessionLookupMissing) {
		t.Fatal("missing identity was not preserved")
	}
}
