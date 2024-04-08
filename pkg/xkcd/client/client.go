package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client interface {
	GetComicsDataByID(ctx context.Context, id int) (Entity, error)
}

type clientImpl struct {
	baseUrl string
	client  *http.Client
}

type Entity struct {
	ID         int    `json:"num"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
}

func New(baseUrl string) (Client, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	return &clientImpl{
		baseUrl: baseUrl,
		client:  client,
	}, nil
}

func (s *clientImpl) GetComicsDataByID(ctx context.Context, id int) (Entity, error) {
	resp, err := s.client.Get(fmt.Sprintf("%s/%d/info.0.json", s.baseUrl, id))
	if err != nil {
		return Entity{}, err
	}
	defer resp.Body.Close()

	var entity Entity
	err = json.NewDecoder(resp.Body).Decode(&entity)
	if err != nil {
		return Entity{}, err
	}

	return entity, nil
}
