package chapter1

import "testing"

func TestReturns(t *testing.T) {
	n1, n2 := Names()
	t.Log(n1, n2)

	n3, _ := Names()
	t.Log(n3)
}

func TestReturns2(t *testing.T) {
	n1, n2 := Names2()
	t.Log(n1, n2)
}
