package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"os"
	"unicode"
)

func main() {
	app := cli.NewApp()
	app.Name = "decritic"
	app.Usage = "remove diacritics (accents) from strings"
	app.Version = "0.0.2"
	app.Action = action
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	if c.NArg() > 0 {
		for i, in := range c.Args().Slice() {
			if i > 0 {
				fmt.Print(" ")
			}
			out, _, _ := transform.String(t, in)
			fmt.Print(out)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			out, _, err := transform.String(t, scanner.Text())
			if err != nil {
				fmt.Fprintln(os.Stderr, "decritic: transform error: ", err)
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
