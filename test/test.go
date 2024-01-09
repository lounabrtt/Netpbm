package main

import (
	"fmt"

	netpbm "github.com/lounabrtt/Netpbm"
)

func main() {
	filename := "image1.pbm"
	pbm, err := netpbm.ReadPBM(filename)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Printf("Magic Number: %s\n", pbm.MagicNumber)
	fmt.Printf("Width: %d\n", pbm.Width)
	fmt.Printf("Height: %d\n", pbm.Height)
	fmt.Println("Data:")

	for _, row := range pbm.Data {
		for _, pixel := range row {
			if pixel {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println()
	}
}
