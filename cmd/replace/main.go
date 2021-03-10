package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// A program that replaces hex values for Jetbrains Themes
func main() {
	// Load in theme file
	theme, err := ioutil.ReadFile("data/DesignerCode.icls")
	panicOnErr(err)

	// Load in original file
	midnight, err := ioutil.ReadFile("data/midnight.txt")
	panicOnErr(err)

	var original []string
	for i := 0; i < len(midnight)-6; i += 6 {
		original = append(original, string(midnight[i:i+6]))
	}

	// Load in replacements
	charcoal, err := ioutil.ReadFile("data/charcoal.txt")
	panicOnErr(err)

	var replacement []string
	for i := 0; i < len(charcoal)-6; i += 6 {
		replacement = append(replacement, string(charcoal[i:i+6]))
	}

	for i, oc := range original {
		theme = bytes.ReplaceAll(theme, []byte(oc), []byte(replacement[i]))
	}

	// replace all 5 long hex
	re5 := regexp.MustCompile("value=\"([0-9|a-f]{5})\"")
	res := re5.FindAllStringSubmatch(string(theme), -1)
	hex5Replacement := make(map[string]string)
	for _, item := range res {
		hex5Replacement[item[1]] = ""
	}
	for oc, _ := range hex5Replacement {
		for i, oc6 := range original {
			if strings.Contains(oc6, oc) {
				hex5Replacement[oc] = replacement[i]
				break
			}
		}
	}
	for oc, nc := range hex5Replacement {
		theme = bytes.ReplaceAll(theme, []byte(oc), []byte(nc))
	}
	// Get all matches
	// Loop through each determine which in the original does it match, and then record what it should be replaced with

	// Replace all 4 long hex
	re4 := regexp.MustCompile("value=\"([0-9|a-f]{4})\"")
	res2 := re4.FindAllStringSubmatch(string(theme), -1)
	hex4Replacement := make(map[string]string)
	for _, item := range res2 {
		hex4Replacement[item[1]] = ""
	}
	for oc, _ := range hex4Replacement {
		for i, oc6 := range original {
			if strings.Contains(oc6, oc) {
				hex4Replacement[oc] = replacement[i]
				break
			}
		}
	}
	for oc, nc := range hex4Replacement {
		theme = bytes.ReplaceAll(theme, []byte(oc), []byte(nc))
	}

	fmt.Println(string(theme))
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
