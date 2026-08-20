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
	return fmt.Errorf("load sessionlookup: %v", ErrSessionLookupMissing)
}
