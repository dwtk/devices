// Code generated by generate.go. DO NOT EDIT.

package devices

var mcus = []*MCU{
	{"AT90PWM1", 0x9383, 0x40, 0x2000, 0x1800, 0x0200, 0x3f, true, true, 1 << 6, true},
	{"AT90PWM161", 0x948b, 0x80, 0x4000, 0x3800, 0x0200, 0x3c, true, true, 1 << 6, true},
	{"AT90PWM216", 0x9483, 0x80, 0x4000, 0x3000, 0x0200, 0x3f, true, true, 1 << 6, true},
	{"AT90PWM81", 0x9388, 0x40, 0x2000, 0x1800, 0x0200, 0x3c, true, true, 1 << 6, true},
	{"AT90USB162", 0x9482, 0x80, 0x4000, 0x3000, 0x0200, 0x3f, true, true, 1 << 7, true},
	{"AT90USB82", 0x9382, 0x80, 0x2000, 0x1000, 0x0200, 0x3f, true, true, 1 << 7, true},
	{"ATmega168", 0x9406, 0x80, 0x4000, 0x3800, 0x0200, 0x3f, true, true, 1 << 6, true},
	{"ATmega168P", 0x940b, 0x80, 0x4000, 0x3800, 0x0200, 0x3f, true, true, 1 << 6, true},
	{"ATmega168PB", 0x9415, 0x80, 0x4000, 0x3800, 0x0200, 0x3f, true, true, 1 << 6, true},
	{"ATmega16HVA", 0x940c, 0x80, 0x4000, 0x0000, 0x0100, 0x3f, false, true, 1 << 4, false},
	{"ATmega16HVB", 0x940d, 0x80, 0x4000, 0x3000, 0x0200, 0x3f, true, true, 1 << 3, true},
	{"ATmega16M1", 0x9484, 0x80, 0x4000, 0x3000, 0x0200, 0x3f, true, true, 1 << 6, true},
	{"ATmega16U2", 0x9489, 0x80, 0x4000, 0x3000, 0x0200, 0x3f, true, true, 1 << 7, true},
	{"ATmega328", 0x9514, 0x80, 0x8000, 0x7000, 0x0400, 0x3f, true, true, 1 << 6, true},
	{"ATmega328P", 0x950f, 0x80, 0x8000, 0x7000, 0x0400, 0x3f, true, true, 1 << 6, true},
	{"ATmega328PB", 0x9516, 0x80, 0x8000, 0x7000, 0x0400, 0x3f, true, true, 1 << 6, true},
	{"ATmega32C1", 0x9586, 0x80, 0x8000, 0x7000, 0x0400, 0x3f, true, true, 1 << 6, true},
	{"ATmega32HVB", 0x9510, 0x80, 0x8000, 0x7000, 0x0400, 0x3f, true, true, 1 << 3, true},
	{"ATmega32M1", 0x9584, 0x80, 0x8000, 0x7000, 0x0400, 0x3f, true, true, 1 << 6, true},
	{"ATmega32U2", 0x958a, 0x80, 0x8000, 0x7000, 0x0400, 0x3f, true, true, 1 << 7, true},
	{"ATmega48", 0x9205, 0x40, 0x1000, 0x0000, 0x0100, 0x3f, false, true, 1 << 6, false},
	{"ATmega48P", 0x920a, 0x40, 0x1000, 0x0000, 0x0100, 0x3f, false, true, 1 << 6, false},
	{"ATmega48PB", 0x9210, 0x40, 0x1000, 0x0000, 0x0100, 0x3f, false, true, 1 << 6, false},
	{"ATmega88", 0x930a, 0x40, 0x2000, 0x1800, 0x0200, 0x3f, true, true, 1 << 6, true},
	{"ATmega88P", 0x930f, 0x40, 0x2000, 0x1800, 0x0200, 0x3f, true, true, 1 << 6, true},
	{"ATmega88PB", 0x9316, 0x40, 0x2000, 0x1800, 0x0200, 0x3f, true, true, 1 << 6, true},
	{"ATmega8HVA", 0x9310, 0x80, 0x2000, 0x0000, 0x0100, 0x3f, false, true, 1 << 4, false},
	{"ATmega8U2", 0x9389, 0x80, 0x2000, 0x1000, 0x0200, 0x3f, true, true, 1 << 7, true},
	{"ATtiny13", 0x9007, 0x20, 0x0400, 0x0000, 0x0040, 0x3c, false, false, 1 << 3, false},
	{"ATtiny1634", 0x9412, 0x20, 0x4000, 0x0000, 0x0100, 0x3c, true, true, 1 << 6, false},
	{"ATtiny167", 0x9487, 0x80, 0x4000, 0x0000, 0x0200, 0x3f, true, true, 1 << 6, false},
	{"ATtiny2313", 0x910a, 0x20, 0x0800, 0x0000, 0x0080, 0x3c, false, false, 1 << 7, false},
	{"ATtiny24", 0x910b, 0x20, 0x0800, 0x0000, 0x0080, 0x3c, true, false, 1 << 6, false},
	{"ATtiny25", 0x9108, 0x20, 0x0800, 0x0000, 0x0080, 0x3c, true, false, 1 << 6, false},
	{"ATtiny261", 0x910c, 0x20, 0x0800, 0x0000, 0x0080, 0x3c, true, false, 1 << 6, false},
	{"ATtiny4313", 0x920d, 0x40, 0x1000, 0x0000, 0x0100, 0x3c, false, true, 1 << 7, false},
	{"ATtiny43U", 0x920c, 0x40, 0x1000, 0x0000, 0x0040, 0x3c, false, true, 1 << 6, false},
	{"ATtiny44", 0x9207, 0x40, 0x1000, 0x0000, 0x0100, 0x3c, true, true, 1 << 6, false},
	{"ATtiny441", 0x9215, 0x10, 0x1000, 0x0000, 0x0100, 0x3c, true, true, 1 << 6, false},
	{"ATtiny45", 0x9206, 0x40, 0x1000, 0x0000, 0x0100, 0x3c, true, true, 1 << 6, false},
	{"ATtiny461", 0x9208, 0x40, 0x1000, 0x0000, 0x0100, 0x3c, true, true, 1 << 6, false},
	{"ATtiny48", 0x9209, 0x40, 0x1000, 0x0000, 0x0040, 0x3f, false, true, 1 << 6, false},
	{"ATtiny828", 0x9314, 0x40, 0x2000, 0x1800, 0x0100, 0x3f, false, true, 1 << 6, true},
	{"ATtiny84", 0x930c, 0x40, 0x2000, 0x0000, 0x0200, 0x3c, true, true, 1 << 6, false},
	{"ATtiny841", 0x9315, 0x10, 0x2000, 0x0000, 0x0200, 0x3c, true, true, 1 << 6, false},
	{"ATtiny85", 0x930b, 0x40, 0x2000, 0x0000, 0x0200, 0x3c, true, true, 1 << 6, false},
	{"ATtiny861", 0x930d, 0x40, 0x2000, 0x0000, 0x0200, 0x3c, true, true, 1 << 6, false},
	{"ATtiny87", 0x9387, 0x80, 0x2000, 0x0000, 0x0200, 0x3f, true, true, 1 << 6, false},
	{"ATtiny88", 0x9311, 0x40, 0x2000, 0x0000, 0x0040, 0x3f, false, true, 1 << 6, false},
}