package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config map[string]string

func (c *Config) Set(key string, val string) {
	(*c)[key] = val
}

func (c *Config) get(key string) string {
	if val, exists := (*c)[key]; exists {
		return val
	}
	return ""
}

func loadConfig() (config *Config) {
	root, _ := filepath.Split(filepath.Clean(os.Args[0]))
	b, err := ioutil.ReadFile(filepath.Join(root, "config.json"))
	if err != nil {
		panic(err)
		return &Config{}
	}
	err = json.Unmarshal(b, &config)
	return
}
