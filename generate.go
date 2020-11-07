// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
)

type deviceType struct {
	name          string
	signature     uint16
	flashpagesize uint16
	flashsize     uint16
	nrwwoffset    uint16
	eepromsize    uint16
	eear          byte
	eearsize      byte
	eedr          byte
	eedrsize      byte
	eecr          byte
	eecrsize      byte
	spmcsr        byte
	spmcsrsize    byte
	sp            byte
	spsize        byte
	sreg          byte
	sregsize      byte
	dwenmask      byte
}

func findDevices() ([]deviceType, error) {
	devices := []deviceType{}

	err := filepath.Walk("atdf", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".atdf") {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		doc, err := xmlquery.Parse(f)
		if err != nil {
			return err
		}

		device := xmlquery.FindOne(doc, "//device")
		if device == nil {
			return nil
		}

		if xmlquery.FindOne(device, "//interface[@type='dw']") == nil {
			return nil
		}

		name := device.SelectAttr("name")

		signature := uint16(0)
		for _, prop := range xmlquery.Find(device, "//property-group[@name='SIGNATURES']/property") {
			n := prop.SelectAttr("name")
			if n == "SIGNATURE1" {
				v, _ := strconv.ParseUint(prop.SelectAttr("value"), 0, 8)
				signature |= uint16(v) << 8
			} else if n == "SIGNATURE2" {
				v, _ := strconv.ParseUint(prop.SelectAttr("value"), 0, 8)
				signature |= uint16(v)
			}
		}

		flash := xmlquery.FindOne(device, "//address-space[@name='prog']/memory-segment[@name='FLASH']")

		flashpagesize, err := strconv.ParseUint(flash.SelectAttr("pagesize"), 0, 16)
		if err != nil || flashpagesize > 0x80 {
			return nil
		}

		flashsize, err := strconv.ParseUint(flash.SelectAttr("size"), 0, 16)
		if err != nil || flashsize > 0xffff {
			return nil
		}

		bootsections := xmlquery.Find(device, "//address-space[@name='prog']/memory-segment[starts-with(@name, 'BOOT_SECTION_')]")
		nrwwoffset := uint64(0)
		for _, bootsection := range bootsections {
			start, err := strconv.ParseUint(bootsection.SelectAttr("start"), 0, 16)
			if err != nil {
				continue
			}
			if nrwwoffset == 0 || start < nrwwoffset {
				nrwwoffset = start
			}
		}

		eeprom := xmlquery.FindOne(device, "//address-space[@name='eeprom']/memory-segment[@name='EEPROM']")
		eepromsize, err := strconv.ParseUint(eeprom.SelectAttr("size"), 0, 16)
		if err != nil {
			return err
		}

		eeprommodule := xmlquery.FindOne(doc, "//module[@name='EEPROM']/register-group[@name='EEPROM']")

		eear := uint64(0)
		eearsize := uint64(0)
		if eearreg := xmlquery.FindOne(eeprommodule, "register[@name='EEAR']"); eearreg != nil {
			eear, err = strconv.ParseUint(eearreg.SelectAttr("offset"), 0, 8)
			if err != nil {
				return err
			}
			eearsize, err = strconv.ParseUint(eearreg.SelectAttr("size"), 0, 8)
			if err != nil {
				return err
			}
		} else if eearlreg := xmlquery.FindOne(eeprommodule, "register[@name='EEARL']"); eearlreg != nil {
			eear, err = strconv.ParseUint(eearlreg.SelectAttr("offset"), 0, 8)
			if err != nil {
				return err
			}
			eearsize, err = strconv.ParseUint(eearlreg.SelectAttr("size"), 0, 8)
			if err != nil {
				return err
			}
			if eearhreg := xmlquery.FindOne(eeprommodule, "register[@name='EEARH']"); eearhreg != nil {
				size, err := strconv.ParseUint(eearhreg.SelectAttr("size"), 0, 8)
				if err != nil {
					return err
				}
				eearsize += size
			}
		} else {
			return fmt.Errorf("failed to find EEAR")
		}

		eedrreg := xmlquery.FindOne(eeprommodule, "register[@name='EEDR']")
		eedr, err := strconv.ParseUint(eedrreg.SelectAttr("offset"), 0, 8)
		if err != nil {
			return err
		}
		eedrsize, err := strconv.ParseUint(eedrreg.SelectAttr("size"), 0, 8)
		if err != nil {
			return err
		}

		eecrreg := xmlquery.FindOne(eeprommodule, "register[@name='EECR']")
		eecr, err := strconv.ParseUint(eecrreg.SelectAttr("offset"), 0, 8)
		if err != nil {
			return err
		}
		eecrsize, err := strconv.ParseUint(eecrreg.SelectAttr("size"), 0, 8)
		if err != nil {
			return err
		}

		spmcsr := uint64(0)
		spmcsrsize := uint64(0)
		if spmcsrreg := xmlquery.FindOne(doc, "//register[@name='SPMCSR']"); spmcsrreg != nil {
			spmcsr, err = strconv.ParseUint(spmcsrreg.SelectAttr("offset"), 0, 8)
			if err != nil {
				return err
			}
			spmcsrsize, err = strconv.ParseUint(spmcsrreg.SelectAttr("size"), 0, 8)
			if err != nil {
				return err
			}
		} else if spmcrreg := xmlquery.FindOne(doc, "//register[@name='SPMCR']"); spmcrreg != nil {
			spmcsr, err = strconv.ParseUint(spmcrreg.SelectAttr("offset"), 0, 8)
			if err != nil {
				return err
			}
			spmcsrsize, err = strconv.ParseUint(spmcrreg.SelectAttr("size"), 0, 8)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("failed to find SPMCSR")
		}

		sp := uint64(0)
		spsize := uint64(0)
		if spreg := xmlquery.FindOne(doc, "//register[@name='SP']"); spreg != nil {
			sp, err = strconv.ParseUint(spreg.SelectAttr("offset"), 0, 8)
			if err != nil {
				return err
			}
			spsize, err = strconv.ParseUint(spreg.SelectAttr("size"), 0, 8)
			if err != nil {
				return err
			}
		} else if splreg := xmlquery.FindOne(doc, "//register[@name='SPL']"); splreg != nil {
			sp, err = strconv.ParseUint(splreg.SelectAttr("offset"), 0, 8)
			if err != nil {
				return err
			}
			spsize, err = strconv.ParseUint(splreg.SelectAttr("size"), 0, 8)
			if err != nil {
				return err
			}
			if sphreg := xmlquery.FindOne(doc, "//register[@name='SPH']"); sphreg != nil {
				size, err := strconv.ParseUint(sphreg.SelectAttr("size"), 0, 8)
				if err != nil {
					return err
				}
				spsize += size
			}
		} else {
			return fmt.Errorf("failed to find SP")
		}

		sregreg := xmlquery.FindOne(doc, "//register[@name='SREG']")
		sreg, err := strconv.ParseUint(sregreg.SelectAttr("offset"), 0, 8)
		if err != nil {
			return err
		}
		sregsize, err := strconv.ParseUint(sregreg.SelectAttr("size"), 0, 8)
		if err != nil {
			return err
		}

		dwen := xmlquery.FindOne(doc, "//bitfield[@name='DWEN']")
		dwenmask, err := strconv.ParseUint(dwen.SelectAttr("mask"), 0, 8)
		if err != nil {
			return err
		}

		devices = append(devices, deviceType{
			name:          name,
			signature:     signature,
			flashpagesize: uint16(flashpagesize),
			flashsize:     uint16(flashsize),
			nrwwoffset:    uint16(nrwwoffset),
			eepromsize:    uint16(eepromsize),
			eear:          byte(eear),
			eearsize:      byte(eearsize),
			eedr:          byte(eedr),
			eedrsize:      byte(eedrsize),
			eecr:          byte(eecr),
			eecrsize:      byte(eecrsize),
			spmcsr:        byte(spmcsr),
			spmcsrsize:    byte(spmcsrsize),
			sp:            byte(sp),
			spsize:        byte(spsize),
			sreg:          byte(sreg),
			sregsize:      byte(sregsize),
			dwenmask:      byte(dwenmask),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func formatBitmask(b byte) string {
	f := -1

	for i := 0; i < 8; i++ {
		if (1<<i)&b != 0 {
			f = i
		}
	}

	if f == -1 {
		return "0"
	}

	return fmt.Sprintf("1 << %d", f)
}

func main() {
	devices, err := findDevices()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("mcus.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "// Code generated by generate.go. DO NOT EDIT.\n\n")
	fmt.Fprintf(f, "package devices\n\nvar mcus = []*MCU{\n")
	for _, device := range devices {
		fmt.Fprintf(f, "\t{")
		fmt.Fprintf(f, "\"%s\", ", device.name)
		fmt.Fprintf(f, "0x%04x, ", device.signature)
		fmt.Fprintf(f, "0x%02x, ", device.flashpagesize)
		fmt.Fprintf(f, "0x%04x, ", device.flashsize)
		fmt.Fprintf(f, "0x%04x, ", device.nrwwoffset)
		fmt.Fprintf(f, "0x%04x, ", device.eepromsize)
		fmt.Fprintf(f, "0x%02x, ", device.eear)
		fmt.Fprintf(f, "%d, ", device.eearsize)
		fmt.Fprintf(f, "0x%02x, ", device.eedr)
		fmt.Fprintf(f, "%d, ", device.eedrsize)
		fmt.Fprintf(f, "0x%02x, ", device.eecr)
		fmt.Fprintf(f, "%d, ", device.eecrsize)
		fmt.Fprintf(f, "0x%02x, ", device.spmcsr)
		fmt.Fprintf(f, "%d, ", device.spmcsrsize)
		fmt.Fprintf(f, "0x%02x, ", device.sp)
		fmt.Fprintf(f, "%d, ", device.spsize)
		fmt.Fprintf(f, "0x%02x, ", device.sreg)
		fmt.Fprintf(f, "%d, ", device.sregsize)
		fmt.Fprintf(f, "%s", formatBitmask(device.dwenmask))
		fmt.Fprintf(f, "},\n")
	}
	fmt.Fprintf(f, "}\n")
}
