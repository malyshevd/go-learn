package config

import (
	"flag"
)

type Config struct {
	Port  string
	Name  string
	Count int
}

func NewConfig() *Config {

	port := flag.String("port", "8080", "port number")
	name := flag.String("name", "username", "name of user")
	count := flag.Int("count", 1, "count of repeat")
	flag.Parse()

	return &Config{
		Port:  *port,
		Name:  *name,
		Count: *count,
	}
}
