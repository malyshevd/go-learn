package server

import (
	"fmt"
	"github.com/malyshevd/go-learn/lesson-8/config"
	"net/http"
)

func RunServer(c *config.Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ""
		for i := 0; i < c.Count; i++ {
			data += fmt.Sprintf("Hello, %s\n", c.Name)
		}
		_, err := fmt.Fprint(w, data)
		if err != nil {
			return
		}

	})

	err := http.ListenAndServe("localhost:"+c.Port, nil)
	if err != nil {
		return
	}
}
