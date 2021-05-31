package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strconv"
)

var (
	Bot struct {
		Token  string
		Prefix string
	}
	Database struct {
		Type string
		Host string
		User string
		Password string
		Protocol string
		Port int
		DatabaseName string
		SSL bool
		Path string
	}
	Log struct {
		Level string
		Path string
		MaxSize int
		MaxAge int
		MaxBackups int
		Compress bool
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
		Host string `yaml:"host"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Protocol string `yaml:"protocol"`
		Port int `yaml:"port"`
		DatabaseName string `yaml:"DatabaseName"`
		SSL bool `yaml:"ssl"`
		Path string `yaml:"path"`
	} `yaml:"Database"`
	Log struct {
		Level string `yaml:"Level"`
		Path string `yaml:"Path"`
		MaxSize int `yaml:"max_size"`
		MaxAge int `yaml:"max_age"`
		MaxBackups int `yaml:"max_backups"`
		Compress bool `yaml:"compress"`
	} `yaml:"Log"`
}
func LoadConfig(configPath string) error {
	// Load config file
	log.Println("Loading Archivist Config from " + configPath)
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalln("Error loading config file " + err.Error())
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
		log.Fatalln("Error parsing config file " + err.Error())
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
	if value, present := os.LookupEnv(envPrefix + "Host"); present {
		Database.Host = value
	} else {
		Database.Host = config.Database.Host
	}
	if value, present := os.LookupEnv(envPrefix + "User"); present {
		Database.User = value
	} else {
		Database.User = config.Database.User
	}
	if value, present := os.LookupEnv(envPrefix + "Password"); present {
		Database.Password = value
	} else {
		Database.Password = config.Database.Password
	}
	if value, present := os.LookupEnv(envPrefix + "Protocol"); present {
		Database.Protocol = value
	} else {
		Database.Protocol = config.Database.Protocol
	}
	//// Validate protocols
	if Database.Protocol != "tcp" && Database.Protocol != "udp" {
		log.Fatalln("Invalid value for Database:Protocol]")
	}
	if value, present := os.LookupEnv(envPrefix + "Port"); present {
		parsedVal, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			log.Fatalln("Invalid value for Database:Port\n" + err.Error())
		}
		Database.Port = int(parsedVal)
	} else {
		Database.Port = config.Database.Port
	}
	if value, present := os.LookupEnv(envPrefix + "DatabaseName"); present {
		Database.DatabaseName = value
	} else {
		Database.DatabaseName = config.Database.DatabaseName
	}
	if value, present := os.LookupEnv(envPrefix + "SSL"); present {
		parsedVal, err := strconv.ParseBool(value)
		if err != nil {
			log.Fatalln("Invalid value for Database:SSL\n" + err.Error())
		}
		Database.SSL = parsedVal
	} else {
		Database.SSL = config.Database.SSL
	}
	if value, present := os.LookupEnv(envPrefix + "Path"); present {
		Database.Path = value
	} else {
		Database.Path = config.Database.Path
	}
	// Log
	envPrefix = "Archivist_Log_"
	if value, present := os.LookupEnv(envPrefix + "Level"); present {
		Log.Level = value
	} else {
		Log.Level = config.Log.Level
	}
	if value, present := os.LookupEnv(envPrefix + "Path"); present {
		Log.Path = value
	} else {
		Log.Path = config.Log.Path
	}
	if value, present := os.LookupEnv(envPrefix + "MaxSize"); present {
		parsedVal, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			log.Fatalln("Invalid value for Log:MaxSize\n" + err.Error())
		}
		Log.MaxSize = int(parsedVal)
	} else {
		Log.MaxSize = config.Log.MaxSize
	}
	if value, present := os.LookupEnv(envPrefix + "MaxAge"); present {
		parsedVal, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			log.Fatalln("Invalid value for Log:MaxAge\n" + err.Error())
		}
		Log.MaxAge = int(parsedVal)
	} else {
		Log.MaxAge = config.Log.MaxAge
	}
	if value, present := os.LookupEnv(envPrefix + "MaxBackups"); present {
		parsedVal, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			log.Fatalln("Invalid value for Log:MaxBackups\n" + err.Error())
		}
		Log.MaxBackups = int(parsedVal)
	} else {
		Log.MaxBackups = config.Log.MaxBackups
	}
	if value, present := os.LookupEnv(envPrefix + "Compress"); present {
		parsedVal, err := strconv.ParseBool(value)
		if err != nil {
			log.Fatalln("Invalid value for Log:Compress\n" + err.Error())
		}
		Log.Compress = parsedVal
	} else {
		Log.Compress = config.Log.Compress
	}

	return nil
}