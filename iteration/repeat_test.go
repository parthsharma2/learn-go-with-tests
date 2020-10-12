package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertEqual := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeats 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertEqual(t, repeated, expected)
	})

	t.Run("repeats 7 times", func(t *testing.T) {
		repeated := Repeat("a", 7)
		expected := "aaaaaaa"
		assertEqual(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
