package main

import (
	"fmt"
	"github.com/GranderStark/go-aes-256-cbc-cli/internal/mainActions"
	"github.com/apex/log"
	"github.com/urfave/cli"
	"os"
)

// Golang v1.8

var (
	version   = "unknown"
	gitCommit = "unknown"
	buildDate = "unknown"
)

func main() {
	cliapp := cli.NewApp()
	cliapp.Name = "aes-256-cbc-cli"
	cliapp.Usage = "Tool for aes-256-cbc-cli encryption/decryption"
	cliapp.UsageText = "aes-256-cbc-cli <command> [-t, --tables=<db>.<table>] <backup_name>"
	cliapp.Version = version

	cliapp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "cipher_key, c",
			Usage:  "cipher_key",
			EnvVar: "AES_256_CBC_CLI_CIPHER_KEY",
		},
		cli.StringFlag{
			Name:   "from_stdin, i",
			Usage:  "from_stdin",
			EnvVar: "AES_256_CBC_CLI_FROM_STDIN",
		},
		cli.BoolFlag{
			Name:   "to_stdout, s",
			Usage:  "to_stdout",
			EnvVar: "AES_256_CBC_CLI_TO_STDOUT",
		},
		cli.StringFlag{
			Name:   "from_file, f",
			Usage:  "from_file",
			EnvVar: "AES_256_CBC_CLI_FROM_FILE",
		},
		cli.StringFlag{
			Name:   "to_file, t",
			Usage:  "to_file",
			EnvVar: "AES_256_CBC_CLI_TO_FILE",
		},
	}

	cliapp.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Printf("Error. Unknown command: '%s'\n\n", command)
		cli.ShowAppHelpAndExit(c, 1)
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println("Version:\t", c.App.Version)
		fmt.Println("Git Commit:\t", gitCommit)
		fmt.Println("Build Date:\t", buildDate)
	}
	cliapp.Commands = []cli.Command{
		{
			Name:      "decrypt",
			Usage:     "Decrypt process",
			UsageText: "aes-256-cbc-cli decrypt",
			Action: func(c *cli.Context) error {
				var (
					err error
				)
				_, err = mainActions.RunDecrypt(
					c.String("c"),
					c.String("i"),
					c.Bool("s"),
					c.String("f"),
					c.String("t"))
				return err
			},
			Flags: cliapp.Flags,
		},
		{
			Name:      "encrypt",
			Usage:     "Encrypt process",
			UsageText: "aes-256-cbc-cli encrypt",
			Action: func(c *cli.Context) error {
				var (
					err error
				)
				_, err = mainActions.RunEncrypt(
					c.String("c"),
					c.String("i"),
					c.Bool("s"),
					c.String("f"),
					c.String("t"))
				return err
			},
			Flags: cliapp.Flags,
		},
	}
	if err := cliapp.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
