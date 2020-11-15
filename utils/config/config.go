package config

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Prop map[string]string

type config struct {
	path string
}

//InitConfig loads config file with specified path relative to project runtime directory.
func InitConfig(path string) *config {
	return &config{path: path}
}

//Load is used to load app config to map.
func (cfg *config) Load() Prop {
	execPath, _ := os.Getwd()
	// build config file path.
	configPath := execPath + string(os.PathSeparator) + cfg.path
	// read config file.
	cfgBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("failed to load config file with path " + configPath)
	}
	cfgContent := string(cfgBytes)
	cfgEntries := strings.Split(cfgContent, "\n")
	kvMap := make(map[string]string, len(cfgEntries))
	for _, entry := range cfgEntries {
		kv := strings.Split(entry, "=")
		if len(kv) != 2 {
			log.Printf("one config entry parse faield, %s\n", entry)
			continue
		}
		kvMap[kv[0]] = kv[1]
	}
	return kvMap
}
