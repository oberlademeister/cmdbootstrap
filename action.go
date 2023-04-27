package main

import "github.com/urfave/cli/v2"

func action(c *cli.Context) error {
	debug := c.Bool("debug")
	setup_log(debug)
	return nil
}
