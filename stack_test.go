package stack

import "testing"

func TestEmpty(t *testing.T) {
	s := New()
	if s.Empty() == false {
		t.Error("Stack should be empty")
	}
	s.Push(10)
	if s.Empty() == true {
		t.Error("Stack should not be empty")
	}
	s.Pop()
	if s.Empty() == false {
		t.Error("Stack should be empty")
	}
}

func TestLen(t *testing.T) {
	s := New()

	// Test that s.Len() is incremented as items pushed
	for i := 0; i < 10; i++ {
		if s.Len() != i {
			t.Errorf("Stack length is %d instead of %d", i, s.Len())
		}
		s.Push(1)
	}
	// Test that s.Len() is decremented as items popped
	for i := 10; i >= 0; i-- {
		if s.Len() != i {
			t.Errorf("Stack length is %d instead of %d", i, s.Len())
		}
		s.Pop()
	}
	// Test that s.Len() doesn't become negative when nothing to pop
	s.Pop()
	if s.Len() != 0 {
		t.Errorf("Stack length is %d instead of 0", s.Len())
	}
}

func TestPush(t *testing.T) {
	s := New()
	err := s.Push(10)
	if err != nil {
		t.Errorf("Stack Push error: %s", err)
	}
	err = s.Push(11)
	if err != nil {
		t.Errorf("Stack Push error: %s", err)
	}
}

func TestPop(t *testing.T) {
	s := New()
	v1 := 10
	v2 := 11

	_, err := s.Pop()
	if err == nil {
		t.Errorf("stack.Pop() should have returned Empty Stack error")
	}
	s.Push(v1)
	s.Push(v2)
	val, err := s.Pop()
	if err != nil {
		t.Errorf("stack.Pop() returned error: %s", err)
	}
	if val != v2 {
		t.Error("stack.Pop() returned %d but should have returned %d", val, v2)
	}

	val, err = s.Pop()
	if err != nil {
		t.Errorf("stack.Pop() returned error: %s", err)
	}
	if val != v1 {
		t.Error("stack.Pop() returned %d but should have returned %d", val, v1)
	}
}
