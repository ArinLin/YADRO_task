package database

import (
	"encoding/json"
	"os"
)

type Database interface {
	AddToEntries(comicsMap []Comics)
	SaveInFile() error
}

type Comics struct {
	ID       int      `json:"-"`
	URL      string   `json:"url"`
	Keywords []string `json:"keywords"`
}

type databaseImpl struct {
	entries  map[int]Comics
	filePath string
}

func New(filePath string) Database {
	return &databaseImpl{
		entries:  make(map[int]Comics),
		filePath: filePath,
	}
}

func (d *databaseImpl) AddToEntries(comics []Comics) {
	for _, c := range comics {
		d.entries[c.ID] = c
	}
}

func (d *databaseImpl) SaveInFile() error {
	updatedData, err := json.Marshal(d.entries)
	if err != nil {
		return err
	}

	return os.WriteFile(d.filePath, updatedData, 0644)
}
