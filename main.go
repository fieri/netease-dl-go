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
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, V",
		Usage: "print only the version",
	}
	cli.AppHelpTemplate = fmt.Sprintf(`%s

EXAMPLE:
    netease-dl-go --down s 554242032     ;;Download the song which id = 554242032
    netease-dl-go --down s https://music.163.com/#/song?id=554242032    ;;Download the song which id = 554242032
    netease-dl-go --down s 554242032 531295576     ;;Download songs id = 554242032 & 531295576 etc.
    netease-dl-go -d s https://music.163.com/#/song?id=554242032 https://music.163.com/#/song?id=531295576    ;;Download songs id = 554242032 & 531295576
    netease-dl-go --down l  115707002    ;;Download the playlist which id = 115707002
    netease-dl-go -d l 115707002 6683129    ;;Download playlists id = 115707002 & 6683129
    netease-dl-go --down l https://music.163.com/#/my/m/music/playlist?id=115707002    ;;Download the playlist which id = 115707002
    netease-dl-go -d l https://music.163.com/#/my/m/music/playlist?id=115707002 https://music.163.com/#/my/m/music/playlist?id=6683129    ;;Download playlists id = 115707002 & 6683129

EMAIL: dream@trytwice.me

`, cli.AppHelpTemplate)

	app.Action = func(c *cli.Context) error {
		if downVar := c.String("down"); downVar != "s" && downVar != "l" {
			return errors.New("flag error: do you use the command correctly")
		}
		if c.NArg() <= 0 {
			return errors.New("arg error: no url or id")
		}
		ids := []string{}
		for i := 0; i < c.NArg(); i++ {
			ids = append(ids, basic.GetID(c.Args().Get(i)))
		}
		if c.String("down") == "s" {
			for _, v := range ids {
				err := netease.DownloadSongByID(v)
				if err != nil {
					return err
				}
			}
		}
		if c.String("down") == "l" {
			for _, v := range ids {
				err := netease.DownloadSongByPlaylist(v)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
