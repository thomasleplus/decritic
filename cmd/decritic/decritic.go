/*
 * Decritic - Removes diacritics
 * Copyright (C) 2016 Thomas Leplus
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"os"
	"unicode"
)

func main() {
	app := cli.NewApp()
	app.Name = "decritic"
	app.Usage = "remove diacritics (accents) from strings"
	app.Version = "0.0.1"
	app.Action = action
	if app.Run(os.Args) != nil {
		os.Exit(1)
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
