package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbName"`
}

type MQTTConfig struct {
	ServerAndPort string `json:"serverAndPort"`
}

type Config struct {
	Database DatabaseConfig `json:"database"`
	MQTT     MQTTConfig     `json:"mqtt"`
}

func LoadConfig(pathToJsonFile string) (config *Config, err error) {
	file, err := os.ReadFile(pathToJsonFile)
	if err != nil {
		return config, fmt.Errorf("failed to read file:%s %s", "\n",
			err.Error())
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal into Config struct:%s %s", "\n",
			err.Error())
	}

	return config, err
}
