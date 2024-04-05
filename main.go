package main

import (
	"context"

	"flag"

	"log"

	"yadrotask/internal/config"
	"yadrotask/internal/service"
	"yadrotask/pkg/database"
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

	service, err := service.New(cfg.SourceURL)
	if err != nil {
		log.Fatal(err)
	}

	var comicsInfo []database.Comics
	for i := 1; i <= cfg.ComicsCount; i++ {
		comics, err := service.GetComicsDataByID(ctx, i)
		if err != nil {
			log.Fatal(err)
		}
		comicsInfo = append(comicsInfo, comics)
	}

	if outputFlag {
		if numberOfComicsFlag != 0 && numberOfComicsFlag <= len(comicsInfo) {
			comicsInfo = comicsInfo[:numberOfComicsFlag]
		}
		log.Print(comicsInfo)
	}
}
