package main

import (
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nunoki/gorest/internal/beetroot/postgres"
)

func main() {
	// do we want to migrate up or down?
	var down uint
	flag.UintVar(&down, "down", 0, "Specify to migrate down, defining how many migrations to roll back")
	flag.Parse()

	connStr := postgres.ConnectionStringFromEnv()
	m, err := migrate.New(
		"file://.docker/postgres/",
		connStr,
	)

	if err != nil {
		panic(err)
	}

	if down <= 0 {
		fmt.Print("Running up migrations...\n")
		err = m.Up()
	} else {
		fmt.Printf("Running %d down migration(s)...\n", down)
		err = m.Steps(-int(down))
	}

	if err != nil {
		// non-panic error here, because it might just be "no change"
		fmt.Println(err)
	} else {
		fmt.Println("Finished.")
	}
}
