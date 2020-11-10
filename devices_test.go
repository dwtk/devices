package devices

import (
	"testing"
)

func TestGetBySignature(t *testing.T) {
	mcu, err := GetBySignature(0x910b)
	if err != nil {
		t.Error(err)
	}

	if mcu == nil {
		t.Errorf("MCU not found for 0x910b")
	}

	if m := mcu.Name(); m != "ATtiny24" {
		t.Errorf("Bad MCU name: %s", m)
	}

	if m := mcu.EECR().Io8(); m != 0x1c {
		t.Errorf("Bad EECR.Io8(): 0x%02x", m)
	}

	if m := mcu.EECR().Io16(); m != 0x1c {
		t.Errorf("Bad EECR.Io16(): 0x%04x", m)
	}

	if m := mcu.EECR().Mem8(); m != 0x3c {
		t.Errorf("Bad EECR.Mem8(): 0x%02x", m)
	}

	if m := mcu.EECR().Mem16(); m != 0x3c {
		t.Errorf("Bad EECR.Mem16(): 0x%04x", m)
	}
}

func TestGetBySignatureError(t *testing.T) {
	mcu, err := GetBySignature(0xdead)
	if err == nil {
		t.Errorf("Error was expected")
	}

	if mcu != nil {
		t.Errorf("Returned MCU for invalid signature: %s", mcu.Name())
	}

	if e := err.Error(); e != "devices: MCU lookup failed for signature: 0xdead" {
		t.Errorf("Bad error message: %s", e)
	}
}

func TestGetByName(t *testing.T) {
	mcu, err := GetByName("attiny24a")
	if err != nil {
		t.Error(err)
	}

	if mcu == nil {
		t.Errorf("MCU not found for 0x910b")
	}

	if m := mcu.Name(); m != "ATtiny24A" {
		t.Errorf("Bad MCU name: %s", m)
	}

	if m := mcu.EECR().Io8(); m != 0x1c {
		t.Errorf("Bad EECR.Io8(): 0x%02x", m)
	}

	if m := mcu.EECR().Io16(); m != 0x1c {
		t.Errorf("Bad EECR.Io16(): 0x%04x", m)
	}

	if m := mcu.EECR().Mem8(); m != 0x3c {
		t.Errorf("Bad EECR.Mem8(): 0x%02x", m)
	}

	if m := mcu.EECR().Mem16(); m != 0x3c {
		t.Errorf("Bad EECR.Mem16(): 0x%04x", m)
	}
}

func TestGetByNameError(t *testing.T) {
	mcu, err := GetByName("attiny345")
	if err == nil {
		t.Errorf("Error was expected")
	}

	if mcu != nil {
		t.Errorf("Returned MCU for invalid name: %s", mcu.Name())
	}

	if e := err.Error(); e != "devices: MCU lookup failed for name: attiny345" {
		t.Errorf("Bad error message: %s", e)
	}
}
