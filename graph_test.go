package imedge

import "testing"

func TestGraph_Create(t *testing.T) {
	g := NewGraph()
	var want, got uint
	want, got = 0, g.Create()
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
	want, got = 1, g.Create()
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestGraph_Connect(t *testing.T) {
	g := NewGraph()
	tail := g.Create()
	head := g.Create()
	g.Connect(tail, head)

	heads := g.ReadPositive(tail)

	if got := heads[0]; got != 1 {
		t.Errorf("want %v, got %v", 1, got)
	}

	tails := g.ReadNegative(head)

	if got := tails[0]; got != 0 {
		t.Errorf("want %v, got %v", 0, got)
	}
}