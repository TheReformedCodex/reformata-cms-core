package configs

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"sigs.k8s.io/yaml"
)

type SiteConfig struct {
	Name   string `yaml:"name"`
	Server struct {
		host string
		Port int
	} `yaml:"server"`
}

var once sync.Once
var cfg *SiteConfig

func LoadConfig(p string) error {
	// var in_scope_cfg *SiteConfig
	current_dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to locate local directory", err)
		return err
	}

	config_path := filepath.Join(current_dir, "configs", p)

	config_file, err := os.ReadFile(config_path)

	if err != nil {
		log.Fatal("Encountered Error Loading Config", err)
		return err
	}

	if err := yaml.Unmarshal(config_file, &cfg); err != nil {
		log.Fatal("Encountered Error Loading Config", err)
		return err
	}

	log.Println("Config file loaded")
	return nil
}


func Get(filename string) *SiteConfig {
	once.Do(func() { LoadConfig(filename) })
	if cfg == nil {
		log.Fatal("Config not loaded")
	}
	return cfg
}
