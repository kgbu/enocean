package main

import (
	"fmt"

	"github.com/kgbu/enocean"
)

func main() {

	fmt.Println("Data packets")

	A5_02_05pattern := []byte{85, 0, 10, 2, 10, 155, 34, 4, 0, 87, 200, 0, 0, 76, 8, 93, 1, 50, 139}
	A5_04_01pattern := []byte{85, 0, 10, 2, 10, 155, 34, 4, 1, 121, 77, 0, 166, 192, 10, 11, 1, 45, 214}

	// A5_02_05 Temperture
	err, c, e := enocean.NewESPData(A5_02_05pattern)
	if err != nil {
		fmt.Printf("ERROR: %v, parse failed on %v. consumed %v", err, A5_02_05pattern, c)
	}

	temp := (255 - int(e.PayloadData[2])) * 40.0 / 255
	fmt.Printf("%v : temperature %v\n", e, temp)

	// A5_04_01 Humidity/Temperture
	err, c, e2 := enocean.NewESPData(A5_04_01pattern)
	if err != nil {
		fmt.Printf("ERROR: %v, parse failed on %v. consumed %v", err, A5_04_01pattern, c)
	}

	humid := int(e2.PayloadData[1]) * 100.0 / 250
	temp2 := int(e2.PayloadData[2]) * 40.0 / 250

	fmt.Printf("%v : humidity: %v, temperature: %v\n", e2, humid, temp2)
}
