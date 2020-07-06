package devices

//go:generate go run generate.go

import (
	"fmt"
)

type MCU struct {
	name          string
	signature     uint16
	flashPageSize uint16
	flashSize     uint16
	nrwwStart     uint16
	eepromSize    uint16
	eecr          byte
	withEEARH     bool
	withSPH       bool
	dwenMask      byte
	withRWW       bool
}

func GetMCU(signature uint16) (*MCU, error) {
	for _, mcu := range mcus {
		if signature == mcu.signature {
			return mcu, nil
		}
	}
	return nil, fmt.Errorf("devices: MCU lookup failed for signature: 0x%04x", signature)
}

func (m *MCU) Name() string {
	return m.name
}

func (m *MCU) Signature() uint16 {
	return m.signature
}

func (m *MCU) FlashPageSize() uint16 {
	return m.flashPageSize
}

func (m *MCU) FlashSize() uint16 {
	return m.flashSize
}

func (m *MCU) NRWWStart() uint16 {
	return m.nrwwStart
}

func (m *MCU) EEPROMSize() uint16 {
	return m.eepromSize
}

func (m *MCU) EECR() SFR {
	return SFR{uint16(m.eecr), 1}
}

func (m *MCU) EEDR() SFR {
	return SFR{uint16(m.eecr) + 1, 1}
}

func (m *MCU) EEAR() SFR {
	s := byte(1)
	if m.withEEARH {
		s++
	}
	return SFR{uint16(m.eecr) + 2, s}
}

func (m *MCU) SPMCSR() SFR {
	return SFR{0x57, 1}
}

func (m *MCU) SP() SFR {
	s := byte(1)
	if m.withSPH {
		s++
	}
	return SFR{0x5d, s}
}

func (m *MCU) SREG() SFR {
	return SFR{0x5f, 1}
}

func (m *MCU) DWENMask() byte {
	return m.dwenMask
}

func (m *MCU) WithRWW() bool {
	return m.withRWW
}
