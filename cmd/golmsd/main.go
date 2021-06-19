package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/abc101/golms/pkg/config"
	"github.com/abc101/golms/pkg/pgdb"
	_ "github.com/lib/pq"
	"github.com/pelletier/go-toml/v2"
)

const (
	// default configuration file
	default_config = "./configs/config.toml"
)

func main() {
	fmt.Println("Starting Go LMS: ")

	// Read options
	cFile := flag.String("config", default_config, "a config file")
	flag.Parse()

	// Read config
	cf, err := ioutil.ReadFile(*cFile)
	if err != nil {
		panic(err)
	}

	// Database connection
	fmt.Print("  * Database ----- ")
	config := config.Config{}
	toml.Unmarshal(cf, &config)
	pginfo := config.Postgres

	db, err := pgdb.Connect(pginfo)

	defer db.Close()

	// TEST: DB connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("OK!")
}
