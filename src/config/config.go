package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)


type config struct {
	HTTPPort  uint16 `json:"http_port"`
	StaticFilePath string `json:"static_file_path"`
}

func LoadConfig(path string) *config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("LoadConfig:Error: Failed reading configuration from %s\nError: %s\n", path, err.Error())
	}

	s := &config{}
	err = json.Unmarshal(data, s)
	if err != nil {
		log.Fatalf("LoadConfig:Error: Failed unmarshalling json from config file Error: %s\nRaw data: %v\n", err.Error(),
			string(data))
	}
	return s
}
