package configs

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	// "github.com/lpernett/godotenv"
	"sigs.k8s.io/yaml"
)

type SiteConfigFile struct {
	Name   string `yaml:"name"`
	Server struct {
		host string
		Port int
	} `yaml:"server"`
	YouTubeApiUrl    string `yaml:"youtubeapiurl"`
	YouTubeChannelId string `yaml:"youtubechannelid"`
}

type SiteSecrets struct {
	YouTubeAPIKey string
}

type SiteConfig struct {
	ConfigFile SiteConfigFile
	Secrets    SiteSecrets
}

var once sync.Once
var cfg *SiteConfigFile
var secrets *SiteSecrets
var config *SiteConfig

func LoadConfigFile(p string) error {
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

func LoadSecrets() error {

	// if err := godotenv.Load(".secrets.env"); err != nil {
	// 	log.Fatal("Error Loading Secrets")
	// }

	secrets = &SiteSecrets{YouTubeAPIKey: os.Getenv("YOUTUBE_API")}

	return nil
}

func GetConfig(filename string) *SiteConfig {
	once.Do(func() {
		LoadConfigFile(filename)
		LoadSecrets()
	})

	if cfg == nil || secrets == nil {
		log.Fatal("Config not loaded")
	}

	config = &SiteConfig{ConfigFile: *cfg, Secrets: *secrets}
	return config
}

var Config *SiteConfig = GetConfig("config.yaml")
