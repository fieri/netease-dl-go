package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/trytwice/netease-dl-go/basic"
	"github.com/trytwice/netease-dl-go/netease"
	"github.com/urfave/cli"
)

var (
	ids = []string{}
)

func main() {
	app := cli.NewApp()
	app.Name = "netease-dl-go"
	app.Usage = "A cli based netease-cloud-music downloader."
	app.Version = "1.1.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "down, d",
			Usage: "Download song by url. s for song, l for playlist.",
		},
		// cli.BoolFlag{
		// 	Name:  "version, V",
		// 	Usage: "print only the version",
		// },
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:    "down",
			Aliases: []string{"d"},
			Usage:   "download song or playlist",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "forever, forevvarr"},
				cli.BoolFlag{Name: "song, s"},
				cli.BoolFlag{Name: "list, l"},
			},
			Before: func(c *cli.Context) error {
				for i := 0; i < c.NArg(); i++ {
					ids = append(ids, basic.GetID(c.Args().Get(i)))
				}
				return nil
			},
			Action: func(c *cli.Context) error {
				if c.NArg() <= 0 {
					return errors.New("arg error: no url or id")
				}
				if c.Bool("forever") {
					// fmt.Println(c.Command.FullName())
					fmt.Println(c.Command.Names())
				}
				if c.Bool("song") {
					for _, v := range ids {
						err := netease.DownloadSongByID(v)
						if err != nil {
							return err
						}
					}
				}
				if c.Bool("list") {
					for _, v := range ids {
						err := netease.DownloadSongByPlaylist(v)
						if err != nil {
							return err
						}
					}
				}
				return nil
			},
		},
	}
	cli.AppHelpTemplate = fmt.Sprintf(`%s

EXAMPLE:
    netease-dl-go down --song 554242032     ;;Download the song which id = 554242032
    netease-dl-go down -s https://music.163.com/#/song?id=554242032    ;;Download the song which id = 554242032
    netease-dl-go d --song 554242032 531295576     ;;Download songs id = 554242032 & 531295576 etc.
    netease-dl-go d -s https://music.163.com/#/song?id=554242032 https://music.163.com/#/song?id=531295576    ;;Download songs id = 554242032 & 531295576
    netease-dl-go down --list  115707002    ;;Download the playlist which id = 115707002
    netease-dl-go down -l 115707002 6683129    ;;Download playlists id = 115707002 & 6683129
    netease-dl-go d --list https://music.163.com/#/my/m/music/playlist?id=115707002    ;;Download the playlist which id = 115707002
    netease-dl-go d -l https://music.163.com/#/my/m/music/playlist?id=115707002 https://music.163.com/#/my/m/music/playlist?id=6683129    ;;Download playlists id = 115707002 & 6683129

EMAIL: dream@trytwice.me

`, cli.AppHelpTemplate)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
