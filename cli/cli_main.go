package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func SetupCLI() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Flags: setupFlags(),
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
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
		cli.StringFlag{
			Name:  "lang",
			Value: "en",
		},
	}
	return flags
}

// func setupCommands() []*cli.Command {
// 	{
// 		commands = []*cli.Command{
// 			Name:  "topic",
// 			Usage: "Shows news by topic",
// 			Flags: setupFlags(),
// 			Action: func(c *cli.Context) error {
// 				fmt.Println("Done")
// 			},
// 		}
// 	}
// 	return commands
// }
