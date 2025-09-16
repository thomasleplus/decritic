package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/urfave/cli/v2"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	app := cli.NewApp()
	app.Name = "decritic"
	app.Usage = "remove diacritics (accents) from strings"
	app.Version = "0.0.3"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "extended",
			Aliases: []string{"x"},
			Usage:   "enable extended transliterations (e.g. æ -> ae, œ -> oe, ð -> d, þ -> th, ł -> l)",
		},
	}
	app.Action = action
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	// Custom transliterations
	customReplacer := strings.NewReplacer(
		"æ", "ae", "Æ", "AE",
		"œ", "oe", "Œ", "OE",
		"ð", "d", "Ð", "D",
		"þ", "th", "Þ", "TH",
		"ł", "l", "Ł", "L",
	)

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	if c.NArg() > 0 {
		for i, in := range c.Args().Slice() {
			if i > 0 {
				fmt.Print(" ")
			}
			out, _, err := transform.String(t, in)
			if err != nil {
				fmt.Fprintln(os.Stderr, "decritic: transform error: ", err)
			}
			if c.Bool("extended") {
				out = customReplacer.Replace(out)
			}
			fmt.Print(out)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			out, _, err := transform.String(t, scanner.Text())
			if err != nil {
				fmt.Fprintln(os.Stderr, "decritic: transform error: ", err)
			}
			if c.Bool("extended") {
				out = customReplacer.Replace(out)
			}
			fmt.Println(out)
		}
		err := scanner.Err()
		if err != nil {
			fmt.Fprintln(os.Stderr, "decritic: read error: ", err)
		}
	}
	return nil
}
