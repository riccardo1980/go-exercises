package base

import (
	"choose/story"
	"flag"
	"log"
)

type CommonFlags struct {
	InputFile string
	FlagSet   *flag.FlagSet
}

type Base interface {
	Read() story.Story
	Parse([]string)
	Run(story.Story)
}

func CommonFlagsDefine(c *CommonFlags) {
	log.Println("Common flags definition")
	c.FlagSet.StringVar(&c.InputFile, "inputFile", "gopher.json", "Input json file")
}
