package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slog"
)

func main() {
	app := cli.App{
		Name: "cmdbootstrap",
		Version: fmt.Sprintf("version: %s commit: %s date: %s builtby: %s",
			version, commit, date, builtBy),
		Action: action,
	}
	err := app.Run(os.Args)
	if err != nil {
		slog.Error("main action failed", slog.String("err", err.Error()))
	}
}
