package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cowboiscott/hippo/library/config"
	"github.com/cowboiscott/hippo/library/database"
)

func main() {
	configFile := flag.String("config", "config.toml", "Path to the configuration file")

	flag.Parse()

	config, err := config.NewConfig(*configFile)
	if err != nil {
		fmt.Printf("Error: could not load config: %v\n", err)
		os.Exit(1)
	}

	_, err = database.NewDB(config)
	if err != nil {
		fmt.Printf("Error: could not load database: %v\n", err)
		os.Exit(1)
	}
}
