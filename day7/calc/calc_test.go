package calc

import "testing"

func TestAdd(t *testing.T) {
	n := Add(10, 20)
	if n != 30 {
		t.Errorf("Expect 30 but got %d", n)
	}
}

func TestSub(t *testing.T) {
	n := Sub(10, 20)
	if n != -10 {
		t.Errorf("Expect -10 but got %d", n)
	}
}

func TestMul(t *testing.T) {
	n := Mul(10, 20)
	if n != 200 {
		t.Errorf("Expect 200 but got %d", n)
	}
}

func TestDiv(t *testing.T) {
	n := Div(10, 20)
	if n != 0.5 {
		t.Errorf("Expect 0.5 but got %f", n)
	}
}
