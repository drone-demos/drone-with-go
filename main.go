package main

import (
	. "github.com/josebarn/drone-with-go/api"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	var redis_url string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "redis, r",
			Usage:       "redis server uri",
			Destination: &redis_url,
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(redis_url) <= 0 {
			log.Println("not using redis server")
		}
		ApiInit()
		return nil
	}

	app.HideVersion = true

	app.Run(os.Args)
}
