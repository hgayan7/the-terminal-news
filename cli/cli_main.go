package cli

import (
	"log"
	"os"

	"github.com/hgayan7/the-terminal-news/api"
	"github.com/urfave/cli/v2"
)

func SetupCLI() {
	app := &cli.App{
		Name:  "theterminalnews",
		Usage: "Read news in your terminal",
		Flags: setupFlags(),
		Action: func(c *cli.Context) error {
			api.FetchNews(c.String("lang"), c.StringSlice("topic"))
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
