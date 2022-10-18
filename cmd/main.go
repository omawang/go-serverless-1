package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"example.portless.io/simple-go-func/handler"
	"gopkg.in/yaml.v3"
)

type config struct {
	Trigger string `yaml:"trigger"`
	Event   string `yaml:"event"`
}

func main() {
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic("failed read config.yaml file " + err.Error())
	}

	var cfg config
	err = yaml.Unmarshal(configFile, &cfg)
	if err != nil {
		panic("failed parsing config.yaml " + err.Error())
	}

	if cfg.Trigger != "HTTP" {
		panic(fmt.Sprintf("not supported trigger \"%s\"", cfg.Trigger))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != cfg.Event {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("{\"message\":\"method not allowed\"}"))
		} else {
			handler.Handler(w, req)
		}
	})

	http.ListenAndServe(":8090", nil)
}
