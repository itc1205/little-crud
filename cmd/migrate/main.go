package main

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/itc1205/little-crud/internal/bootstrap"
	"github.com/itc1205/little-crud/internal/config"
	"github.com/itc1205/little-crud/internal/migrate"
)

const helpMsg = "Simple migration tool\nUsage:\n\thelp - Display that message\n\tup - Apply all migrations\n\tdown - Remove all migrations"

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		println(helpMsg)
		os.Exit(0)
	}

	if args[0] == "help" {
		println(helpMsg)
		os.Exit(0)
	}

	if !(args[0] == "up") && !(args[0] == "down") {
		println(helpMsg)
		os.Exit(0)
	}

	pgCfg := config.PostgresConfig{}
	if err := env.Parse(&pgCfg); err != nil {
		log.Fatal("Could not parse migration config:", err)
	}

	chCfg := config.ClickHouseConfig{}
	if err := env.Parse(&chCfg); err != nil {
		log.Fatal("Could not parse migration config:", err)
	}

	pgDb, err := bootstrap.InitSqlDB(pgCfg)
	if err != nil {
		log.Fatal("Could not init postgres connection:", err)
	}

	chDb, err := bootstrap.InitClickHouse(chCfg)
	if err != nil {
		log.Fatal("Could not init clickhouse connection:", err)
	}
	if args[0] == "up" {
		log.Println("Applying up migrations to postgres...")
		if err := migrate.PostgresMigrateUp(pgDb); err != nil {
			log.Fatal("Could not apply migrations to postgres:", err)
		}
		log.Println("Done!")

		log.Println("Applying up migrations to clickhouse...")

		if err := migrate.ClickMigrateUp(chDb); err != nil {
			log.Fatal("Could not apply migrations to clickhouse:", err)
		}
		log.Println("Done!")
	} else if args[0] == "down" {
		log.Println("Applying down migrations to postgres...")
		if err := migrate.PostgresMigrateDown(pgDb); err != nil {
			log.Fatal("Could not apply migrations to postgres:", err)
		}
		log.Println("Done!")

		log.Println("Applying down migrations to clickhouse...")
		if err := migrate.ClickMigrateDown(chDb); err != nil {
			log.Fatal("Could not apply migrations to clickhouse:", err)
		}
		log.Println("Done!")
	}
}
