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
	nrwwOffset    uint16
	eepromSize    uint16
	eear          byte
	eearSize      byte
	eedr          byte
	eedrSize      byte
	eecr          byte
	eecrSize      byte
	spmcsr        byte
	spmcsrSize    byte
	sp            byte
	spSize        byte
	sreg          byte
	sregSize      byte
	dwenMask      byte
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

func (m *MCU) NRWWOffset() uint16 {
	return m.nrwwOffset
}

func (m *MCU) EEPROMSize() uint16 {
	return m.eepromSize
}

func (m *MCU) EEAR() SFR {
	return SFR{m.eear, m.eearSize}
}

func (m *MCU) EEDR() SFR {
	return SFR{m.eedr, m.eedrSize}
}

func (m *MCU) EECR() SFR {
	return SFR{m.eecr, m.eecrSize}
}

func (m *MCU) SPMCSR() SFR {
	return SFR{m.spmcsr, m.spmcsrSize}
}

func (m *MCU) SP() SFR {
	return SFR{m.sp, m.spSize}
}

func (m *MCU) SREG() SFR {
	return SFR{m.sreg, m.sregSize}
}

func (m *MCU) DWENMask() byte {
	return m.dwenMask
}
