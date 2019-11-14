package proliferate

import (
	"encoding/json"
	"os"
)

// Config root config populated by config.json
type Config struct {
	Logging Logging       `json:"logging"`
	Couch   CouchConfig   `json:"couchDB"`
	Network NetworkConfig `json:"network"`
}

// CouchConfig couchConfig populated by config.json
type CouchConfig struct {
	Enabled  bool   `json:"enabled"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
}

// Logging logging read from config.json
type Logging struct {
	Enabled      bool   `json:"enabled"`
	Level        int    `json:"level"`
	Console      bool   `json:"console"`
	File         bool   `json:"file"`
	FileLocation string `json:"fileLocation"`
}

// NetworkConfig networkConfig read from config.json
type NetworkConfig struct {
	Algorithm string   `json:"consensusAlgorithm"`
	Role      int      `json:"role"`
	MaxPeers  int      `json:"maxPeers"`
	Discovery []string `json:"DiscoveryURL"`
}

// LoadConfig returns json as struct (TODO!)
func LoadConfig() Config {
	var config Config

	file := "config.json"

	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		LogRaw(Message{
			Level: 1,
			Text:  err.Error(),
		})
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return (config)
}