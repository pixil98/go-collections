package collections

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	tests := map[string]struct {
		pushItems      []int
		expectedLength int
	}{
		"push single item": {
			pushItems:      []int{1},
			expectedLength: 1,
		},
		"push multiple items": {
			pushItems:      []int{1, 2, 3},
			expectedLength: 3,
		},
		"push no items": {
			pushItems:      []int{},
			expectedLength: 0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var s Stack[int]
			for _, item := range tc.pushItems {
				s.Push(item)
			}
			if s.Length() != tc.expectedLength {
				t.Errorf("expected length %d, got %d", tc.expectedLength, s.Length())
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := map[string]struct {
		pushItems   []int
		popCount    int
		expectedPop int
		expectedLen int
	}{
		"pop from stack with one item": {
			pushItems:   []int{42},
			popCount:    1,
			expectedPop: 42,
			expectedLen: 0,
		},
		"pop from stack with multiple items returns last pushed": {
			pushItems:   []int{1, 2, 3},
			popCount:    1,
			expectedPop: 3,
			expectedLen: 2,
		},
		"pop from empty stack returns zero value": {
			pushItems:   []int{},
			popCount:    1,
			expectedPop: 0,
			expectedLen: 0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var s Stack[int]
			for _, item := range tc.pushItems {
				s.Push(item)
			}

			var popped int
			for i := 0; i < tc.popCount; i++ {
				popped = s.Pop()
			}

			if popped != tc.expectedPop {
				t.Errorf("expected popped value %d, got %d", tc.expectedPop, popped)
			}
			if s.Length() != tc.expectedLen {
				t.Errorf("expected length %d, got %d", tc.expectedLen, s.Length())
			}
		})
	}
}

func TestStack_Peek(t *testing.T) {
	tests := map[string]struct {
		pushItems    []int
		expectedPeek int
		expectedLen  int
	}{
		"peek at stack with one item": {
			pushItems:    []int{42},
			expectedPeek: 42,
			expectedLen:  1,
		},
		"peek at stack with multiple items returns last pushed": {
			pushItems:    []int{1, 2, 3},
			expectedPeek: 3,
			expectedLen:  3,
		},
		"peek at empty stack returns zero value": {
			pushItems:    []int{},
			expectedPeek: 0,
			expectedLen:  0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var s Stack[int]
			for _, item := range tc.pushItems {
				s.Push(item)
			}

			peeked := s.Peek()

			if peeked != tc.expectedPeek {
				t.Errorf("expected peeked value %d, got %d", tc.expectedPeek, peeked)
			}
			if s.Length() != tc.expectedLen {
				t.Errorf("expected length %d after peek, got %d", tc.expectedLen, s.Length())
			}
		})
	}
}

func TestStack_Length(t *testing.T) {
	tests := map[string]struct {
		pushItems      []int
		popCount       int
		expectedLength int
	}{
		"empty stack has length zero": {
			pushItems:      []int{},
			popCount:       0,
			expectedLength: 0,
		},
		"stack with items has correct length": {
			pushItems:      []int{1, 2, 3, 4, 5},
			popCount:       0,
			expectedLength: 5,
		},
		"length decreases after pop": {
			pushItems:      []int{1, 2, 3},
			popCount:       2,
			expectedLength: 1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var s Stack[int]
			for _, item := range tc.pushItems {
				s.Push(item)
			}
			for i := 0; i < tc.popCount; i++ {
				s.Pop()
			}

			if s.Length() != tc.expectedLength {
				t.Errorf("expected length %d, got %d", tc.expectedLength, s.Length())
			}
		})
	}
}

func TestStack_LIFO(t *testing.T) {
	tests := map[string]struct {
		pushItems   []int
		expectedPop []int
	}{
		"items are popped in LIFO order": {
			pushItems:   []int{1, 2, 3, 4, 5},
			expectedPop: []int{5, 4, 3, 2, 1},
		},
		"single item": {
			pushItems:   []int{42},
			expectedPop: []int{42},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var s Stack[int]
			for _, item := range tc.pushItems {
				s.Push(item)
			}

			for i, expected := range tc.expectedPop {
				popped := s.Pop()
				if popped != expected {
					t.Errorf("pop %d: expected %d, got %d", i, expected, popped)
				}
			}
		})
	}
}
