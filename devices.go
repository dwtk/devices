package devices

//go:generate go run generate.go

import (
	"fmt"
)

type MCU struct {
	Name          string
	Signature     uint16
	FlashPageSize uint16
	FlashSize     uint16
	EEPROMSize    uint16
	EECR          byte
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
