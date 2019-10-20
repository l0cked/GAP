package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

var config Config

type Config struct {
	BindAddress      string `toml:"bind_addr"`
	HttpDirectory    string `toml:"http_dir"`
	DatabaseFileName string `toml:"db_filename"`
}

func (cfg *Config) Init() {
	cfg.BindAddress = ":8080"
	cfg.HttpDirectory = "http"
	cfg.DatabaseFileName = "gap.db"

	_, err := toml.DecodeFile("gap.toml", &cfg)
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create("gap.toml")
			if err != nil {
				panic(err)
			}
			_, _ = f.WriteString("bind_addr = ':8080'\n")
			_, _ = f.WriteString("http_dir = 'http'\n")
			_, _ = f.WriteString("db_filename = 'gap.db'\n")
		} else {
			panic(err)
		}
	}
}
