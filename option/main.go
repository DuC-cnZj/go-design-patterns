package main

import (
	"design_patterns/option/option"
	"fmt"
)

func main()  {
	options := option.New(option.SetAge(21), option.SetId(1), option.SetName("duc"))
	fmt.Println(options)
}