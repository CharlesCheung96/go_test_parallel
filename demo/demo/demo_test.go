package demo

import "testing"

func TestInDemoA(t *testing.T) {
	t.Parallel()
	accessInner("TestInDemoA")
}

func TestInDemoB(t *testing.T) {
	t.Parallel()
	accessInner("TestInDemoA")
}
