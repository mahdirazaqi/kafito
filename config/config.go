package config

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		Port      string    `json:"port"`
		SecretKey string    `json:"secret_key"`
		Mongo     Mongo     `json:"mongo"`
		Kavenegar Kavenegar `json:"kavenegar"`
	}

	Mongo struct {
		Host     string `json:"host"`
		DB       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	}

	Kavenegar struct {
		Token    string `json:"token"`
		Template string `json:"template"`
	}
)

var C = new(Config)

func Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, C)
}
