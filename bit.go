package circuit

type Bit bool

func (b Bit) String() string {
	if bool(b) {
		return "1"
	} else {
		return "0"
	}
}

const (
	LOW Bit = false
	HIGH Bit = true
)
