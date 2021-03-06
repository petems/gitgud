package main

import (
	"fmt"
	"os"

	"github.com/petems/gitgud"
	"github.com/caarlos0/spin"
	"github.com/urfave/cli"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "gitgud"
	app.Version = version
	app.Author = "Carlos Alexandro Becker (root@carlosbecker.com)"
	app.Usage = "This is an gitgud app written in Go"
	app.Action = func(c *cli.Context) error {
		spin := spin.New("\033[36m %s Working...\033[m")
		spin.Start()
		err := gitgud.Foo()
		spin.Stop()
		if err != nil {
			return err
		}
		fmt.Println("Done!")
		return nil
	}
	_ = app.Run(os.Args)
}
