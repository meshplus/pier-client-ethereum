package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gobuffalo/packr/v2"

	"github.com/fatih/color"
	"github.com/gobuffalo/packd"
	"github.com/hashicorp/go-plugin"
	"github.com/meshplus/pier/pkg/plugins"
	"github.com/urfave/cli"
)

var initCMD = cli.Command{
	Name:  "init",
	Usage: "Get appchain default configuration",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "target",
			Usage:    "Sepcify where to put the default configuration files",
			Required: false,
		},
	},
	Action: func(ctx *cli.Context) error {
		target := ctx.String("target")
		box := packr.NewBox("config")

		if err := box.Walk(func(s string, file packd.File) error {
			p := filepath.Join(target, s)
			dir := filepath.Dir(p)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				err := os.MkdirAll(dir, 0755)
				if err != nil {
					return err
				}
			}
			return ioutil.WriteFile(p, []byte(file.String()), 0644)
		}); err != nil {
			return err
		}

		return nil
	},
}

var startCMD = cli.Command{
	Name:  "start",
	Usage: "Start ethereum appchain plugin",
	Action: func(ctx *cli.Context) error {
		plugin.Serve(&plugin.ServeConfig{
			HandshakeConfig: plugins.Handshake,
			Plugins: map[string]plugin.Plugin{
				plugins.PluginName: &plugins.AppchainGRPCPlugin{Impl: &Client{}},
			},
			Logger:     logger,
			GRPCServer: plugin.DefaultGRPCServer,
		})

		logger.Info("Plugin server down")

		return nil
	},
}

var (
	// CurrentCommit current git commit hash
	CurrentCommit = ""
	// CurrentBranch current git branch
	CurrentBranch = ""
	// CurrentVersion current project version
	CurrentVersion = "0.0.0"
	// BuildDate compile date
	BuildDate = ""
	// GoVersion system go version
	GoVersion = runtime.Version()
	// Platform info
	Platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

var versionCMD = cli.Command{
	Name:  "version",
	Usage: "eth-client version",
	Action: func(ctx *cli.Context) error {
		printVersion()
		return nil
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "ethereum-plugin"
	app.Usage = "Manipulate the ethereum blockchain"
	app.Compiled = time.Now()
	cli.VersionPrinter = func(c *cli.Context) {
		printVersion()
	}

	app.Commands = []cli.Command{
		initCMD,
		startCMD,
		versionCMD,
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Red(err.Error())
		os.Exit(-1)
	}
}

func printVersion() {
	fmt.Printf("eth-client version: %s-%s-%s\n", CurrentVersion, CurrentBranch, CurrentCommit)
	fmt.Printf("App build date: %s\n", BuildDate)
	fmt.Printf("System version: %s\n", Platform)
	fmt.Printf("Golang version: %s\n", GoVersion)
	fmt.Println()
}
