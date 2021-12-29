package main

import (
	base "choose/base"
	cmdline "choose/cmdline"
	web "choose/web"
	"fmt"
	"log"
	"os"
)

func New(s string) base.Base {
	mapping := map[string](func() base.Base){
		"cmdline": cmdline.New,
		"web":     web.New,
	}

	if builder, ok := mapping[s]; ok {
		return builder()
	}
	log.Fatal(fmt.Sprintf("Can't find %s command", s))
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	b := New(os.Args[1])
	b.Parse(os.Args[2:])
	story := b.Read()

	log.Println("Start")

	b.Run(story)

}
