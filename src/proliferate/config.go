package proliferate

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	UseCouchDB bool    `json:"useCouchDB"`
	Logging    Logging `json:"logging"`
}

type Logging struct {
	Enabled      bool   `json:"enabled"`
	Level        int    `json:"level"`
	Console      bool   `json:"console"`
	File         bool   `json:"file"`
	FileLocation string `json:"fileLocation"`
}

// LoadJSON returns json as struct (TODO!)
func LoadConfig() Config {
	var config Config

	file := "config.json"

	configFile, err := os.Open(file)
	defer configFile.Close()

	// TODO enable logging
	if err != nil {
		//Log("util.LoadJSON", err.Error(), 4)
		fmt.Println(err)
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	// Build struct to match
	return (config)
}
