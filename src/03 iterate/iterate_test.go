package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("Apple")
	expected := "AppleAppleAppleAppleApple"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	} else {
		println("expected " + expected + " and got " + repeated)
	}
}
