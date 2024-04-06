package main

import (
	"context"
	"flag"
	"log"

	"yadrotask/internal/config"
	"yadrotask/internal/service"
	"yadrotask/pkg/database"
	"yadrotask/pkg/xkcd/client"
)

var (
	outputFlag         bool
	numberOfComicsFlag int
)

func main() {
	flag.BoolVar(&outputFlag, "o", false, "activate console output")
	flag.IntVar(&numberOfComicsFlag, "n", 0, "number of output comics")
	flag.Parse()

	ctx := context.Background()

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := database.New(cfg.DBFile)

	client, err := client.New(cfg.SourceURL)
	if err != nil {
		log.Fatal(err)
	}

	service, err := service.New(client, db, cfg.StopWords)
	if err != nil {
		log.Fatal(err)
	}

	comicsInfo, err := service.GetComicsData(ctx, cfg.ComicsCount)
	if err != nil {
		log.Fatal(err)
	}

	if outputFlag {
		if numberOfComicsFlag != 0 && numberOfComicsFlag <= len(comicsInfo) {
			comicsInfo = comicsInfo[:numberOfComicsFlag]
		}
		log.Print(comicsInfo)
	}
}
