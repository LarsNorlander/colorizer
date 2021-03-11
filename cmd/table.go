package cmd

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/pkg/color"
	"github.com/LarsNorlander/colorizer/pkg/theme"
	"github.com/spf13/cobra"
)

var themePath string

func init() {
	command := cobra.Command{
		Use:   "table",
		Short: "Generate a table",
		RunE:  tableFn,
	}
	command.Flags().StringVarP(&themePath, "theme", "t", "", "Path to theme file")

	rootCmd.AddCommand(&command)
}

func tableFn(cmd *cobra.Command, args []string) error {
	// Setup
	var cw *color.ColorWheel
	var blk, wht color.Color
	var err error

	if themePath == "" {
		cw, err = color.GenerateColorWheel(map[color.Name]color.Color{
			color.Red:   color.MustParseHex("#f00"),
			color.Green: color.MustParseHex("#0f0"),
			color.Blue:  color.MustParseHex("#00f"),
		})
		if err != nil {
			return err
		}
		blk = color.MustParseHex("#000000")
		wht = color.MustParseHex("#ffffff")
	} else {
		thm, err := theme.ParseFile(themePath)
		if err != nil {
			return err
		}

		cw, err = color.GenerateColorWheel(thm.Colors)
		if err != nil {
			return err
		}
		blk = thm.Black
		wht = thm.White
	}

	// Calculate grays
	gray := color.HSLBlend(blk, wht, color.HueDistanceCW)
	grays := color.HSLGradient(6, color.HueDistanceCW, blk, gray, wht)
	for _, color := range grays {
		fmt.Print(color.RGB().Hex())
	}
	fmt.Println()

	// Calculate colors with shades and tints
	for i := 0; i < 12; i++ {
		line := color.RGBGradient(6,
			blk,
			cw.Get(i),
			wht,
		)
		for _, clr := range line {
			fmt.Print(clr.RGB().Hex())
		}
		fmt.Println()
	}

	// Calculate tones for colors
	fmt.Println()
	fmt.Println("Tones:") // 12 steps
	fmt.Println()

	for i := 0; i < 12; i++ {
		line := color.RGBGradient(12,
			gray,
			cw.Get(i),
		)
		for _, clr := range line {
			fmt.Print(clr.RGB().Hex())
		}
		fmt.Println()
	}

	return nil
}
