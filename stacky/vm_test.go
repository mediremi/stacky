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
