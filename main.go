package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "wctool",
		Usage: "A command line tool challenge",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "c",
				Usage:   "Return number of bytes in text",
				Aliases: []string{"C"},
			},
			&cli.BoolFlag{
				Name:    "w",
				Usage:   "Return number of words in text",
				Aliases: []string{"W"},
			},
		},
		Action: func(c *cli.Context) error {
			switch {
			case c.Bool("c"):
				runC()
			case c.Bool("w"):
				runW()
			default:
				fmt.Println("Invalid flag")
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runC() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	byteCount := len(file)
	fmt.Println(byteCount)
}

func runW() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(file)
	words := strings.Fields(text)
	wordCount := len(words)

	fmt.Println(wordCount)
}
