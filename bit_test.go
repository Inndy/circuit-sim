package circuit

import "testing"

func TestHighLowConst(t *testing.T) {
	if LOW {
		t.Error("LOW definiation error")
	}

	if !HIGH {
		t.Error("HIGH definiation error")
	}
}
