package main

import (
	"fmt"

	netpbm "github.com/lounabrtt/Netpbm"
)

func main() {
	filename := "imageP4.pbm"
	pbm, err := netpbm.ReadPBM(filename)
	if err != nil {
		fmt.Println("error", err)
		return
	}

		// stets the magic number
		//fmt.Printf("Current Magic Number: %s\n", pbm.MagicNumber)
	//pbm.SetMagicNumber("P3")
		//fmt.Printf("Modified Magic Number: %s\n", pbm.MagicNumber)

	fmt.Printf("Magic Number: %s\n", pbm.MagicNumber)
	fmt.Printf("Width: %d\n", pbm.Width)
	fmt.Printf("Height: %d\n", pbm.Height)
	fmt.Println("Data before:")
	fmt.Println("here it's ", pbm.At(0, 0))

	//pbm.Invert()
	//pbm.Flip()


	// change the value of the pixel
	pbm.Set(0, 0, true)
	fmt.Println("Data after:")
	fmt.Println("here it's ", pbm.At(0, 0))

	//convert from number to pixel
	for _, row := range pbm.Data {
		for _, pixel := range row {
			if pixel {
				fmt.Print("°")
			} else {
				fmt.Print("■")
			}
		}
		fmt.Println()
	}
	// new save
	newFilename := "imagehexa_two.pbm"
	err = pbm.Save(newFilename)
	if err != nil {
		fmt.Println("error saving image:", err)
	}
}
