package service

import (
	"context"
	"yadrotask/pkg/database"
	"yadrotask/pkg/words"
	"yadrotask/pkg/xkcd/client"
)

type Service interface {
	GetComicsData(ctx context.Context, n int) ([]database.Comics, error)
}

type serviceImpl struct {
	stopWords map[string]struct{}
	client    client.Client
	db        database.Database
}

func New(client client.Client, db database.Database, fileName string) (Service, error) {
	stopWordsMap, err := words.CreateStopWordsMap(fileName)
	if err != nil {
		return nil, err
	}

	return &serviceImpl{
		stopWords: stopWordsMap,
		client:    client,
		db:        db,
	}, nil
}

func (s *serviceImpl) GetComicsData(ctx context.Context, n int) ([]database.Comics, error) {
	var comics []database.Comics
	for i := 1; i <= n; i++ {
		comicsData, err := s.client.GetComicsDataByID(ctx, i)
		if err != nil {
			return nil, err
		}

		entity, err := dto(comicsData, s.stopWords)
		if err != nil {
			return nil, err
		}

		comics = append(comics, entity)
	}

	s.db.AddToEntries(comics)
	s.db.SaveInFile()

	return comics, nil
}

func dto(entity client.Entity, stopWords map[string]struct{}) (database.Comics, error) {
	var keywords []string
	if entity.Alt != "" {
		normalizedSentence, err := words.NormalizeSentence(entity.Alt, stopWords)
		if err != nil {
			return database.Comics{}, err
		}
		keywords = append(keywords, normalizedSentence...)
	}

	if entity.Transcript != "" {
		normalizedSentence, err := words.NormalizeSentence(entity.Transcript, stopWords)
		if err != nil {
			return database.Comics{}, err
		}
		keywords = append(keywords, normalizedSentence...)
	}

	return database.Comics{
		ID:       entity.ID,
		URL:      entity.Img,
		Keywords: keywords,
	}, nil
}
