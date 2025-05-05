package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port       int
	LdapHost   string
	LdapDomain string
}

var config Config

func loadConfig(filename string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatal(err)
	}

	if _, err := toml.DecodeFile(filename, &config); err != nil {
		log.Fatal(err)
	}
}
