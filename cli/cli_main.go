package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func SetupCLI() {
	app := &cli.App{
		Name:     "theterminalnews",
		Usage:    "Read news in your terminal",
		Flags:    setupFlags(),
		Commands: setupCommands(),
		Action: func(c *cli.Context) error {
			// api.FetchNews()
			fmt.Println(c.StringSlice("topic"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func setupFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "lang",
			Value:   "en",
			Aliases: []string{"l"},
		},
		&cli.StringSliceFlag{
			Name:    "topic",
			Aliases: []string{"t"},
		},
	}
	return flags
}

func setupCommands() []*cli.Command {
	commands := []*cli.Command{
		{
			Name:    "tech",
			Aliases: []string{"te"},
			Usage:   "Get news about technology",
			Action: func(c *cli.Context) error {
				fmt.Println("tech command used")
				return nil
			},
		},
	}
	return commands
}
