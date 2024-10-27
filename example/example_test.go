package example

import "testing"

func TestReturnsTrue(t *testing.T) {
	if ReturnsTrue() != true {
		t.Error("Expected true, got false")
	}
}

func TestReturnsFalse(t *testing.T) {
	if ReturnsFalse() != false {
		t.Error("Expected false, got true")
	}
}
