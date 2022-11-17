package cmd

import (
	"easygoing/demo"
	"flag"
	"log"
)

var pname string

type Cmd struct {
}

func NewCmd() *Cmd {
	return &Cmd{}
}

func (c *Cmd) Run() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 || args[0] != "new" {
		log.Fatal("go build first, and then use command easygoing new [project name], e.g 'easygoing new easygoing-demo' ")
		return
	}
	if len(args) == 1 {
		pname = "easygoing-demo"
	} else {
		pname = args[1]
	}

	t := demo.NewTemplate(pname)
	t.Run()
}
