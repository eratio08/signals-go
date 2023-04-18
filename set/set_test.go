package set

import "testing"

func TestAdd(t *testing.T) {
	s := New[int]()

	s.Add(23)

	if len(s.elements) != 1 {
		t.Fatalf("Unexpected size, expected %d but was %d", 1, len(s.elements))
	}

	if _, ok := s.elements[23]; !ok {
		t.Fatalf("Set did not include element %d, elements are %#v", 23, s.elements)
	}
}

func TestRemove(t *testing.T) {
  s := New(1,2,3)

  s.Remove(2)

  if len(s.elements) != 2 {
    t.Fatalf("Unexpected size, expected %d but was %d (%v)", 2, len(s.elements), s)
  }

  if _, ok := s.elements[2]; ok {
    t.Fatalf("%d is still in the set %v", 2, s)
  }
}

func TestLen(t *testing.T) {
  s := New("1", "2", "3")

  l := s.Len()

  if l != 3 {
  t.Fatalf("Expected length of %d, got %d", len(s.elements), l)
  }
}

func TestToSlice(t *testing.T) {
  s := New(1.2,1.3,1.4)

  sl := s.ToSlice()

  if len(sl) != len(s.elements) {
    t.Fatalf("Expected slice to have a lenght of %d, got %d (%v)", len(s.elements), len(sl), sl)
  }
}
