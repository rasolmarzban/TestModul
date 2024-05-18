package main

import (
	"fmt"
	"github.com/jalaali/go-jalaali"
)

func main() {

	year, mounth, day, error := jalaali.ToGregorian(1402, 14, 12)

	if error == nil {
		fmt.Printf("%v/%v/%v\n", year, mounth, day)
	} else {
		fmt.Printf("%v\n", error)
	}
}
