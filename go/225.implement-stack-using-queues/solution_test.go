package implementstackusingqueues

import (
	"testing"
)

func TestMyStack(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)

	if top := s.Top(); top != 2 {
		t.Errorf("s.Top() = %v, want: %v", top, 2)
	}
	if pop := s.Pop(); pop != 2 {
		t.Errorf("s.Pop() = %v, want: %v", pop, 2)
	}
	if s.Empty() {
		t.Errorf("s.Empty() == true, want false")
	}
}
