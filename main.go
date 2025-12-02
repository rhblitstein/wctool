package main

import (
	"fmt"
	"log"
	"os"

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
		},
		Action: func(c *cli.Context) error {
			if c.Bool("c") {
				runC()
			} else {
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
	fmt.Print("test")
}
