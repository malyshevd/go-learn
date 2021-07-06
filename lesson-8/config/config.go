package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port  string
	Name  string
	Count int
}

func NewConfig() *Config {

	fileName := "./lesson-8/config/config.json"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	fs, err := os.Stat(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	data := make([]byte, fs.Size())
	_, err = f.Read(data)
	if err != nil {
		log.Fatalln(err)
	}

	config := Config{}
	err = json.Unmarshal(data, &config)

	if config.Port == "" || config.Name == "" {
		log.Fatalf("config.json incorrect: port=%v; name=%v;\n", config.Port, config.Name)
	}

	if config.Count == 0 {
		config.Count = 1
	}

	/*port := flag.String("port", "8080", "port number")
	name := flag.String("name", "username", "name of user")
	count := flag.Int("count", 1, "count of repeat")
	flag.Parse()*/

	return &config
}
