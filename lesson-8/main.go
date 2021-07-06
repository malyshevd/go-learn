package main

import (
	"fmt"
	"github.com/malyshevd/go-learn/lesson-8/config"
	"github.com/malyshevd/go-learn/lesson-8/server"
)

func main() {
	config := config.NewConfig()
	fmt.Println(*config)

	server.RunServer(config)
}
