package cmd

import (
	"errors"
	"fmt"
	"github.com/LarsNorlander/colorizer/pkg/color"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "hue COLOR [flags]",
		Short: "Returns the pure hue of a color",
		RunE:  hueFn,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("hue requires a color argument")
			}
			return nil
		},

	})
}

func hueFn(command *cobra.Command, args []string) error {
	arg := args[0]

	if !strings.HasPrefix(arg, "#") {
		arg = "#" + arg
	}

	hex, err := color.ParseHex(arg)
	if err != nil {
		return err
	}

	hue := hex.RGB().Hue()
	hsl := color.HSL{
		H: hue,
		S: 1,
		L: 0.5,
	}
	fmt.Println(hsl.RGB().Hex())
	return nil
}
