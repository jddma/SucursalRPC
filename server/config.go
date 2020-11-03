package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type serverConfig struct {

	Port string
	Ssl bool
	SslCertPath string
	SslKeyPath string

}

func newServerConfig() *serverConfig {

	var result serverConfig
	configStr, err := ioutil.ReadFile("resources/config.json")
	if err != nil{
		log.Fatal(err.Error())
	}

	json.Unmarshal(configStr, &result)

	return &result

}