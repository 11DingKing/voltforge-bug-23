package service

import "errors"

func ClassifySessionLookup(present bool) string {
	err := LoadSessionLookup(present)
	if err == nil {
		return "ok"
	}
	if err.Error() == ErrSessionLookupMissing.Error() {
		return "retest_required"
	}
	return "internal_error"
}
