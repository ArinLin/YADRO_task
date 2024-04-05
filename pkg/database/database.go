package database

type Comics struct {
	ID       int      `json:"-"`
	URL      string   `json:"url"`
	Keywords []string `json:"keywords`
}

type Database struct {
	Entries  map[int]Comics
	FilePath string
}
