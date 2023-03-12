package main

import (
	"flag"
	"goqueue/configs"
)

func main() {
	dbgMode := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	_ = configs.New(*dbgMode)
}
