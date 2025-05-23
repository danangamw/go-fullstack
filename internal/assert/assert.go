package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got %v; want %v", actual, expected)
	}
}

func StripContains(t *testing.T, actual, expectedSubString string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubString) {
		t.Errorf("got: %q; expected to contain: %q", actual, expectedSubString)
	}
}

func NilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("got: %v; expected: nil", actual)
	}
}
