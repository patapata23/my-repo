package main
import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "even" {
		t.Error("expected: evn, actual: %s", result)
	}
}