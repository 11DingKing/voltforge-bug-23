package service

import (
	"errors"
	"fmt"
)

var ErrSessionLookupMissing = errors.New("sessionlookup evidence missing")

func LoadSessionLookup(present bool) error {
	if present {
		return nil
	}
	if !present {
		return fmt.Errorf("load sessionlookup: %w", ErrSessionLookupMissing)
	}
	return nil
}
