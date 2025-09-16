package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestDecritic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty", "", ""},
		{"ASCII", "hello world", "hello world"},
		{"French", "déjà vu", "deja vu"},
		{"Spanish", "pingüino", "pinguino"},
		{"German", "über", "uber"},
		{"Vietnamese", "phở", "pho"},
		{"Mixed", "crème brûlée", "creme brulee"},
		{"More French", "français", "francais"},
		{"More Spanish", "mañana", "manana"},
		{"More German", "schön", "schon"},
		{"Accents on capitals", "ÉÀÇ", "EAC"},
		{"Czech", "žluťoučký kůň", "zlutoucky kun"},
		{"Polish", "zażółć gęślą jaźń", "zazołc gesla jazn"},
		{"Romanian", "învățământ", "invatamant"},
		{"Slovak", "ľúbozvučný", "lubozvucny"},
		{"Hungarian", "árvíztűrő tükörfúrógép", "arvizturo tukorfurogep"},
		{"Latvian", "glāžšķūņu rūķīši", "glazskunu rukisi"},
		{"Lithuanian", "įgarsinkite", "igarsinkite"},
		{"Turkish", "şeker", "seker"},
		{"Icelandic", "hæð", "hæð"},
		{"No accents", "12345", "12345"},
		{"Symbols", "!@#$%^&*()", "!@#$%^&*()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate command-line arguments
			app := &cli.App{
				Action: action,
			}

			// Redirect stdout to a buffer to capture the output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Run the action with the test input
			args := os.Args[0:1]
			args = append(args, strings.Split(tt.input, " ")...)
			app.Run(args)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			io.Copy(&buf, r)
			got := strings.TrimSpace(buf.String())

			if got != tt.expected {
				t.Errorf("decritic(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}