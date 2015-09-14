package main

import (
	"fmt"

	"github.com/kgbu/enocean"
)

func main() {

	fmt.Println("Teach In packets")

	patterns := [][]byte{
		{85, 0, 12, 2, 10, 230, 98, 0, 0, 4, 0, 87, 200, 8, 40, 11, 128, 131, 1, 53, 158},
		{85, 0, 12, 2, 10, 230, 98, 0, 0, 4, 1, 121, 77, 16, 8, 11, 128, 48, 1, 41, 202},
	}

	for _, v := range patterns {
		err, consumedBytes, e := enocean.NewESPData(v)
		if err != nil {
			fmt.Printf("ERROR: %v, parse failed on %v. consumed %v, ", err, v, consumedBytes)
		}
		if (e.TeachIn == false) ||
			(e.RORG == 0) ||
			(e.FUNC == 0) ||
			(e.TYPE == 0) ||
			(e.ManufacturerId == 0) {
			fmt.Printf("ERROR: %v, teach in data extraction failed on %v. consumed %v, ", e, v, consumedBytes)
			return
		}
		fmt.Printf("\nRaw data: %v\n", e)
		err, m := enocean.GetManufacturerName(e.ManufacturerId)
		if err != nil {
			fmt.Printf("ERROR: %v, manufacturerID is wrong", e.ManufacturerId)
			return
		}
		fmt.Printf("RORG: %x, FUNC: %v, TYPE: %v ManufacturerId: %v\n", e.RORG, e.FUNC, e.TYPE, m)
	}
}
