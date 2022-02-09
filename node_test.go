package circuit

import "testing"

func TestConstNode(t *testing.T) {
	if EvalBit(High) != HIGH {
		t.Error("ConstNode High evalulate error")
	}

	if EvalBit(Low) != LOW {
		t.Error("ConstNode Low evalulate error")
	}

	if EvalBit(GetConstNode(HIGH)) != HIGH {
		t.Error("GetConstNode(HIGH) error")
	}

	if EvalBit(GetConstNode(LOW)) != LOW {
		t.Error("GetConstNode(HIGH) error")
	}

	if s := High.String(); s != "High" {
		t.Errorf("High.String() is %q, but \"High\" expected", s)
	}

	if s := Low.String(); s != "Low" {
		t.Errorf("Low.String() is %q, but \"Low\" expected", s)
	}
}
