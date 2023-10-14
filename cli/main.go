package cli

import (
	"log"
	"mci/aws/ec2"

	"os"

	"github.com/urfave/cli/v2"
)

func Init() {

	app := &cli.App{
		Name:        "mci",
		Usage:       "To be able to manipulate deployment in multiple cloud environment like AWS, AZURE, GOOGLE CLOUD",
		Description: "Multi-Cloud Interface",
	}

	c := []*cli.Command{}
	c = append(c, ec2.Command())
	app.Commands = c
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
