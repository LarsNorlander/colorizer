package color

import (
	"fmt"
	"testing"
)

func TestGenerateColorWheel(t *testing.T) {
	// TODO Actually test this logic
	fmt.Print(GenerateColorWheel(map[Name]Color{
		Green: MustParseHex("#00ff00"),
		Blue:  MustParseHex("#0000ff"),
	}))
}
