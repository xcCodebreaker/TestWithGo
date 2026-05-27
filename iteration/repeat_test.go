package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated, repeatedTimes := Repeat("a", 5)
	expected, expectedTimes := "aaaaa", 5

	if repeated != expected && repeatedTimes != expectedTimes {
		t.Errorf("expected %q when num of repeat is %d but got %q with %d num of repeat", expected, expectedTimes, repeated, repeatedTimes)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated, timesRepeated := Repeat("a", 5)
	fmt.Printf("%q, %d", repeated, timesRepeated)
	//Output: "aaaaa", 5
}
