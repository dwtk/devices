package devices

type SFR struct {
	address uint16
	size    byte
}

func (s SFR) Io8() byte {
	return byte(s.address - 0x20)
}

func (s SFR) Io16() uint16 {
	return s.address - 0x20
}

func (s SFR) Mem8() byte {
	return byte(s.address)
}

func (s SFR) Mem16() uint16 {
	return s.address
}

func (s SFR) Size() byte {
	return s.size
}
