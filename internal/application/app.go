package application

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/lvrach/testprof/internal/pkg"
	"github.com/lvrach/testprof/internal/test2json"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name:  "testprof",
		Usage: "testprof: performance tools for go tests",
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:   "timings",
				Usage:  "compute pkg timings for tests",
				Action: Timings,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "test2json",
						Value: "test2json.json",
					},
					&cli.StringFlag{
						Name:  "timings",
						Value: "testdata/testprof/timings.json",
					},
				},
			},
			{
				Name:   "list",
				Usage:  "",
				Action: List,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "timings",
						Value: "testdata/testprof/timings.json",
					},
					&cli.StringFlag{
						Name:    "portion",
						Value:   "1/1",
						EnvVars: []string{"TEST_PORTION"},
					},
				},
			},
		},

		EnableBashCompletion: true,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}

func Timings(c *cli.Context) error {
	f, err := os.Open(c.String("test2json"))
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := test2json.Parse(f)
	if err != nil {
		return err
	}

	ts := pkg.Timings(r)

	dir, _ := path.Split(c.String("timings"))
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	out, err := os.Create(c.String("timings"))
	if err != nil {
		return err
	}

	outJSON := json.NewEncoder(out)
	outJSON.SetIndent("", "  ")

	err = outJSON.Encode(ts)
	return err
}

func List(c *cli.Context) error {
	f, err := os.Open(c.String("timings"))
	if err != nil {
		return err
	}

	ts, err := pkg.Parse(f)
	if err != nil {
		return err
	}

	var i, n int
	fmt.Sscanf(c.String("portion"), "%d/%d", &i, &n)

	s := pkg.Splitter{
		Timings: ts,
		Parts:   n,
	}

	fmt.Fprintln(os.Stdout, strings.Join(s.List(i-1), "\n"))

	return nil
}
