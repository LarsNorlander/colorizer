package theme_test

import (
	"fmt"
	"github.com/LarsNorlander/colorizer/pkg/theme"
	"testing"
)

func Test(t *testing.T) {
	// TODO Actually test this logic

	thm, err := theme.ParseFile("../../themes/charcoal.yaml")
	if err != nil {
		t.Fatalf("%v", err)
	}

	fmt.Println(thm)
}
