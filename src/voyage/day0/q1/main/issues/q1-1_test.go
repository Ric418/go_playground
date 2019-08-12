package issues

import (
	"testing"
)

func Test(t *testing.T) {
	actual := Fib(10)
	expected := 55
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}