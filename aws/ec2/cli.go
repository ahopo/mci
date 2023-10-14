package ec2

import (
	"github.com/urfave/cli/v2"
)

func Command() (mainCommand *cli.Command) {

	sc := []*cli.Command{
		{
			Name:        "create",
			Aliases:     []string{"c"},
			Description: "To Create AWS EC2 Instance",
			Action:      CreateEc2Instance,
		},
		{
			Name:        "update",
			Aliases:     []string{"u"},
			Description: "To Update AWS EC2 Instance",
			Action:      CreateEc2Instance,
		},
		{
			Name:        "delete",
			Aliases:     []string{"d"},
			Description: "To Delete AWS EC2 Instance",
			Action:      CreateEc2Instance,
		},
		{
			Name:        "duplicate",
			Aliases:     []string{"dp"},
			Description: "To duplicate AWS EC2 Instance",
			Action:      CreateEc2Instance,
		},
	}

	ec2Command := []*cli.Command{
		{Name: "ec2",
			// Aliases:     []string{"cei"},
			Description: "AWS Elastic Compute Cloud",
			Subcommands: sc},
	}
	mainCommand = &cli.Command{
		Name:        "aws",
		Description: "AWS Cloud",
		Subcommands: ec2Command,
	}
	return
}
