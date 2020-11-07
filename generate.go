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
	nrwwstart     uint16
	eepromsize    uint16
	eecr          byte
	witheearh     bool
	withsph       bool
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
		nrwwstart := uint64(0)
		for _, bootsection := range bootsections {
			start, err := strconv.ParseUint(bootsection.SelectAttr("start"), 0, 16)
			if err != nil {
				continue
			}
			if nrwwstart == 0 || start < nrwwstart {
				nrwwstart = start
			}
		}

		eeprom := xmlquery.FindOne(device, "//address-space[@name='eeprom']/memory-segment[@name='EEPROM']")
		eepromsize, err := strconv.ParseUint(eeprom.SelectAttr("size"), 0, 16)
		if err != nil {
			return nil
		}

		eeprommodule := xmlquery.FindOne(doc, "//module[@name='EEPROM']/register-group[@name='EEPROM']")

		eecrreg := xmlquery.FindOne(eeprommodule, "register[@name='EECR']")
		eecr, err := strconv.ParseUint(eecrreg.SelectAttr("offset"), 0, 8)
		if err != nil {
			return nil
		}

		witheearh := false
		if eearreg := xmlquery.FindOne(eeprommodule, "register[@name='EEAR']"); eearreg != nil {
			witheearh = eearreg.SelectAttr("size") == "2"
		} else {
			witheearh = xmlquery.FindOne(eeprommodule, "register[@name='EEARH']") != nil
		}

		withsph := false
		if spreg := xmlquery.FindOne(doc, "//register[@name='SP']"); spreg != nil {
			withsph = spreg.SelectAttr("size") == "2"
		} else {
			withsph = xmlquery.FindOne(doc, "//register[@name='SPH']") != nil
		}

		dwen := xmlquery.FindOne(doc, "//bitfield[@name='DWEN']")
		dwenmask, err := strconv.ParseUint(dwen.SelectAttr("mask"), 0, 8)
		if err != nil {
			return nil
		}

		devices = append(devices, deviceType{
			name:          name,
			signature:     signature,
			flashpagesize: uint16(flashpagesize),
			flashsize:     uint16(flashsize),
			nrwwstart:     uint16(nrwwstart),
			eepromsize:    uint16(eepromsize),
			eecr:          byte(eecr),
			witheearh:     witheearh,
			withsph:       withsph,
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
		fmt.Fprintf(f, "0x%04x, ", device.nrwwstart)
		fmt.Fprintf(f, "0x%04x, ", device.eepromsize)
		fmt.Fprintf(f, "0x%02x, ", device.eecr)
		fmt.Fprintf(f, "%v, ", device.witheearh)
		fmt.Fprintf(f, "%v, ", device.withsph)
		fmt.Fprintf(f, "%s", formatBitmask(device.dwenmask))
		fmt.Fprintf(f, "},\n")
	}
	fmt.Fprintf(f, "}\n")
}
