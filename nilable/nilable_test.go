package nilable_test

import (
	"sus/nilable"
	"testing"
)

func TestJust(t *testing.T) {
	n := nilable.Just("Hello")
	if !n.HasValue() {
		t.Error("Expected the nilable to have HasValue be set to true")
	}
	v, ok := n.Value()
	if !ok {
		t.Error("Expected the nilable to have its value actually exist")
	}
	if v != "Hello" {
		t.Errorf("Expected %s but got %s", "Hello", v)
	}
	v2 := n.ValueOrDefault("Cool")
	if v2 != "Hello" {
		t.Errorf("Expected the nilable to contain %s, but it likely contains nothing for some odd reason", v2)
	}
}

func TestNil(t *testing.T) {
	n := nilable.Nil[string]()
	if n.HasValue() {
		t.Error("Expected the nilable to not have any value")
	}
	_, ok := n.Value()
	if ok {
		t.Error("Expected the nilable to be empty")
	}
	v := n.ValueOrDefault("Cool")
	if v != "Cool" {
		t.Errorf("Expected the nilable to default to %s, but got %s", "Cool", v)
	}
}
