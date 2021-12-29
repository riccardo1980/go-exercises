package choose

import (
	base "choose/base"
	story "choose/story"
	"flag"
	"fmt"
	"log"
)

type cmdline struct {
	base.CommonFlags
}

func (c *cmdline) Read() story.Story {
	st, _ := story.FromFile(c.InputFile)
	return st
}

func New() base.Base {
	var c cmdline
	c.FlagSet = flag.NewFlagSet("cmdline", flag.ExitOnError)
	base.CommonFlagsDefine(&c.CommonFlags)
	return &c
}

func (c *cmdline) Parse(args []string) {
	log.Println("cmdline parsing")
	c.FlagSet.Parse(args)
}

func (c *cmdline) Run(st story.Story) {
	log.Println("Start!")
	arc := st["intro"]
	for len(arc.Options) > 0 {
		choice := make(chan int, 1)

		go func() {
			ArcShow(arc)
			var c int
			fmt.Scanf("%d", &c)
			choice <- c
		}()

		c := <-choice
		if c < 0 || c >= len(arc.Options) {
			continue
		}
		arc = st[arc.Options[c].Arc]
	}
	log.Println("End!")
}

func ArcShow(arc story.Arc) {
	fmt.Println(arc.Title)

	for _, s := range arc.Story {
		fmt.Println(s)
		fmt.Println()
	}

	for idx, opt := range arc.Options {
		fmt.Printf("[ %d ] - %s\n", idx, opt.Text)
	}
}
