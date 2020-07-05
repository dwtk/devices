package devices

//go:generate go run generate.go

import (
	"fmt"
)

type SFR uint16

func (s SFR) Io8() byte {
	return byte(uint16(s) - 0x20)
}

func (s SFR) Io16() uint16 {
	return uint16(s) - 0x20
}

func (s SFR) Mem8() byte {
	return byte(s)
}

func (s SFR) Mem16() uint16 {
	return uint16(s)
}

type MCU struct {
	Name          string
	Signature     uint16
	FlashPageSize uint16
	FlashSize     uint16
	EEPROMSize    uint16
	EECR          SFR
	WithEEARH     bool
	DWENMask      byte
}

func GetMCU(signature uint16) (*MCU, error) {
	for _, mcu := range mcus {
		if signature == mcu.Signature {
			return mcu, nil
		}
	}
	return nil, fmt.Errorf("devices: MCU lookup failed for signature: 0x%04x", signature)
}
