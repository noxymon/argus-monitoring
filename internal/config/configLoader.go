package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type ArgusConfig struct {
	ConfigList []Config `json:"configList"`
}

type Config struct {
	Type           string `json:"type"`
	HealthCheckURL string `json:"healthCheckUrl"`
	ProcessName    string `json:"processName"`
	CommandLine    string `json:"commandLine"`
}

type ConfigLoader struct {
	ConfigPath string
}

func (c ConfigLoader) Load() (*ArgusConfig, error) {
	file, err := ioutil.ReadFile(c.ConfigPath)

	if(err!=nil){
		return nil, errors.New("error on opening json")
	}

	var loadedConfig ArgusConfig
	err = json.Unmarshal(file, &loadedConfig)

	if(err!=nil){
		return nil, errors.New("Failed on deserializing json.")
	}

	return &loadedConfig, nil
}