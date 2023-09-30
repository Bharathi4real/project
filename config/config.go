package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Configuration struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
	ServerPort string `json:"server_port"`
}

var Config Configuration

func InitConfig() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to determine caller's file path.")
	}

	configFilePath := filepath.Join(filepath.Dir(filename), "config.json")

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	err = json.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalf("Error decoding configuration JSON: %v", err)
	}
}
