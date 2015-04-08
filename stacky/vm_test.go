package stacky

import "testing"

func TestInterpret(t *testing.T) {
	cases := []struct {
		arg      instructions
		expected stack
	}{
		{instructions{instLiteral, 0x04}, stack{0x04}},
		{instructions{instLiteral, 0x04, instLiteral, 0x06}, stack{0x04, 0x06}},
		{instructions{instLiteral, 0x01, instLiteral, 0x01, instAdd}, stack{0x02}},
		{instructions{instLiteral, 0x04, instLiteral, 0x01, instMult}, stack{0x04}},
	}

	for _, c := range cases {
		vm := NewVM()
		vm.Interpret(c.arg)

		for i := range vm.stack {
			if vm.stack[i] != c.expected[i] {
				t.Errorf("expected stack to be %v, got %v", vm.stack, c.expected)
			}
		}
	}
}

func TestPush(t *testing.T) {
	vm := NewVM()

	vm.stack.push(stackVal(1))
	if l := len(vm.stack); l != 1 {
		t.Error("expected vm.stack to have 1 element, but it has a length of ", l)
	}

	vm.stack.push(stackVal(1))
	if l := len(vm.stack); l != 2 {
		t.Error("expected vm.stack to have 2 elements, but it has a length of ", l)
	}
}

func TestPop(t *testing.T) {
	vm := NewVM()

	var expected stackVal = 1
	vm.stack = append(vm.stack, expected)
	v := vm.stack.pop()
	if v != expected {
		t.Errorf("expected vm.stack.pop to return %d, but got %d", expected, v)
	}
	if l := len(vm.stack); l != 0 {
		t.Error("expected stack to be empty after calling vm.stack.pop, but vm.stack has a length of ", l)
	}

	var expected2 stackVal = 42
	vm.stack = append(vm.stack, expected)
	vm.stack = append(vm.stack, expected)
	vm.stack = append(vm.stack, expected2)
	v2 := vm.stack.pop()
	if v2 != expected2 {
		t.Error("expected vm.stack.pop to return the last element on the stack, but got ", v)
	}
}

func TestStackOverflow(t *testing.T) {
	vm := NewVM()

	for i := 0; i < maxStackSize; i++ {
		vm.Interpret(instructions{instLiteral, 0x2A})
	}

	called := false
	fatal = func(_ string) {
		called = true
	}
	// Cause a stack overflow, which will result in fatal being called
	vm.stack.push(0x2A)
	if called == false {
		t.Error("Adding too many elements to the stack did not cause a stack overflow")
	}
}
