package main

import (
	"flag"
	"fmt"
)

// baseFontSize is the base font size in pixels
var baseFontSize int;

/* 
A function to convert a pixel value to rem

@param px int - The pixel value to convert
@param baseFontSize int - The base font size in pixels

@return float64 - The converted value in rem
*/
func pxToRem(px int, baseFontSize int) float64 {
	return float64(px) / float64(baseFontSize)
}

/* 
A function to convert a rem value to pixels

@param rem float64 - The rem value to convert
@param baseFontSize int - The base font size in pixels

@return int - The converted value in pixels
*/
func remToPx(rem float64, baseFontSize int) int {
	return int(rem * float64(baseFontSize))
}

/*
The main function
*/
func main() {
	// Parse the command line flags
	pxValue := flag.Int("px", 0, "Pixel value to convert to rem")
	remValue := flag.Float64("rem", 0, "Rem value to convert to px")
	flag.IntVar(&baseFontSize, "baseFontSize", 16, "Base font size")
	flag.Parse()

	// Check if the user provided a valid input
	switch {
	case *pxValue > 0:
		fmt.Println(pxToRem(*pxValue, baseFontSize))
	case *remValue > 0:
		fmt.Println(remToPx(*remValue, baseFontSize))
	default:
		fmt.Println("Please provide a valid input")
	}
}