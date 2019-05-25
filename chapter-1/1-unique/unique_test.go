package unique

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	if !IsUnique("abcdef") {
		t.Fatalf("String %q should be unique but was %v", "abcdef", false)
	}

	if IsUnique("abcdefaa") {
		t.Fatalf("String %q shouldn't be unique but was %v", "abcdefaa", false)
	}
}
