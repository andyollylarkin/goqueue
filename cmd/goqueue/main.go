package main

import (
	"flag"
	"fmt"
	application "goqueue/internal/app"
	"goqueue/internal/configs"
	"log"
	"time"
)

var (
	Version   = "dev"
	BuildTime = time.Now().Format("2006-01-02 15:04:05")
)

func main() {
	fmt.Printf("Current build version: %s\nBuild time: %s\n", Version, BuildTime)
	dbgMode := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	config := configs.MustNewConfig(*dbgMode)
	app := application.NewApplication(config)
	err := app.Start()
	if err != nil {
		log.Fatalf("Application error: %s", err.Error())
	}
}
