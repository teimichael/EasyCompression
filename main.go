package main

import (
	"EasyCompression/core"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var outputName string

	app := &cli.App{
		Name:    "EasyCompression",
		Usage:   "A CLI-based tool for (de)compression",
		Version: "v0.1.0",
		Authors: []*cli.Author{
			{
				Name: "Michael Tei",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "Sets the output name",
				Destination: &outputName,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "compress",
				Aliases: []string{"c"},
				Usage:   "Compresses the input",
				Action: func(c *cli.Context) error {
					path := c.Args().First()
					if len(path) > 0 {
						if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
							fmt.Println("File does not exist.")
						} else {
							if len(outputName) == 0 {
								outputName = "compressed"
							}
							fmt.Println("Compressing: ", path)
							core.Compress(path, outputName)
						}
					} else {
						fmt.Println("Please designate an input path.")
					}

					return nil
				},
			},
			{
				Name:    "decompress",
				Aliases: []string{"d"},
				Usage:   "Decompresses the input",
				Action: func(c *cli.Context) error {
					path := c.Args().First()
					if len(path) > 0 {
						if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
							fmt.Println("File doest not exist.")
						} else {
							if len(outputName) == 0 {
								outputName = "decompressed"
							}
							fmt.Println("Decompressing: ", path)
							core.Decompress(path, outputName)
						}
					} else {
						fmt.Println("Please designate an input path.")
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
