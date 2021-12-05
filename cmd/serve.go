package cmd

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/pkg/color"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "Start the colorizer server",
		RunE:  serveFn,
	})
}

func serveFn(command *cobra.Command, args []string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		cw, err := color.GenerateColorWheel(map[color.Name]color.Color{
			color.Red:   color.MustParseHex("#FF0000").RGB(),
			color.Green: color.MustParseHex("#00ff00").RGB(),
			color.Blue:  color.MustParseHex("#0000ff").RGB(),
		})
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}

		blk := color.MustParseHex("#1A1F27").RGB()
		wht := color.MustParseHex("#f7f8fa").RGB()

		original := color.MustParseHex("#503030").RGB()

		mapped := color.NewMapper(cw, blk, wht)(original)

		_, _ = fmt.Fprintf(writer, "original: %s\nmapped  : %s\n", original, mapped)
	})
	log.Print("Serving colorizer...")
	return http.ListenAndServe(":8000", nil)
}
