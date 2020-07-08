package devices

type SFR struct {
	address byte
	size    byte
}

func (s SFR) Io8() byte {
	return s.address - 0x20
}

func (s SFR) Io16() uint16 {
	return uint16(s.address - 0x20)
}

func (s SFR) Mem8() byte {
	return s.address
}

func (s SFR) Mem16() uint16 {
	return uint16(s.address)
}

func (s SFR) Size() byte {
	return s.size
}
