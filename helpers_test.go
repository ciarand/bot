package bot

import "testing"

func assertSame(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("Expected %s to equal %s", a, b)
	}
}
