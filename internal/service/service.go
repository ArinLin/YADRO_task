package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"yadrotask/pkg/database"
	"yadrotask/pkg/words"
)

type Service interface {
	GetComicsDataByID(ctx context.Context, id int) (database.Comics, error)
}

type serviceImpl struct {
	baseUrl   string
	client    *http.Client
	stopWords map[string]struct{}
}

type Entity struct {
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
}

func New(baseUrl string) (Service, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	stopWordsMap, err := words.CreateStopWordsMap()
	if err != nil {
		return nil, err
	}

	return &serviceImpl{
		baseUrl:   baseUrl,
		client:    client,
		stopWords: stopWordsMap,
	}, nil
}

func (s *serviceImpl) GetComicsDataByID(ctx context.Context, id int) (database.Comics, error) {
	resp, err := s.client.Get(fmt.Sprintf("%s/%d/info.0.json", s.baseUrl, id))
	if err != nil {
		return database.Comics{}, err
	}

	defer resp.Body.Close()

	// экземпляр структуры, куда будут парситься данные
	var entity Entity

	// Декодирование JSON-ответа в структуру
	err = json.NewDecoder(resp.Body).Decode(&entity)
	if err != nil {
		return database.Comics{}, err
	}
	var keywords []string
	if entity.Alt != "" {
		normalizedSentence, err := words.NormalizeSentence(entity.Alt, s.stopWords)
		if err != nil {
			return database.Comics{}, err
		}
		keywords = append(keywords, normalizedSentence...)
	}

	if entity.Transcript != "" {
		normalizedSentence, err := words.NormalizeSentence(entity.Transcript, s.stopWords)
		if err != nil {
			return database.Comics{}, err
		}
		keywords = append(keywords, normalizedSentence...)
	}

	return database.Comics{
		ID:       id,
		URL:      entity.Img,
		Keywords: keywords,
	}, nil
}
