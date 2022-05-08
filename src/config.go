package main

type Config struct {
	DatabasePath   string
	DatabaseName   string
	CollectionName string
	ServerPort     string
}

func NewConfig() *Config {
	return &Config{
		DatabasePath:   "mongodb://root:example@localhost:27017",
		DatabaseName:   "basec",
		CollectionName: "users-datas",
		ServerPort:     "3000",
	}
}
