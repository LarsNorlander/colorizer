package theme

import (
	"github.com/LarsNorlander/colorizer/pkg/color"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	mapper = map[string]color.Name{
		"red":         0,
		"orange":      1,
		"yellow":      2,
		"yellowGreen": 3,
		"green":       4,
		"greenCyan":   5,
		"cyan":        6,
		"cyanBlue":    7,
		"blue":        8,
		"blueMagenta": 9,
		"magenta":     10,
		"magentaRed":  11,
	}
)

type theme struct {
	Metadata struct {
		Name   string `yaml:"name"`
		Author string `yaml:"author"`
	} `yaml:"metadata"`
	Black  string            `yaml:"black"`
	White  string            `yaml:"white"`
	Colors map[string]string `yaml:"colors"`
}

func ParseFile(path string) (Theme, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return Theme{}, err
	}

	temp := theme{}
	err = yaml.Unmarshal(file, &temp)
	if err != nil {
		return Theme{}, err
	}

	blk, err := color.ParseHex(temp.Black)
	if err != nil {
		return Theme{}, err
	}

	wht, err := color.ParseHex(temp.White)
	if err != nil {
		return Theme{}, err
	}

	clrs := map[color.Name]color.Color{}
	for k, v := range temp.Colors {
		name, ok := mapper[k]
		if !ok {
			continue
		}

		clr, err := color.ParseHex(v)
		if err != nil {
			return Theme{}, err
		}

		clrs[name] = clr
	}

	return Theme{
		Metadata: Metadata{
			Name:   temp.Metadata.Name,
			Author: temp.Metadata.Author,
		},
		Black:  blk,
		White:  wht,
		Colors: clrs,
	}, nil
}
