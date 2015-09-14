package enocean

import (
	"testing"
)

func TestSampleTelegrams(t *testing.T) {
	testVector := [][]byte{
		// trivial cases
		{0, 0},
		{0, 0, 0},
		{0, 0, 0, 0, 0x00, 0, 0},

		// EnOcean Serial Protocol header without sync byte
		{0, 7, 2, 10, 10},
		{0, 10, 2, 10, 0x9b},
		{0, 10, 2, 10, 155},
		{0, 12, 2, 10, 230},

		// EnOcean Serial Protocol data samples
		// wrong sample?{0x20, 0, 0x29, 0x91, 0xf1, 0x88, 0, 2, 0x2a, 0xfc},
		{0x22, 4, 0, 0x1d, 0x6e, 0, 0, 0x4b, 8, 0xe5, 3, 0x31, 0xa8},
		{32, 0, 41, 145, 150, 0, 14, 1, 41, 202},
		{32, 0, 41, 145, 150, 130, 137, 1, 41, 202},
		{32, 0, 41, 145, 150, 136, 191, 1, 41, 202},

		// EnOcean Radio Protocol Data CRC
		{98, 0, 0, 4, 0, 87, 200, 8, 40, 11, 128, 131},
	}

	for _, vec := range testVector {
		sample := vec[:len(vec)-1]
		if CRC8(sample) != vec[len(vec)-1] {
			t.Errorf("%v caused %x shall be %x", sample, CRC8(sample), vec[len(vec)-1])
		}
	}
}
