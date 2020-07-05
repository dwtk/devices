package devices

import (
	"testing"
)

func TestGetMCU(t *testing.T) {
	mcu, err := GetMCU(0x910b)
	if err != nil {
		t.Error(err)
	}

	if mcu == nil {
		t.Errorf("MCU not found for 0x910b")
	}

	if mcu.Name != "ATtiny24" {
		t.Errorf("Bad MCU name: %s", mcu.Name)
	}
}

func TestGetMCUError(t *testing.T) {
	mcu, err := GetMCU(0xdead)
	if err == nil {
		t.Errorf("Error was expected")
	}

	if mcu != nil {
		t.Errorf("Returned MCU for invalid signature: %s", mcu.Name)
	}

	if err.Error() != "devices: MCU lookup failed for signature: 0xdead" {
		t.Errorf("Bad error message: %s", err)
	}
}
