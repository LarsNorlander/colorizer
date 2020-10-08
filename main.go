package main

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/color"
)

func main() {
	grad := color.GenerateGradient(color.MustParseHex("#1a1f27"), color.New(100, 30.2, 39.6), 11)
	for i := range grad {
		fmt.Print(grad[i].AsHex())
	}
	fmt.Println()
}
