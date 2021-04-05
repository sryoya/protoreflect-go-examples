package fuzzing

import "testing"

func TestNew(t *testing.T) {
	res := New()

	t.Errorf("%v", res)
}
