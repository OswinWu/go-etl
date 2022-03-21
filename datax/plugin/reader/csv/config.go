package csv

import (
	"encoding/json"

	"github.com/Breeze0806/go-etl/config"
	//csv storage
	"github.com/Breeze0806/go-etl/storage/stream/file/csv"
)

type Config struct {
	csv.Config

	Path []string `json:"path"`
}

func NewConfig(conf *config.JSON) (c *Config, err error) {
	c = &Config{}
	if err = json.Unmarshal([]byte(conf.String()), c); err != nil {
		return
	}
	return
}
