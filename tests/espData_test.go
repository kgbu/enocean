package main

import (
	"testing"
	"github.com/kgbu/enocean"
)

func TestVariousPacketAccept(t *testing.T) {
	patterns := [][]byte{
		{85, 0, 12, 2, 10, 230, 98, 0, 0, 4, 0, 87, 200, 8, 40, 11, 128, 131, 1, 53, 158},
		{85, 0, 12, 2, 10, 230, 98, 0, 0, 4, 0, 87, 200, 8, 40, 11, 128, 131, 1, 55, 144},
		{85, 0, 10, 2, 10, 155, 34, 4, 0, 87, 200, 0, 0, 80, 8, 246, 1, 41, 202},
		{85, 0, 10, 2, 10, 155, 34, 4, 0, 87, 200, 0, 0, 80, 8, 246, 1, 55, 144},
		{85, 0, 12, 2, 10, 230, 98, 0, 0, 4, 1, 121, 77, 16, 8, 11, 128, 48, 1, 41, 202},
		{85, 0, 10, 2, 10, 155, 34, 4, 1, 121, 77, 0, 166, 192, 10, 11, 1, 45, 214},
		{85, 0, 10, 2, 10, 155, 34, 4, 1, 121, 77, 0, 197, 205, 10, 154, 1, 45, 214},
		{85, 0, 10, 2, 10, 155, 34, 4, 1, 121, 77, 0, 180, 176, 10, 221, 1, 43, 196},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 132, 242, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 136, 214, 1, 73, 237},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 61, 166},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 132, 242, 1, 46, 223},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 45, 214},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 136, 214, 1, 42, 195},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 136, 214, 1, 42, 195},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 136, 214, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 136, 214, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 132, 242, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 132, 242, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 132, 242, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 132, 242, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 0, 43, 146, 87, 0, 103, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 4, 0, 62, 127, 0, 24, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 4, 0, 62, 127, 0, 24, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 4, 0, 62, 127, 128, 145, 1, 41, 202},
		{85, 0, 7, 2, 10, 10, 32, 4, 0, 62, 127, 128, 145, 1, 41, 202},
	}

	for _, v := range patterns {
		err, consumedBytes, _ := enocean.NewESPData(v)
		if err != nil {
			t.Errorf("ERROR: %v, parse failed on %v. consumed %v, ", err, v, consumedBytes)
		}
	}
}

func TestVariousRORGTeachInPacketAccept(t *testing.T) {
	patterns := [][]byte{
		{85, 0, 12, 2, 10, 230, 98, 0, 0, 4, 0, 87, 200, 8, 40, 11, 128, 131, 1, 53, 158},
		{85, 0, 12, 2, 10, 230, 98, 0, 0, 4, 1, 121, 77, 16, 8, 11, 128, 48, 1, 41, 202},
		{85, 0, 10, 2, 10, 155, 34, 4, 0, 79, 66, 0, 0, 62, 8, 78, 1, 55, 144, 10},	// not TeachIn case
	}

	validity := []bool{
		true,
		true,
		false,
	}

	for i, v := range patterns {
		err, consumedBytes, e := enocean.NewESPData(v)
		if err != nil {
			t.Errorf("ERROR: %v, parse failed on %v. consumed %v, ", err, v, consumedBytes)
		}
		if (!validity[i]) {
			if (e.TeachIn == true) {
				t.Errorf("ERROR: %v, not teach in data is flagged as teach-in at %v", e, i)
			} else {
				continue
			}
		}
		if (e.TeachIn == false) ||
			(e.RORG == 0) ||
			(e.FUNC == 0) ||
			(e.TYPE == 0) ||
			(e.ManufacturerId == 0) {
			t.Errorf("ERROR: %v, teach in data extraction failed on %v. consumed %v, ", e, v, consumedBytes)
		}
	}
}

func TestRSSIvalue(t *testing.T) {
	patterns := [][]byte{
		{85, 0, 12, 2, 10, 230, 98, 0, 0, 4, 0, 87, 200, 8, 40, 11, 128, 131, 1, 53, 158},
		{85, 0, 12, 2, 10, 230, 98, 0, 0, 4, 1, 121, 77, 16, 8, 11, 128, 48, 1, 41, 202},
	}

	for _, v := range patterns {
		err, consumedBytes, e := enocean.NewESPData(v)
		if err != nil {
			t.Errorf("ERROR: %v, parse failed on %v. consumed %v, ", err, v, consumedBytes)
		}
		if e.RSSI != v[len(v)-2] {
			t.Errorf("ERROR: RSSI data %v is differ from %v on %v. consumed %v, ", e.RSSI, v[len(v)-2], v, consumedBytes)
		}
	}
}
