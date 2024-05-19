package object

import "testing"

func TestIntegerHashKey(t *testing.T) {
	i1 := &Integer{Value: 5}
	i2 := &Integer{Value: 5}
	i3 := &Integer{Value: 10}
	i4 := &Integer{Value: 0}

	if i1.HashKey() != i2.HashKey() {
		t.Errorf("objects with same value have different hash keys")
	}

	if i1.HashKey() == i3.HashKey() {
		t.Errorf("objects with different values have same hash keys")
	}

	if i4.HashKey() == i3.HashKey() {
		t.Errorf("objects with different values have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	b1 := &Boolean{Value: true}
	b2 := &Boolean{Value: true}
	b3 := &Boolean{Value: false}

	if b1.HashKey() != b2.HashKey() {
		t.Errorf("objects with same value have different hash keys")
	}

	if b1.HashKey() == b3.HashKey() {
		t.Errorf("objects with different values have same hash keys")
	}
}

func TestStringHashKey(t *testing.T) {
	s1 := &String{Value: "Hello World"}
	s2 := &String{Value: "Hello World"}
	s3 := &String{Value: "My name is johnny"}
	s4 := &String{Value: ""}

	if s1.HashKey() != s2.HashKey() {
		t.Errorf("objects with same value have different hash keys")
	}

	if s1.HashKey() == s3.HashKey() {
		t.Errorf("objects with different values have same hash keys")
	}

	if s4.HashKey() == s3.HashKey() {
		t.Errorf("objects with different values have same hash keys")
	}
}
