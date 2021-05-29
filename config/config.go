package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	Bot struct {
		Token  string
		Prefix string
	}
	Database struct {
		Type string
	}

	config = &configStruct{}
)

type configStruct struct {
	Bot struct {
		Prefix string `yaml:"Prefix"`
		Token string `yaml:"Token"`
	} `yaml:"Bot"`
	Database struct {
		Type string `yaml:"Type"`
	} `yaml:"Database"`
}
func LoadConfig(configPath string) error {
	// Load config file
	fmt.Println("Loading Archivist Config")
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("Error loading config file " + err.Error())
		return err
	}
	// Close file silently, since we only read we can safely ignore any errors.
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Parse YAML config file
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error parsing config file " + err.Error())
		return err
	}

	// Assign loaded config to variables
	/* I really don't like how I'm doing env variable overloading here.
	However, I haven't yet found a better and more cleaner way. Yet, I'm
	fairly positive there is a better way (Using generics would work, but
	it's not quite there yet).  A big issue is they're not just one type,
	they can be several different types.
	 */
	//// Bot
	////// Token
	var envPrefix = "Archivist_Bot_"
	if value, present := os.LookupEnv(envPrefix + "Token"); present {
		Bot.Token = value
	} else {
		Bot.Token = config.Bot.Token
	}
	////// Prefix
	if value, present := os.LookupEnv(envPrefix + "Prefix"); present {
		Bot.Prefix = value
	} else {
		Bot.Prefix = config.Bot.Prefix
	}
	// Database
	envPrefix = "Archivist_Database_"
	if value, present := os.LookupEnv(envPrefix + "Type"); present {
		Database.Type = value
	} else {
		Database.Type = config.Database.Type
	}
	return nil
}