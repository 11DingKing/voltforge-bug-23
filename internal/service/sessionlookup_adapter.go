package service

import "errors"

func ClassifySessionLookup(present bool) string {
	err := LoadSessionLookup(present)
	if err == nil {
		return "ok"
	}
	if errors.Is(err, ErrSessionLookupMissing) {
		return "retest_required"
	}
	return "internal_error"
}
